package com.github.kiarttantasi.lolapi.services;

import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.when;

import com.github.kiarttantasi.lolapi.configurations.RiotConfig;
import com.github.kiarttantasi.lolapi.models.riot.response.AccountResponse;
import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.http.HttpClient;
import java.net.http.HttpHeaders;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.Optional;
import java.util.concurrent.CompletableFuture;
import javax.net.ssl.SSLSession;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

@ExtendWith(MockitoExtension.class)
public class MatchServiceTest {

  @InjectMocks
  private MatchService matchService;

  @Mock
  private RiotConfig riotConfig;

  @Mock
  private ApiService apiService;

  @BeforeEach
  public void beforeEach() {
    when(riotConfig.getRiotApiKey()).thenReturn("mock");
    when(riotConfig.getRegionAccount()).thenReturn("mock");
    when(riotConfig.getRegionMatch()).thenReturn("mock");
    when(riotConfig.getMatchAmount()).thenReturn(5);
  }

  @Test
  public void getMatchesSuccess() throws URISyntaxException, IOException, InterruptedException {
    when(apiService.send(any(), eq(AccountResponse.class))).thenReturn(mockAccountResponse());
    when(apiService.send(any(), eq(String[].class))).thenReturn(mockMatchIds());
    when(apiService.sendAsync(any())).thenReturn(mockCompletable());
    assertNotNull(matchService.getMatches("gameName", "tagLine"));
  }

  private AccountResponse mockAccountResponse() {
    final AccountResponse accountResponse = new AccountResponse();
    accountResponse.setPuuid("mock");
    accountResponse.setGameName("mock");
    accountResponse.setTagLine("mock");
    return accountResponse;
  }

  private String[] mockMatchIds() {
    return new String[] {"mock1", "mock2", "mock3"};
  }

  private CompletableFuture<HttpResponse<String>> mockCompletable() {
    return CompletableFuture.completedFuture(mockHttpResponse());
  }

  private HttpResponse<String> mockHttpResponse() {
    return new HttpResponse<>() {
      @Override
      public int statusCode() {
        return 0;
      }

      @Override
      public HttpRequest request() {
        return null;
      }

      @Override
      public Optional<HttpResponse<String>> previousResponse() {
        return Optional.empty();
      }

      @Override
      public HttpHeaders headers() {
        return null;
      }

      @Override
      public String body() {
        return "mock";
      }

      @Override
      public Optional<SSLSession> sslSession() {
        return Optional.empty();
      }

      @Override
      public URI uri() {
        return null;
      }

      @Override
      public HttpClient.Version version() {
        return null;
      }
    };
  }
}
