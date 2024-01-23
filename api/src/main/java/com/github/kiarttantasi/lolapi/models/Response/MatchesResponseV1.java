package com.github.kiarttantasi.lolapi.models.Response;

import java.util.List;

import lombok.Getter;
import lombok.RequiredArgsConstructor;

@RequiredArgsConstructor
@Getter
public class MatchesResponseV1 {
    private final List<MatchDetailV1> matches;
}
