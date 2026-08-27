package com.ifeanyi.Gateway.security.filter;

import com.ifeanyi.Gateway.exception.Unauthorized;
import com.ifeanyi.Gateway.util.Util;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.security.Keys;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.ReactiveSecurityContextHolder;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ResponseStatusException;
import org.springframework.web.server.ServerWebExchange;
import org.springframework.web.server.WebFilter;
import org.springframework.web.server.WebFilterChain;
import reactor.core.publisher.Mono;

import java.nio.charset.StandardCharsets;
import java.util.Date;
import java.util.List;
import java.util.function.Function;

@Slf4j
@Component
public class SecurityFilter implements WebFilter {

    private boolean verifyToken(String token) throws Unauthorized {

        Date date = extractClaims(token,Claims::getExpiration);

        return new Date().before(date);
    }

    public <T> T extractClaims(String token,Function<Claims,T> tFunction)throws Unauthorized {

        Claims claims = Jwts.parserBuilder()
                .setSigningKey(Keys.hmacShaKeyFor(Util.TOKEN_KEY.getBytes(StandardCharsets.UTF_8)))
                .build().parseClaimsJws(token)
                .getBody();

        return tFunction.apply(claims);
    }

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, WebFilterChain chain) {

        String header = exchange.getRequest().getHeaders().getFirst("Authorization");
        String path = exchange.getRequest().getURI().getPath();

        if (path.startsWith("/auth")) {
            return chain.filter(exchange);
        }

        if (header != null && header.startsWith("Bearer ")) {

            String token = header.substring(7).trim();

            boolean authPassed;

            try {
                 authPassed = verifyToken(token);
            } catch (Unauthorized e) {
                return Mono.error(new ResponseStatusException(
                        HttpStatus.UNAUTHORIZED, "Auth failed"));
            }

            if(!authPassed){

                return Mono.error(new ResponseStatusException(
                        HttpStatus.UNAUTHORIZED, "Auth failed"));
//                throw new ResponseStatusException(HttpStatus.UNAUTHORIZED,"Auth failed");
            }

            String username = null;
            try {
                username = extractClaims(token, Claims::getSubject);
            } catch (Unauthorized e) {
                return Mono.error(new ResponseStatusException(
                        HttpStatus.UNAUTHORIZED, "Auth failed"));
            }

            Authentication authentication = new UsernamePasswordAuthenticationToken(username,null, List.of());

            return chain.filter(exchange).contextWrite(ReactiveSecurityContextHolder.withAuthentication(authentication));
        }

        return chain.filter(exchange);
    }

}

