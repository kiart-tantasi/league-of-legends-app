package com.github.kiarttantasi.lolapi.models;

import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor
public class MatchDetailV1 {
    private final String championName; // TODO: implement with actual type
    private final Integer kills;
    private final Integer deaths;
    private final Integer assists;
    private final Boolean win;
}
