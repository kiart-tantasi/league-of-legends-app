package com.github.kiarttantasi.lolapi.configurations;

import lombok.Getter;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Configuration;

@Configuration
@Getter
public class RiotConfig {
    @Value("${riot.api.key}")
    private String riotApiKey;
    // TEMP USE, we don't have production key so we hit rate limit so often. let's use two 2 dev keys at the same time
    @Value("${riot.api.key.secondary:#{null}}")
    private String riotApiKeySecondary;
    @Value("${riot.api.region.account}")
    private String regionAccount;
    @Value("${riot.api.region.match}")
    private String regionMatch;
    @Value("#{new Integer('${riot.api.match.amount}')}")
    private Integer matchAmount;

    // simple load-balancer to select riot api key
    public String getRiotApiKey() {
        if (this.riotApiKeySecondary != null && Math.random() > 0.5) {
            return this.riotApiKeySecondary;
        }
        return this.riotApiKey;
    }
}
