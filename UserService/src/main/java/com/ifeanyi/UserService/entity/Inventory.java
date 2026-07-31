package com.ifeanyi.UserService.entity;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import lombok.Data;

import java.util.Date;

@Data
@Entity(name = "inventory")
public class Inventory {

    @Id
    @JsonProperty("id")
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;
    @JsonProperty("user_id")
    private String userId;
    @JsonProperty("product_id")
    private String productId;
    @JsonProperty("created_at")
    private Date createdAt;

}
