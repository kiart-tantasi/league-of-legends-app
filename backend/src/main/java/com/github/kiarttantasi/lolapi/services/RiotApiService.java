package com.github.kiarttantasi.lolapi.services;

import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class RiotApiService {

    // [how to get match history from riot api]
    // 1. /riot/account/v1/accounts/by-riot-id/{gameName}/{tagLine}
    // 2. /lol/match/v5/matches/by-puuid/{puuid}/ids
    // 3. /lol/match/v5/matches/{matchId}

    public List<String> getMatches() {
        System.out.println("getting matches from Riot API...");
        return null;
    }
}
