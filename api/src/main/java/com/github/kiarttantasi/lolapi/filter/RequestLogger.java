package com.github.kiarttantasi.lolapi.filter;

import java.io.IOException;

import org.springframework.core.annotation.Order;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.extern.slf4j.Slf4j;

@Component
@Order
@Slf4j
public class RequestLogger extends OncePerRequestFilter {

    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain)
            throws ServletException, IOException {
        try {
            final long start = System.currentTimeMillis();
            filterChain.doFilter(request, response);
            final long end = System.currentTimeMillis();
            final String msg = String.format("%s, %d ms", request.getRequestURI(), (end - start));
            log.info(msg);
        } catch (Exception e) {
            log.error(e.getMessage(), e);
        }
    }
}
