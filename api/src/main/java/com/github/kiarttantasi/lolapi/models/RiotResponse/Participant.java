package com.github.kiarttantasi.lolapi.models.RiotResponse;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

import lombok.Getter;

@JsonIgnoreProperties(ignoreUnknown = true)
@Getter
public class Participant {
    private String riotIdGameName;
    private Integer kills;
    private Integer deaths;
    private Integer assists;
    private String championName;
    private Boolean win;
}
