package com.github.kiarttantasi.lolapi.configurations;

import lombok.Getter;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Configuration;

@Configuration
@Getter
public class RiotConfig {
  @Value("${riot.api.key}")
  private String riotApiKey;
  @Value("${riot.api.region.account}")
  private String regionAccount;
  @Value("${riot.api.region.match}")
  private String regionMatch;
  @Value("#{new Integer('${riot.api.match.amount}')}")
  private Integer matchAmount;
}
