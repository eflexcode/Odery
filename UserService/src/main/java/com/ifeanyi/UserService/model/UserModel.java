package com.ifeanyi.UserService.model;

import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import lombok.Data;

@Data
@Entity
public class UserModel {

    @Id
    private Integer id;
    private String username;
    private String password;
    private String name;
    private String imgUrl;

}
