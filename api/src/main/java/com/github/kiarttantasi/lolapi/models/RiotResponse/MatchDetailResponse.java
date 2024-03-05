package com.github.kiarttantasi.lolapi.models.RiotResponse;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Getter;

@JsonIgnoreProperties(ignoreUnknown = true)
@Getter
public class MatchDetailResponse {
  private MatchInfo info;
}
