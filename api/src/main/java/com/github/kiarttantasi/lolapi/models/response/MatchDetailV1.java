package com.github.kiarttantasi.lolapi.models.response;

import java.util.ArrayList;
import java.util.List;
import lombok.Getter;

@Getter
public class MatchDetailV1 {
  private final String championName;
  private final Integer kills;
  private final Integer deaths;
  private final Integer assists;
  private final Boolean win;
  private final String gameMode;
  private final Long gameCreation;
  private final List<ParticipantV1> participantList;
  private final List<Integer> itemIds;

  public MatchDetailV1(
      String championName,
      Integer kills,
      Integer deaths,
      Integer assists,
      Boolean win,
      String gameMode,
      Long gameCreation,
      List<ParticipantV1> participantList,
      List<Integer> itemIds
  ) {
    this.championName = championName;
    this.kills = kills;
    this.deaths = deaths;
    this.assists = assists;
    this.win = win;
    this.gameMode = gameMode;
    this.gameCreation = gameCreation;
    this.participantList = new ArrayList<>(participantList);
    this.itemIds = new ArrayList<>(itemIds);
  }

  public List<ParticipantV1> getParticipantList() {
    return new ArrayList<>(participantList);
  }

  public List<Integer> getItemIds() {
    return new ArrayList<>(itemIds);
  }
}
