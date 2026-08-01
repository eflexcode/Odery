package com.ifeanyi.OderService.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.ifeanyi.OderService.entity.OrderStatus;
import lombok.Data;

import java.util.Date;

@Data
public class OrderModel {

    @JsonProperty("product_id")
    private String productId;
    @JsonProperty("product_currency")
    private String productCurrency;
    @JsonProperty("product_img_url")
    private String productImgUrl;
    @JsonProperty("product_name")
    private String productName;
    @JsonProperty("user_id")
    private String userId;
    @JsonProperty("count")
    private int count;
    @JsonProperty("description")
    private String description;

}
