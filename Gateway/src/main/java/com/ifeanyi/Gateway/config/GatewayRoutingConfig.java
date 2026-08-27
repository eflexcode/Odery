package com.ifeanyi.Gateway.config;

import org.springframework.context.annotation.Configuration;
import org.springframework.cloud.gateway.route.RouteLocator;
import org.springframework.cloud.gateway.route.RouteLocatorBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class GatewayRoutingConfig {

    @Bean
    public RouteLocator(RouteLocatorBuilder builder) {

        return builder.routes()
                .route("url", r -> r
                        .path("/url/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("http://localhost:8083"))
                .route("auth", r -> r
                        .path("/auth/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("http://localhost:8084"))
                .route("user", r -> r
                        .path("/user/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("http://localhost:8082"))
                .build();

    }

}
