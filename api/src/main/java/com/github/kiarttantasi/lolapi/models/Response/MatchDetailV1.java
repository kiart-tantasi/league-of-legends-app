package com.github.kiarttantasi.lolapi.models.Response;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public class MatchDetailV1 {
    private final String championName;
    private final Integer kills;
    private final Integer deaths;
    private final Integer assists;
    private final Boolean win;
    private final String gameMode;
    private final Long gameCreation;
    private final List<ParticipantV1> participantList;
}
