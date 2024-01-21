package com.github.kiarttantasi.lolapi.services;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.kiarttantasi.lolapi.models.AccountResponse;
import com.github.kiarttantasi.lolapi.models.MatchDetailResponse;
import com.github.kiarttantasi.lolapi.models.MatchDetailV1;
import com.github.kiarttantasi.lolapi.models.Participant;

import lombok.extern.slf4j.Slf4j;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.URLEncoder;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.charset.Charset;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Optional;
import java.util.concurrent.*;

@Service
@Slf4j
public class MatchService {

    private static final Charset ENCODING_CHARSET = StandardCharsets.UTF_8;
    private static final HttpResponse.BodyHandler<String> BODYHANDLER = HttpResponse.BodyHandlers.ofString();

    @Value("${match.amount:10}")
    private Integer matchAmount;
    @Value("${riot.api.key:no-key-found}")
    private String riotApiKey;
    @Value("${riot.api.region.account}")
    private String regionAccount;
    @Value("${riot.api.region.match}")
    private String regionMatch;

    public List<MatchDetailV1> getMatches(String gameName, String tagLine)
            throws URISyntaxException, IOException, InterruptedException {
        long start = System.currentTimeMillis();
        final String puuid = getPuuid(gameName, tagLine);
        log.info("got puuid in " + (System.currentTimeMillis() - start));

        start = System.currentTimeMillis();
        final String[] matchIds = getMatchIds(puuid);
        log.info("got matchIds in " + (System.currentTimeMillis() - start));

        start = System.currentTimeMillis();
        final List<MatchDetailV1> matches = getMatchesV1SendAsync(matchIds, gameName);
        log.info("got matches in " + (System.currentTimeMillis() - start));

        return matches;
    }

    private String getPuuid(String gameName, String tagLine)
            throws URISyntaxException, IOException, InterruptedException {
        final String encodedGameName = URLEncoder.encode(gameName, ENCODING_CHARSET);
        final URI uriAccount = new URI(
                String.format("https://%s.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s", regionAccount,
                        encodedGameName, tagLine));
        final HttpRequest request = HttpRequest.newBuilder().uri(uriAccount).header("X-Riot-Token", riotApiKey).build();
        final HttpResponse<String> response = HttpClient.newHttpClient().send(request, BODYHANDLER);
        final String body = response.body();
        final AccountResponse accountResponse = new ObjectMapper().readValue(body, AccountResponse.class);
        return accountResponse.getPuuid();
    }

    private String[] getMatchIds(String puuid) throws URISyntaxException, IOException, InterruptedException {
        final URI uri = new URI(
                String.format("https://%s.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=%d",
                        regionMatch, puuid, matchAmount));
        final HttpRequest request = HttpRequest.newBuilder().uri(uri).header("X-Riot-Token", riotApiKey).build();
        final HttpResponse<String> response = HttpClient.newHttpClient().send(request, BODYHANDLER);
        final String[] matchIds = new ObjectMapper().readValue(response.body(), String[].class);
        return matchIds;
    }

    private List<MatchDetailV1> getMatchesV1SendAsync(String[] matchIds, String gameName) {
        // send requests asynchronously
        final List<CompletableFuture<HttpResponse<String>>> completables = new ArrayList<>();
        for (final String matchId : matchIds) {
            URI uri;
            try {
                uri = new URI(
                        String.format("https://%s.api.riotgames.com/lol/match/v5/matches/%s", regionMatch, matchId));
                final HttpRequest request = HttpRequest.newBuilder().uri(uri).header("X-Riot-Token", riotApiKey)
                        .build();
                completables.add(HttpClient.newHttpClient().sendAsync(request, BODYHANDLER));
            } catch (URISyntaxException e) {
                log.error(e.getMessage(), e);
            }
        }
        // get value from completables and map to MatchDetailV1
        final List<MatchDetailV1> matchDetails = new ArrayList<>();
        for (final CompletableFuture<HttpResponse<String>> completable : completables) {
            try {
                final MatchDetailResponse response = new ObjectMapper().readValue(completable.get().body(),
                        MatchDetailResponse.class);
                final Optional<Participant> optiParti = Arrays.stream(response.getInfo().getParticipants())
                        .filter(x -> {
                            return x.getRiotIdGameName().equals(gameName);
                        }).findFirst();
                optiParti.ifPresent(parti -> {
                    matchDetails.add(new MatchDetailV1(
                            parti.getChampionName(),
                            parti.getKills(),
                            parti.getDeaths(),
                            parti.getAssists(),
                            parti.getWin()));
                });
            } catch (Exception ex) {
                log.error(ex.getMessage(), ex);
            }
        }
        return matchDetails;
    }
}
