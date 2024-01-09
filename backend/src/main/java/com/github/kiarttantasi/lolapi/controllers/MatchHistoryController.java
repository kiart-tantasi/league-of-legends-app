package com.github.kiarttantasi.lolapi.controllers;

import com.github.kiarttantasi.lolapi.services.MatchHistoryService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/match")
public class MatchHistoryController {

    private final MatchHistoryService matchHistoryService;

    public MatchHistoryController(MatchHistoryService matchHistoryService) {
        this.matchHistoryService = matchHistoryService;
    }

    @GetMapping
    public ResponseEntity<String> getMatches(@RequestParam String gameName, @RequestParam String tagLine) {
        try {
            matchHistoryService.getMatches(gameName, tagLine);
        } catch (Exception e) {
            System.out.println(e.getMessage());
            return ResponseEntity.internalServerError().build();
        }
        return ResponseEntity.ok().body("matches");
    }
}
