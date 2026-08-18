package com.ifeanyi.OderService.service.OtherService.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

import java.util.Date;

@Data
public class Inventory {

    private String id;
    private String userId;
    private String productId;
    private int count;
    private String orderId;
    private Date createdAt;

}
