package com.ifeanyi.OderService.service.OtherService.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

@Data
public class Payment {

    @JsonProperty("id")
    private String id;
    @JsonProperty("user_id")
    private String userId;
    @JsonProperty("card_id")
    private String cardId;
    @JsonProperty("amount")
    private float amount;
    @JsonProperty("product_id")
    private String productId;
    @JsonProperty("order_id")
    private String orderId;
    @JsonProperty("status")
    private String Status;
    @JsonProperty("reason")
    private String Reason;
    @JsonProperty("type")
    private String type;
    @JsonProperty("item_count")
    private int itemCount;
    @JsonProperty("description")
    private String description;
    @JsonProperty("create_at")
    private String createdAt;
    @JsonProperty("updated_at")
    private String updatedAt;

}
