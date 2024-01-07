package com.github.kiarttantasi.lolapi.services;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.kiarttantasi.lolapi.models.AccountResponse;
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

@Service
public class MatchHistoryService {

    private static final Charset ENCODING_CHARSET = StandardCharsets.UTF_8;
    private static final HttpResponse.BodyHandler<String> BODYHANDLER = HttpResponse.BodyHandlers.ofString();

    @Value("${riot.api.key}")
    private String riotApiKey;

    @Value("${riot.api.region.account}")
    private String regionAccount;

    @Value("${riot.api.region.match}")
    private String regionMatch;

    // TODO: create private methods to separate logic
    public void getMatches(String gameName, String tagLine) throws URISyntaxException, IOException, InterruptedException {
        // shared
        final HttpClient client = HttpClient.newHttpClient();
        final ObjectMapper objectMapper = new ObjectMapper();

        // STEP1: get puuid from gameName and tagLine
        final String encodedGameName = URLEncoder.encode(gameName, ENCODING_CHARSET);
        final URI uriAccount = new URI(String.format("https://%s.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s", regionAccount, encodedGameName, tagLine));
        final HttpRequest request = HttpRequest.newBuilder().uri(uriAccount).header("X-Riot-Token", riotApiKey).build();
        final HttpResponse<String> response = client.send(request, BODYHANDLER);
        final String body = response.body();
        final AccountResponse accountResponse = objectMapper.readValue(body, AccountResponse.class);

        // STEP2: get 10 matches from puuid
        final String puuid = accountResponse.getPuuid();
        final URI uri2 = new URI(String.format("https://%s.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=10", regionMatch, puuid));
        final HttpRequest request2 = HttpRequest.newBuilder().uri(uri2).header("X-Riot-Token", riotApiKey).build();
        final HttpResponse<String> response2 = client.send(request2, BODYHANDLER);
        final String[] matches = objectMapper.readValue(response2.body(), String[].class);
        for (String match : matches) {
            System.out.println("MATCH: " + match);
        }

        // STEP3: get all match info from match list
    }
}
