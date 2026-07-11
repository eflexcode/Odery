package com.ifeanyi.ProductService.service.impl.OtherServices.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

import java.util.Date;

@Data
public class User {

    private String id;
    private String username;
    private String password;
    private Role role;
    private String name;
    private String imgUrl;

    private Date createdAt;
    private Date updatedAt;

}
