package com.github.kiarttantasi.lolapi.models.riot.response;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import edu.umd.cs.findbugs.annotations.SuppressFBWarnings;

@JsonIgnoreProperties(ignoreUnknown = true)
public class MatchDetailResponse {
  private MatchInfo info;

  @SuppressFBWarnings // to suppress "M C NP: Read of unwritten field"
  public MatchInfo getInfo() {
    return this.info.deepClone();
  }
}
