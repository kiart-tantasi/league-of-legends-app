package com.github.kiarttantasi.lolapi.models.riot;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Getter;

@JsonIgnoreProperties(ignoreUnknown = true)
@Getter
public class MatchDetailResponse {
  private MatchInfo info;
}
