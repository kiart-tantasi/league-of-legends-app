package com.github.kiarttantasi.lolapi.services;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.URLEncoder;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.charset.StandardCharsets;
import java.util.List;

@Service
public class MatchHistoryService {

    @Value("${riot.api.key}")
    private String riotApiKey;

    // TODO: create private methods to separate logic
    public List<String> getMatches() {
        // STEP1: get puuid from gameName and tagLine
        final String encodedGameName = URLEncoder.encode("เพชร", StandardCharsets.UTF_8);
        final String tagLine = "ARAM";
        final String apiRegion = "asia";
        final URI uri;
        try {
            uri = new URI(String.format("https://%s.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s", apiRegion, encodedGameName, tagLine));
            System.out.println("URI:" + uri.toString());
        } catch (URISyntaxException e) {
            System.out.println(e.getMessage());
            throw new InternalError();
        }
        final HttpRequest request = HttpRequest.newBuilder().uri(uri).header("X-Riot-Token", riotApiKey).build();
        final HttpClient client = HttpClient.newHttpClient();
        HttpResponse<String> response;
        try {
            response = client.send(request, HttpResponse.BodyHandlers.ofString());
            final String body = response.body();
            System.out.println("RESPONSE BODY: " + body);
        } catch (IOException | InterruptedException e) {
            System.out.println(e.getMessage());
            throw new InternalError();
        }

        // STEP2: get 10 matches from puuid
        //

        return null;
    }
}
