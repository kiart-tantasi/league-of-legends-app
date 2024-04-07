package com.github.kiarttantasi.lolapi.models.riot.response;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
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
    final MatchInfo clone = new MatchInfo();
    clone.setParticipants(Arrays.copyOf(this.participants, this.participants.length));
    clone.setGameMode(this.getGameMode());
    clone.setGameCreation(this.getGameCreation());
    return clone;
  }
}
