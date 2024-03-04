package com.github.kiarttantasi.lolapi.filters;

import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.Cookie;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.extern.slf4j.Slf4j;
import org.springframework.core.annotation.Order;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;
import java.util.regex.Pattern;
import java.util.stream.Collectors;

@Component
@Order
@Slf4j
public class UserFilter extends OncePerRequestFilter {

    private static final String USER_ID_COOKIE_NAME = "user_id";

    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) throws ServletException, IOException {
        try {
            if (shouldSkipFilter(request)) {
                return;
            }
            final Map<String, String> cookiesMap = initCookiesMap(request);
            String userId = getCookie(cookiesMap, USER_ID_COOKIE_NAME);
            if (userId == null) {
                userId = UUID.randomUUID().toString();
                response.addCookie(createCookie(USER_ID_COOKIE_NAME, userId, true, true, 3));
            }
            log.info("user id: " + userId);
        } catch (Exception e) {
            log.error(e.getMessage(), e);
        } finally {
            filterChain.doFilter(request, response);
        }
    }

    private Map<String, String> initCookiesMap(HttpServletRequest request) {
        final Cookie[] cookies = request.getCookies();
        if (cookies == null) {
            return new HashMap<>();
        }
        return Arrays.stream(cookies).collect(Collectors.toMap(c -> formatCookieName(c.getName()), Cookie::getValue));
    }

    private String getCookie(Map<String, String> cookiesMap, String cookieName) {
        return cookiesMap.get(formatCookieName(cookieName));
    }

    private String formatCookieName(String cookieName) {
        if (cookieName == null) {
            return null;
        }
        return cookieName.toLowerCase();
    }

    private Cookie createCookie(String cookieName, String cookieValue, boolean isHttpOnly, boolean isSecure, int expireMonth) {
        final Cookie cookie = new Cookie(cookieName, cookieValue);
        cookie.setHttpOnly(isHttpOnly);
        cookie.setSecure(isSecure);
        expireMonth = Math.max(expireMonth, 1);
        cookie.setMaxAge(60 * 60 * 24 * 30 * expireMonth);
        return cookie;
    }

    private boolean shouldSkipFilter(HttpServletRequest request) {
        return !Pattern.compile("^/api/v1/").matcher(request.getRequestURI()).find();
    }
}
