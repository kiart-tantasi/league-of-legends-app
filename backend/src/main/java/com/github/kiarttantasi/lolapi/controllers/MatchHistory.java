package com.github.kiarttantasi.lolapi.controllers;

import com.github.kiarttantasi.lolapi.services.RiotApiService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/match")
public class MatchHistory {

    @Autowired
    private RiotApiService riotApiService;

    @GetMapping
    public String getMatches(@RequestParam String gameName, @RequestParam String tagLine) {
        System.out.println("gameName: " + gameName);
        System.out.println("tagLine: " + tagLine);
        riotApiService.getMatches();
        return "match!";
    }
}
