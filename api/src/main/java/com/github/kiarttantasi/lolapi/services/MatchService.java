package com.github.kiarttantasi.lolapi.services;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.kiarttantasi.lolapi.models.Response.MatchDetailV1;
import com.github.kiarttantasi.lolapi.models.RiotResponse.AccountResponse;
import com.github.kiarttantasi.lolapi.models.RiotResponse.MatchDetailResponse;
import com.github.kiarttantasi.lolapi.models.RiotResponse.Participant;

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
    private static final Integer MATCH_AMOUNT = 20;

    @Value("${riot.api.key:no-key-found}")
    private String riotApiKey;
    @Value("${riot.api.region.account}")
    private String regionAccount;
    @Value("${riot.api.region.match}")
    private String regionMatch;

    public List<MatchDetailV1> getMatches(String gameName, String tagLine)
            throws URISyntaxException, IOException, InterruptedException {
        final String puuid = getPuuid(gameName, tagLine);
        final String[] matchIds = getMatchIds(puuid);
        return getMatchesV1SendAsync(matchIds, gameName);
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
                        regionMatch, puuid, MATCH_AMOUNT));
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
                log.error(e.getMessage());
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
                            parti.getWin(),
                            response.getInfo().getGameMode(),
                            response.getInfo().getGameCreation()));
                });
            } catch (Exception e) {
                log.error(e.getMessage());
            }
        }
        return matchDetails;
    }
}
