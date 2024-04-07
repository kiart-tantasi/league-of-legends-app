package com.github.kiarttantasi.lolapi.models.response;

import java.util.ArrayList;
import java.util.List;

public record MatchesResponseV1(List<MatchDetailV1> matchDetailList) {
  public MatchesResponseV1(List<MatchDetailV1> matchDetailList) {
    this.matchDetailList = new ArrayList<>(matchDetailList);
  }

  @Override
  public List<MatchDetailV1> matchDetailList() {
    return new ArrayList<>(this.matchDetailList);
  }
}
