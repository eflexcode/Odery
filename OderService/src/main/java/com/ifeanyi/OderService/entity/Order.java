package com.ifeanyi.OderService.entity;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import java.util.Date;

@Document(collation = "orders")
@Data
public class Order {

    @Id
    private String id;
    @JsonProperty("product_id")
    private String productId;
    @JsonProperty("product_name")
    private String productName;
    @JsonProperty("product_currency")
    private String productCurrency;
    @JsonProperty("product_img_url")
    private String productImgUrl;
    @JsonProperty("count")
    private int count;
    @JsonProperty("amount")
    private int amount;
    @JsonProperty("user_id")
    private String userId;
    @JsonProperty("description")
    private String description;
    @JsonProperty("status")
    private OrderStatus status;
    @JsonProperty("created_at")
    private Date createdAt;
    @JsonProperty("updated_at")
    private Date updatedAt;

}
