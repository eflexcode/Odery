package com.ifeanyi.UserService.entity;

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
    private String imgUrl;

    private Date createdAt;
    private Date updatedAt;

}