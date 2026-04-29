package com.ifeanyi.OderService.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.ifeanyi.OderService.entity.OrderStatus;
import lombok.Data;

import java.util.Date;

@Data
public class OrderModel {

    @JsonProperty("product_id")
    private String productId;
    @JsonProperty("description")
    private String description;
    @JsonProperty("status")
    private OrderStatus status;
    @JsonProperty("created_at")
    private Date createdAt;
    @JsonProperty("updated_at")
    private Date updatedAt;

}
