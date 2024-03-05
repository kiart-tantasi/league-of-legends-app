package com.github.kiarttantasi.lolapi.filters;

import jakarta.servlet.FilterChain;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.extern.slf4j.Slf4j;
import org.springframework.core.annotation.Order;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

@Component
@Order
@Slf4j
public class RequestLogger extends OncePerRequestFilter {
  private static String getUri(HttpServletRequest request) {
    if (request.getQueryString() == null) {
      return request.getRequestURI();
    }
    return request.getRequestURI().concat("?" + request.getQueryString());
  }

  @Override
  protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response,
                                  FilterChain filterChain) {
    try {
      final long start = System.currentTimeMillis();
      filterChain.doFilter(request, response);
      final long end = System.currentTimeMillis();
      final String msg = String.format("%d, %s, %d ms",
          response.getStatus(),
          getUri(request),
          (end - start));
      log.info(msg);
    } catch (Exception e) {
      log.error(e.getMessage());
    }
  }
}
