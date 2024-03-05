package com.github.kiarttantasi.lolapi.models.Response;

import java.util.List;
import lombok.Builder;
import lombok.Getter;

@Getter
@Builder
public class ParticipantV1 {
  private final String gameName;
  private final String tagLine;
  private final String championName;
  private final Integer kills;
  private final Integer deaths;
  private final Integer assists;
  private final Boolean win;
  private final List<Integer> itemIds;
}
