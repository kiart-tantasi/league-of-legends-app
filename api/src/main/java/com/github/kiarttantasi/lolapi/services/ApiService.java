package com.github.kiarttantasi.lolapi.services;

import com.fasterxml.jackson.databind.ObjectMapper;
import java.io.IOException;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.concurrent.CompletableFuture;
import org.springframework.stereotype.Service;

@Service
public class ApiService {
  private static final HttpResponse.BodyHandler<String> BODYHANDLER =
      HttpResponse.BodyHandlers.ofString();

  public <T> T send(HttpRequest request, Class<T> mappingClass)
      throws IOException, InterruptedException {
    final HttpResponse<String> response = HttpClient.newHttpClient().send(request, BODYHANDLER);
    return new ObjectMapper().readValue(response.body(), mappingClass);
  }

  public CompletableFuture<HttpResponse<String>> sendAsync(HttpRequest request) {
    return HttpClient.newHttpClient().sendAsync(request, BODYHANDLER);
  }
}
