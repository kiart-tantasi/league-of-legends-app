package com.github.kiarttantasi.lolapi.services;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.kiarttantasi.lolapi.models.AccountResponse;
import com.github.kiarttantasi.lolapi.models.MatchDetailResponse;
import lombok.Getter;
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
import java.util.Collection;
import java.util.List;
import java.util.concurrent.*;

@Service
public class MatchHistoryService {

    private static final Charset ENCODING_CHARSET = StandardCharsets.UTF_8;
    private static final HttpResponse.BodyHandler<String> BODYHANDLER = HttpResponse.BodyHandlers.ofString();
    private static final int THREAD_AMOUNT = 10; // TODO: change this in EC2 to proper number

    @Value("${riot.api.key:no-key-found}")
    private String riotApiKey;

    @Value("${riot.api.region.account}")
    private String regionAccount;

    @Value("${riot.api.region.match}")
    private String regionMatch;

    // TODO: create private methods to separate logic
    public void getMatches(String gameName, String tagLine)
            throws URISyntaxException, IOException, InterruptedException {
        // shared
        final HttpClient client = HttpClient.newHttpClient();
        final ObjectMapper objectMapper = new ObjectMapper();

        // STEP1: get puuid from gameName and tagLine
        final String encodedGameName = URLEncoder.encode(gameName, ENCODING_CHARSET);
        final URI uriAccount = new URI(
                String.format("https://%s.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s", regionAccount,
                        encodedGameName, tagLine));
        final HttpRequest request = HttpRequest.newBuilder().uri(uriAccount).header("X-Riot-Token", riotApiKey).build();
        final HttpResponse<String> response = client.send(request, BODYHANDLER);
        final String body = response.body();
        final AccountResponse accountResponse = objectMapper.readValue(body, AccountResponse.class);

        // STEP2: get 10 matches from puuidd
        final String puuid = accountResponse.getPuuid();
        final URI uri2 = new URI(
                String.format("https://%s.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=10",
                        regionMatch, puuid));
        final HttpRequest request2 = HttpRequest.newBuilder().uri(uri2).header("X-Riot-Token", riotApiKey).build();
        final HttpResponse<String> response2 = client.send(request2, BODYHANDLER);
        final String[] matchIds = objectMapper.readValue(response2.body(), String[].class);

        // STEP3: get all match info from match list
        getMatchesDetails(matchIds);
    }

    private void getMatchesDetails(String[] matchIds) throws InterruptedException {
        final long start = System.currentTimeMillis();
        final ExecutorService ex = Executors.newFixedThreadPool(THREAD_AMOUNT);
        // invokeAll already blocks so no need to .shutdown() or .awaitTermination()
        // ref: https://stackoverflow.com/q/44742444
        final List<Future<Void>> futures = ex.invokeAll(getMatchCallables(matchIds));
        for (final Future<Void> future : futures) {
            try {
                future.get();
            } catch (ExecutionException e) {
                System.out.println(e.getMessage());
            }
        }

        final long end = System.currentTimeMillis();
        System.out.println("got all matches' details in " + (end - start) + " ms");
    }

    private Collection<Callable<Void>> getMatchCallables(String[] matchIds) {
        final List<Callable<Void>> callables = new ArrayList<>();
        for (String matchId : matchIds) {
            final Callable<Void> callable = new Callable<>() {
                @Override
                public Void call() throws Exception {
                    final MatchDetailResponse matchDetailResponse = getMatchDetail(matchId);
                    return null;
                }
            };
            callables.add(callable);
        }
        return callables;
    }

    private MatchDetailResponse getMatchDetail(String matchId) throws URISyntaxException, IOException, InterruptedException {
        final HttpClient client = HttpClient.newHttpClient();
        final URI uri = new URI(
                String.format("https://%s.api.riotgames.com/lol/match/v5/matches/%s", regionMatch, matchId));
        final HttpRequest request = HttpRequest.newBuilder().uri(uri).header("X-Riot-Token", riotApiKey).build();
        final HttpResponse<String> response = client.send(request, BODYHANDLER);
        return new ObjectMapper().readValue(response.body(), MatchDetailResponse.class);
    }
}


