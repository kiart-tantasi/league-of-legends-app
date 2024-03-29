package com.github.kiarttantasi.lolapi.models.RiotResponse;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Getter;

@JsonIgnoreProperties(ignoreUnknown = true)
@Getter
public class MatchInfo {
  private Participant[] participants;
  private String gameMode;
  private Long gameCreation;
}
