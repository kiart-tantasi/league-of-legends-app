package com.github.kiarttantasi.lolapi.models.riot.response;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Getter;
import lombok.Setter;

@JsonIgnoreProperties(ignoreUnknown = true)
public class MatchDetailResponse {
  private MatchInfo info;

  public MatchInfo getInfo() {
    return this.info.deepClone();
  }

  public void setInfo(MatchInfo info) {
    this.info = info.deepClone();
  }
}
