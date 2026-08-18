package com.ifeanyi.UserService.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import lombok.Data;

import java.util.Date;

@Data
public class InventoryModel {

    @JsonProperty("user_id")
    private String userId;
    @JsonProperty("product_id")
    private String productId;
    @JsonProperty("order_id")
    private String orderId;

}
