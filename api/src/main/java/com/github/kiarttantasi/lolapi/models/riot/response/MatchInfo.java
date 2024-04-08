package com.github.kiarttantasi.lolapi.models.riot.response;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.util.Arrays;
import lombok.Getter;
import lombok.Setter;

@JsonIgnoreProperties(ignoreUnknown = true)
@Getter
@Setter
public class MatchInfo {
  private Participant[] participants;
  private String gameMode;
  private Long gameCreation;

  public Participant[] getParticipants() {
    return this.participants.clone();
  }

  public void setParticipants(Participant[] participants) {
    this.participants = Arrays.copyOf(participants, participants.length);
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
