package com.ifeanyi.ProductService.config;

import com.ifeanyi.ProductService.service.impl.OtherServices.User.UserService;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestTemplate;

@Configuration
public class Config {

    @Bean
    public RestTemplate restTemplate(){
        return new RestTemplate();
    }

    @Bean
    public UserService userService(){
        return new UserService(restTemplate());
    }

}
