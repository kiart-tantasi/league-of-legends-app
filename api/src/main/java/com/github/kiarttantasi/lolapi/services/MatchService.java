package com.github.kiarttantasi.lolapi.services;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.kiarttantasi.lolapi.models.Response.MatchDetailV1;
import com.github.kiarttantasi.lolapi.models.Response.ParticipantV1;
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
import java.util.List;
import java.util.concurrent.*;

@Service
@Slf4j
public class MatchService {

    private static final Charset ENCODING_CHARSET = StandardCharsets.UTF_8;
    private static final HttpResponse.BodyHandler<String> BODYHANDLER = HttpResponse.BodyHandlers.ofString();
    private static final Integer MATCH_AMOUNT = 20;

    @Value("${riot.api.key}")
    private String riotApiKey;
    @Value("${riot.api.region.account}")
    private String regionAccount;
    @Value("${riot.api.region.match}")
    private String regionMatch;

    public List<MatchDetailV1> getMatches(String gameName, String tagLine)
            throws URISyntaxException, IOException, InterruptedException {
        final String puuid = getPuuid(gameName, tagLine);
        final String[] matchIds = getMatchIds(puuid);
        return getMatchDetailV1List(matchIds, gameName);
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

    private List<MatchDetailV1> getMatchDetailV1List(String[] matchIds, String gameName) {
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
        return mapMatchDetailList(completables, gameName);
    }

    private List<MatchDetailV1> mapMatchDetailList(List<CompletableFuture<HttpResponse<String>>> completables,
            String gameName) {
        final List<MatchDetailV1> matchDetails = new ArrayList<>();
        for (final CompletableFuture<HttpResponse<String>> completable : completables) {
            try {
                final MatchDetailResponse response = new ObjectMapper().readValue(completable.get().body(),
                        MatchDetailResponse.class);
                final List<ParticipantV1> participants = new ArrayList<>();
                ParticipantV1 user = null;
                for (Participant parti : response.getInfo().getParticipants()) {
                    try {
                        final ParticipantV1 newParti = ParticipantV1.builder().gameName(parti.getRiotIdGameName())
                                .tagLine(parti.getRiotIdTagline()).championName(parti.getChampionName())
                                .kills(parti.getKills()).deaths(parti.getDeaths()).assists(parti.getAssists())
                                .win(parti.getWin()).itemIds(parti.getItemIds()).build();
                        participants.add(newParti);
                        if (parti.getRiotIdGameName().equals(gameName)) {
                            user = newParti;
                        }
                    } catch (Exception e) {
                        log.error(e.getMessage());
                    }
                }
                matchDetails.add(MatchDetailV1.builder().championName(user.getChampionName()).kills(user.getKills())
                        .deaths(user.getDeaths()).assists(user.getAssists()).win(user.getWin())
                        .gameMode(response.getInfo().getGameMode()).gameCreation(response.getInfo().getGameCreation())
                        .participantList(participants).itemIds(user.getItemIds()).build());
            } catch (Exception e) {
                log.error(e.getMessage());
            }
        }
        return matchDetails;
    }
}
