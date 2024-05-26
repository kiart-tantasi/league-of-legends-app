package com.github.kiarttantasi.lolapi.services;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.kiarttantasi.lolapi.configurations.RiotConfig;
import com.github.kiarttantasi.lolapi.models.response.MatchDetailV1;
import com.github.kiarttantasi.lolapi.models.response.ParticipantV1;
import com.github.kiarttantasi.lolapi.models.riot.response.AccountResponse;
import com.github.kiarttantasi.lolapi.models.riot.response.MatchDetailResponse;
import com.github.kiarttantasi.lolapi.models.riot.response.Participant;
import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.URLEncoder;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.charset.Charset;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.CompletableFuture;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

@Service
@Slf4j
public class MatchService {
  private static final Charset ENCODING_CHARSET = StandardCharsets.UTF_8;
  private final ApiService apiService;
  private final RiotConfig riotConfig;

  public MatchService(RiotConfig riotConfig, ApiService apiService) {
    this.riotConfig = riotConfig;
    this.apiService = apiService;
  }

  public List<MatchDetailV1> getMatches(String gameName, String tagLine)
      throws URISyntaxException, IOException, InterruptedException {
    // puuid
    final String puuid = getPuuid(gameName, tagLine);
    if (puuid == null) {
      return null;
    }
    // match ids
    final String[] matchIds = getMatchIds(puuid);
    if (matchIds == null || matchIds.length == 0) {
      return null;
    }
    // matches
    return getMatchDetailV1List(matchIds, puuid);
  }

  private String getPuuid(String gameName, String tagLine)
      throws URISyntaxException, IOException, InterruptedException {
    final String encodedGameName = URLEncoder.encode(gameName, ENCODING_CHARSET);
    final URI uriAccount = new URI(
        String.format("https://%s.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s",
            riotConfig.getRegionAccount(),
            encodedGameName, tagLine));
    final HttpRequest request =
        HttpRequest.newBuilder().uri(uriAccount)
            .header("X-Riot-Token", this.riotConfig.getRiotApiKey())
            .build();
    final AccountResponse accountResponse = apiService.send(request, AccountResponse.class);
    if (accountResponse == null) {
      return null;
    }
    return accountResponse.getPuuid();
  }

  private String[] getMatchIds(String puuid)
      throws URISyntaxException, IOException, InterruptedException {
    final URI uri = new URI(
        String.format(
            "https://%s.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=%d",
            riotConfig.getRegionMatch(), puuid, riotConfig.getMatchAmount()));
    final HttpRequest request =
        HttpRequest.newBuilder().uri(uri).header("X-Riot-Token", this.riotConfig.getRiotApiKey())
            .build();
    return apiService.send(request, String[].class);
  }

  private List<MatchDetailV1> getMatchDetailV1List(String[] matchIds, String puuid) {
    final List<CompletableFuture<HttpResponse<String>>> completables = new ArrayList<>();
    for (final String matchId : matchIds) {
      try {
        final URI uri = new URI(
            String.format("https://%s.api.riotgames.com/lol/match/v5/matches/%s",
                riotConfig.getRegionMatch(), matchId));
        final HttpRequest request =
            HttpRequest.newBuilder().uri(uri)
                .header("X-Riot-Token", this.riotConfig.getRiotApiKey())
                .build();
        completables.add(apiService.sendAsync(request));
      } catch (URISyntaxException e) {
        log.error(e.getMessage());
      }
    }
    return getMatchDetailV1ListHelper(completables, puuid);
  }

  private List<MatchDetailV1> getMatchDetailV1ListHelper(
      List<CompletableFuture<HttpResponse<String>>> completables, String puuid) {
    final List<MatchDetailV1> matchDetails = new ArrayList<>();
    for (final CompletableFuture<HttpResponse<String>> completable : completables) {
      try {
        final MatchDetailResponse response =
            new ObjectMapper().readValue(completable.get().body(), MatchDetailResponse.class);
        final List<ParticipantV1> participants = new ArrayList<>();
        final MatchDetailV1 matchDetailV1 = new MatchDetailV1();
        for (Participant parti : response.getInfo().getParticipants()) {
          try {
            // all cases
            final ParticipantV1 newParti = new ParticipantV1(
                parti.getRiotIdGameName(),
                parti.getRiotIdTagline(),
                parti.getChampionName(),
                parti.getKills(),
                parti.getDeaths(),
                parti.getAssists(),
                parti.getWin(),
                parti.getItemIds()
            );
            participants.add(newParti);
            // id owner case
            if (parti.getPuuid().equals(puuid)) {
              matchDetailV1.setChampionName(parti.getChampionName());
              matchDetailV1.setKills(parti.getKills());
              matchDetailV1.setDeaths(parti.getDeaths());
              matchDetailV1.setAssists(parti.getAssists());
              matchDetailV1.setWin(parti.getWin());
              matchDetailV1.setItemIds(parti.getItemIds());
            }
          } catch (Exception e) {
            log.error(e.getMessage());
          }
        }
        matchDetailV1.setGameMode(response.getInfo().getGameMode());
        matchDetailV1.setGameCreation(response.getInfo().getGameCreation());
        matchDetailV1.setParticipantList(participants);
        matchDetails.add(matchDetailV1);
      } catch (Exception e) {
        log.error(e.getMessage());
      }
    }
    return matchDetails;
  }
}
