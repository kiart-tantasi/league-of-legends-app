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
    private static final Integer MATCH_AMOUNT = 20;

    @Value("${thread.amount}")
    private Integer threadAmount;
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
        return getMatchesV1(matchIds, gameName);
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

    /**
     * V1: return champion, kills, deaths and assists
     */
    private List<MatchDetailV1> getMatchesV1(String[] matchIds, String gameName) throws InterruptedException {
        final List<Future<MatchDetailResponse>> futures = getFutures(matchIds);
        final List<MatchDetailV1> matchDetails = new ArrayList<>();
        for (final Future<MatchDetailResponse> future : futures) {
            try {
                final Optional<Participant> optiParti = Arrays.stream(future.get().getInfo().getParticipants())
                        .filter(x -> {
                            return x.getRiotIdGameName().equals(gameName);
                        }).findFirst();
                optiParti.ifPresent(parti -> {
                    matchDetails.add(new MatchDetailV1(
                            parti.getChampionName(),
                            parti.getKills(),
                            parti.getDeaths(),
                            parti.getAssists()));
                });
            } catch (ExecutionException e) {
                log.error(e.getMessage(), e);
            }
        }
        return matchDetails;
    }

    private List<Future<MatchDetailResponse>> getFutures(String[] matchIds) throws InterruptedException {
        final List<Callable<MatchDetailResponse>> callables = new ArrayList<>();
        for (String matchId : matchIds) {
            final Callable<MatchDetailResponse> callable = new Callable<>() {
                @Override
                public MatchDetailResponse call() throws Exception {
                    return mapMatchDetailResponse(matchId);
                }
            };
            callables.add(callable);
        }
        final ExecutorService ex = Executors.newFixedThreadPool(threadAmount);
        return ex.invokeAll(callables);
    }

    private MatchDetailResponse mapMatchDetailResponse(String matchId)
            throws URISyntaxException, IOException, InterruptedException {
        final HttpClient client = HttpClient.newHttpClient();
        final URI uri = new URI(
                String.format("https://%s.api.riotgames.com/lol/match/v5/matches/%s", regionMatch, matchId));
        final HttpRequest request = HttpRequest.newBuilder().uri(uri).header("X-Riot-Token", riotApiKey).build();
        final HttpResponse<String> response = client.send(request, BODYHANDLER);
        return new ObjectMapper().readValue(response.body(), MatchDetailResponse.class);
    }
}
