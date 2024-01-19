package com.github.kiarttantasi.lolapi.models;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Getter;

@JsonIgnoreProperties(ignoreUnknown = true)
@Getter
public class MatchDetailResponse {
    private MatchInfo info;
}

@JsonIgnoreProperties(ignoreUnknown = true)
@Getter
class MatchInfo {
    Participant[] participants;
}


@JsonIgnoreProperties(ignoreUnknown = true)
@Getter
class Participant {
    private String riotIdGameName;
    private Integer kills;
    private Integer deaths;
    private Integer assists;
}