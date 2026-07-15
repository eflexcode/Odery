package com.ifeanyi.ProductService.service.impl.OtherServices.User;

import com.ifeanyi.ProductService.service.impl.OtherServices.model.User;
import com.ifeanyi.ProductService.util.Util;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.client.RestTemplate;

@RequiredArgsConstructor
public class UserService {

    private final RestTemplate restTemplate;

    public User getUserFromUserService(String id) {
        String endpoint = "" + id;

        ResponseEntity<User> userResponseEntity = restTemplate.getForEntity(Util.USER_SERVICE_BASE_URL + endpoint, User.class);
        if (userResponseEntity.getStatusCode() != HttpStatus.OK) {
            return null;
        }
        return userResponseEntity.getBody();
    }
}
