package com.ifeanyi.UserService.entity;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import lombok.Data;

import java.util.Date;

@Data
@Entity
public class User {

    @Id
    private Integer id;
    private String username;
    private String password;
    private String name;
    @JsonProperty("img_url")
    private String imgUrl;

    @JsonProperty("create_at")
    private Date createdAt;
    @JsonProperty("updated_at")
    private Date updatedAt;

}