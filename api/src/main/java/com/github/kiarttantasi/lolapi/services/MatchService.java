package com.github.kiarttantasi.lolapi.services;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.kiarttantasi.lolapi.configurations.RiotConfig;
import com.github.kiarttantasi.lolapi.models.Response.MatchDetailV1;
import com.github.kiarttantasi.lolapi.models.Response.ParticipantV1;
import com.github.kiarttantasi.lolapi.models.RiotResponse.AccountResponse;
import com.github.kiarttantasi.lolapi.models.RiotResponse.MatchDetailResponse;
import com.github.kiarttantasi.lolapi.models.RiotResponse.Participant;

import lombok.extern.slf4j.Slf4j;

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

    private final RiotConfig riotConfig;

    public MatchService(RiotConfig riotConfig) {
        this.riotConfig = riotConfig;
    }

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
                String.format("https://%s.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s", riotConfig.getRegionAccount(),
                        encodedGameName, tagLine));
        final HttpRequest request = HttpRequest.newBuilder().uri(uriAccount).header("X-Riot-Token", riotConfig.getRiotApiKey()).build();
        final HttpResponse<String> response = HttpClient.newHttpClient().send(request, BODYHANDLER);
        final String body = response.body();
        final AccountResponse accountResponse = new ObjectMapper().readValue(body, AccountResponse.class);
        return accountResponse.getPuuid();
    }

    private String[] getMatchIds(String puuid) throws URISyntaxException, IOException, InterruptedException {
        final URI uri = new URI(
                String.format("https://%s.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=%d",
                        riotConfig.getRegionMatch(), puuid, riotConfig.getMatchAmount()));
        final HttpRequest request = HttpRequest.newBuilder().uri(uri).header("X-Riot-Token", riotConfig.getRiotApiKey()).build();
        final HttpResponse<String> response = HttpClient.newHttpClient().send(request, BODYHANDLER);
        return new ObjectMapper().readValue(response.body(), String[].class);
    }

    private List<MatchDetailV1> getMatchDetailV1List(String[] matchIds, String gameName) {
        final List<CompletableFuture<HttpResponse<String>>> completables = new ArrayList<>();
        final String consistentRiotApiKey = this.riotConfig.getRiotApiKey();
        for (final String matchId : matchIds) {
            URI uri;
            try {
                uri = new URI(
                        String.format("https://%s.api.riotgames.com/lol/match/v5/matches/%s", riotConfig.getRegionMatch(), matchId));
                final HttpRequest request = HttpRequest.newBuilder().uri(uri).header("X-Riot-Token", consistentRiotApiKey)
                        .build();
                completables.add(HttpClient.newHttpClient().sendAsync(request, BODYHANDLER));
            } catch (URISyntaxException e) {
                log.error(e.getMessage());
            }
        }
        return mapMatchDetailList(completables, gameName);
    }

    // TODO: break-down this method (especially at double-try-catch-scope) for less complexity
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
                if (user == null) {
                    continue;
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
