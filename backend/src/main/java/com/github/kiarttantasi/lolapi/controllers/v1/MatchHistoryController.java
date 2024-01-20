package com.github.kiarttantasi.lolapi.controllers.v1;

import com.github.kiarttantasi.lolapi.models.MatchesResponseV1;
import com.github.kiarttantasi.lolapi.services.MatchService;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/matches")
public class MatchHistoryController {

    private final MatchService matchService;

    public MatchHistoryController(MatchService matchHistoryService) {
        this.matchService = matchHistoryService;
    }

    @GetMapping
    public ResponseEntity<MatchesResponseV1> getMatches(@RequestParam String gameName, @RequestParam String tagLine) {
        try {
            return ResponseEntity.ok().body(
                    new MatchesResponseV1(
                            matchService.getMatches(gameName, tagLine)));
        } catch (Exception e) {
            System.out.println(e.getMessage());
            return ResponseEntity.internalServerError().build();
        }
    }
}
