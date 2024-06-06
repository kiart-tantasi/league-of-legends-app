package com.github.kiarttantasi.lolapi.services;

import com.fasterxml.jackson.databind.ObjectMapper;
import java.io.IOException;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.concurrent.CompletableFuture;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;

@Service
@Slf4j
public class ApiService {
  private static final HttpResponse.BodyHandler<String> BODYHANDLER =
      HttpResponse.BodyHandlers.ofString();
  private final HttpClient httpClient;

  public ApiService(HttpClient httpClient) {
    this.httpClient = httpClient;
  }

  public <T> T send(HttpRequest request, Class<T> mappingClass)
      throws IOException, InterruptedException {
    final HttpResponse<String> response = httpClient.send(request, BODYHANDLER);
    if (response.statusCode() != HttpStatus.OK.value()) {
      log.warn(
          String.format("request with URI %s got status code %s and response body %s",
              response.uri(),
              response.statusCode(),
              response.body()));
      return null;
    }
    return new ObjectMapper().readValue(response.body(), mappingClass);
  }

  public CompletableFuture<HttpResponse<String>> sendAsync(HttpRequest request) {
    return httpClient.sendAsync(request, BODYHANDLER);
  }
}
