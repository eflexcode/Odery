package com.ifeanyi.Gateway.config;

import org.springframework.cloud.gateway.route.builder.RouteLocatorBuilder;
import org.springframework.context.annotation.Configuration;
import org.springframework.cloud.gateway.route.RouteLocator;
//import org.springframework.cloud.gateway.route.RouteLocatorBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class GatewayRoutingConfig {
    @Bean
    public RouteLocator routeLocator(RouteLocatorBuilder builder) {

        return builder.routes()
                .route("notification", r -> r
                        .path("/notification/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("http://localhost:8084"))
                .route("auth", r -> r
                        .path("/auth/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("http://localhost:8081"))
                .route("order", r -> r
                        .path("/order/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("http://localhost:8091"))
                .route("payment", r -> r
                        .path("/payment/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("http://localhost:8089"))
                .route("product", r -> r
                        .path("/product/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("http://localhost:8092"))
                .route("user", r -> r
                        .path("/user/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("http://localhost:8095"))
                .route("discovery", r -> r
                        .path("/discovery/**")
                        .filters(f -> f.stripPrefix(1))
                        .uri("http://localhost:8085"))
                .build();

    }

}
