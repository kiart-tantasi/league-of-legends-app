package com.github.kiarttantasi.lolapi.models.RiotResponse;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import java.util.ArrayList;
import java.util.List;
import lombok.Getter;

@JsonIgnoreProperties(ignoreUnknown = true)
@Getter
public class Participant {
  private String riotIdGameName;
  private String riotIdTagline;
  private Integer kills;
  private Integer deaths;
  private Integer assists;
  private String championName;
  private Boolean win;
  private Integer item0;
  private Integer item1;
  private Integer item2;
  private Integer item3;
  private Integer item4;
  private Integer item5;
  private Integer item6;
  private String TestdaDoeNdsadas;

  public List<Integer> getItemIds() {
    final List<Integer> list = new ArrayList<>();
    if (item0 != 0) {
      list.add(item0);
    }
    if (item1 != 0) {
      list.add(item1);
    }
    if (item2 != 0) {
      list.add(item2);
    }
    if (item3 != 0) {
      list.add(item3);
    }
    if (item4 != 0) {
      list.add(item4);
    }
    if (item5 != 0) {
      list.add(item5);
    }
    return list;
  }
}
