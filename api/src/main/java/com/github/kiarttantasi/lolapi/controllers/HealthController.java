package com.github.kiarttantasi.lolapi.controllers;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/health")
public class HealthController {
    @GetMapping
    public ResponseEntity<Integer> getHealth() {
        return ResponseEntity.ok().build();
    }

    @PostMapping
    public ResponseEntity<String> postHealth() {
        return ResponseEntity.badRequest().build();
    }
}
