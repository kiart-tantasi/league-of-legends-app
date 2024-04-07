package com.github.kiarttantasi.lolapi.models.response;

import java.util.ArrayList;
import java.util.List;
import lombok.Getter;

@Getter
public class ParticipantV1 {
  private final String gameName;
  private final String tagLine;
  private final String championName;
  private final Integer kills;
  private final Integer deaths;
  private final Integer assists;
  private final Boolean win;
  private final List<Integer> itemIds;

  // prevent exposing internal rep
  public ParticipantV1(String gameName,
                       String tagLine,
                       String championName,
                       Integer kills,
                       Integer deaths,
                       Integer assists,
                       Boolean win,
                       List<Integer> itemIds
  ) {
    this.gameName = gameName;
    this.tagLine = tagLine;
    this.championName = championName;
    this.kills = kills;
    this.deaths = deaths;
    this.assists = assists;
    this.win = win;
    this.itemIds = new ArrayList<>(itemIds);
  }

  public List<Integer> getItemIds() {
    return new ArrayList<>(this.itemIds);
  }
}
