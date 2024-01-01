package com.github.kiarttantasi.lolapi.controllers;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/match")
public class MatchHistory {

    @GetMapping
    public String getMatches(@RequestParam String gameName, @RequestParam String tagLine) {
        System.out.println("gameName: " + gameName);
        System.out.println("tagLine: " + tagLine);
        return "match!";
    }
}
