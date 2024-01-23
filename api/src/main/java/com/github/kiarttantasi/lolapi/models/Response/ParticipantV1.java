package com.github.kiarttantasi.lolapi.models.Response;

import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor
@Getter
public class ParticipantV1 {
    private final String gameName;
    private final String tagLine;
    private final String championName;
    private final Integer kills;
    private final Integer deaths;
    private final Integer assists;
    private final Boolean win;
}
