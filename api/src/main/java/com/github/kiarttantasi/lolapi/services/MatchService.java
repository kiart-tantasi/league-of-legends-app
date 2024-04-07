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
    final String puuid = getPuuid(gameName, tagLine);
    final String[] matchIds = getMatchIds(puuid);
    return getMatchDetailV1List(matchIds, gameName);
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

  private List<MatchDetailV1> getMatchDetailV1List(String[] matchIds, String gameName) {
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
    return getMatchDetailV1ListHelper(completables, gameName);
  }

  private List<MatchDetailV1> getMatchDetailV1ListHelper(
      List<CompletableFuture<HttpResponse<String>>> completables, String gameName) {
    final List<MatchDetailV1> matchDetails = new ArrayList<>();
    for (final CompletableFuture<HttpResponse<String>> completable : completables) {
      try {
        final MatchDetailResponse response =
            new ObjectMapper().readValue(completable.get().body(), MatchDetailResponse.class);
        final List<ParticipantV1> participants = new ArrayList<>();
        ParticipantV1 self = null;
        for (Participant parti : response.getInfo().getParticipants()) {
          try {
            final ParticipantV1 newParti =
                ParticipantV1.builder().gameName(parti.getRiotIdGameName())
                    .tagLine(parti.getRiotIdTagline()).championName(parti.getChampionName())
                    .kills(parti.getKills()).deaths(parti.getDeaths()).assists(parti.getAssists())
                    .win(parti.getWin()).itemIds(parti.getItemIds()).build();
            participants.add(newParti);
            if (parti.getRiotIdGameName().equals(gameName)) {
              self = newParti;
            }
          } catch (Exception e) {
            log.error(e.getMessage());
          }
        }
        if (self != null) {
          matchDetails.add(
              MatchDetailV1.builder().championName(self.getChampionName()).kills(self.getKills())
                  .deaths(self.getDeaths()).assists(self.getAssists()).win(self.getWin())
                  .gameMode(response.getInfo().getGameMode())
                  .gameCreation(response.getInfo().getGameCreation())
                  .participantList(participants).itemIds(self.getItemIds()).build());
        }
      } catch (Exception e) {
        log.error(e.getMessage());
      }
    }
    return matchDetails;
  }
}
