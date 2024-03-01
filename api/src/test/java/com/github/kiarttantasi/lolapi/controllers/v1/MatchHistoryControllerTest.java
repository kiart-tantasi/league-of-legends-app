package com.github.kiarttantasi.lolapi.controllers.v1;

import com.github.kiarttantasi.lolapi.services.MatchService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.web.servlet.MockMvc;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(MatchHistoryController.class)
public class MatchHistoryControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockBean
    private MatchService matchService;

    @Test
    public void getMatchesOk() throws Exception {
        when(matchService.getMatches(any(), any())).thenReturn(null);
        mockMvc.perform(get("/api/v1/matches?gameName=foo&tagLine=bar")).andExpect(status().isOk());
    }

    @Test
    public void getMatchesBadRequestNoParams() throws Exception {
        mockMvc.perform(get("/api/v1/matches")).andExpect(status().isBadRequest());
    }
}
