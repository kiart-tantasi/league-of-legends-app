package com.github.kiarttantasi.lolapi.models.response;

import java.util.List;
import lombok.Builder;
import lombok.Getter;

@Getter
@Builder
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
}
