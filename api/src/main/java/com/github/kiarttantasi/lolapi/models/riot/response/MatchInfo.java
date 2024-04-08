package com.github.kiarttantasi.lolapi.models.riot.response;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import edu.umd.cs.findbugs.annotations.SuppressFBWarnings;
import lombok.Getter;

@JsonIgnoreProperties(ignoreUnknown = true)
@Getter
public class MatchInfo {
  private Participant[] participants;
  private String gameMode;
  private Long gameCreation;

  @SuppressFBWarnings // to suppress "M C NP: Read of unwritten field"
  public Participant[] getParticipants() {
    return this.participants.clone();
  }

  public MatchInfo deepClone() {
    try {
      final ObjectMapper objectMapper = new ObjectMapper();
      return objectMapper.readValue(objectMapper.writeValueAsString(this), MatchInfo.class);
    } catch (JsonProcessingException e) {
      return null;
    }
  }
}
