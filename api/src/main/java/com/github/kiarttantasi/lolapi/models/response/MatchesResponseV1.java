package com.github.kiarttantasi.lolapi.models.response;

import java.util.List;
import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor
@Getter
public class MatchesResponseV1 {
  private final List<MatchDetailV1> matchDetailList;
}
