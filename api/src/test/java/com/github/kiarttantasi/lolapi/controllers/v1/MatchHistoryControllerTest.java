package com.github.kiarttantasi.lolapi.controllers.v1;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import com.github.kiarttantasi.lolapi.models.response.MatchDetailV1;
import com.github.kiarttantasi.lolapi.services.MatchService;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.web.servlet.MockMvc;

@WebMvcTest(MatchHistoryController.class)
public class MatchHistoryControllerTest {

  @Autowired
  private MockMvc mockMvc;

  @MockBean
  private MatchService matchService;

  @Test
  public void getMatchesOkNonEmptyList() throws Exception {
    final List<MatchDetailV1> matchesMock = new ArrayList<>(Collections.singletonList(mockMatch()));
    when(matchService.getMatches(any(), any())).thenReturn(matchesMock);
    mockMvc.perform(get("/api/v1/matches?gameName=foo&tagLine=bar")).andExpect(status().isOk());
  }

  @Test
  public void getMatchesOkEmptyList() throws Exception {
    final List<MatchDetailV1> matchesMock = new ArrayList<>();
    when(matchService.getMatches(any(), any())).thenReturn(matchesMock);
    mockMvc.perform(get("/api/v1/matches?gameName=foo&tagLine=bar")).andExpect(status().isOk());
  }

  // TOOO: test case when user is not found (404)

  @Test
  public void getMatchesBadRequestNoParams() throws Exception {
    mockMvc.perform(get("/api/v1/matches")).andExpect(status().isBadRequest());
  }

  private MatchDetailV1 mockMatch() {
    return new MatchDetailV1(
        "MOCK",
        1,
        2,
        3,
        true,
        "MOCK",
        123L,
        new ArrayList<>(),
        new ArrayList<>()
    );
  }
}
