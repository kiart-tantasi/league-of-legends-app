package com.github.kiarttantasi.lolapi.models.response;

import java.util.ArrayList;
import java.util.List;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class MatchDetailV1 {
  private String championName;
  private Integer kills;
  private Integer deaths;
  private Integer assists;
  private Boolean win;
  private String gameMode;
  private Long gameCreation;
  private List<ParticipantV1> participantList;
  private List<Integer> itemIds;

  public List<ParticipantV1> getParticipantList() {
    return new ArrayList<>(participantList);
  }

  public void setParticipantList(List<ParticipantV1> participantList) {
    this.participantList = new ArrayList<>(participantList);
  }

  public List<Integer> getItemIds() {
    return new ArrayList<>(itemIds);
  }

  public void setItemIds(List<Integer> itemIds) {
    this.itemIds = new ArrayList<>(itemIds);
  }
}
