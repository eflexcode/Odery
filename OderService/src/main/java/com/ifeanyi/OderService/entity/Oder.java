package com.ifeanyi.OderService.entity;

import com.fasterxml.jackson.annotation.JsonProperty;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import java.util.Date;

@Document
public class Oder {

    @Id
    private String id;
    @JsonProperty("product_id")
    private String productId;
    @JsonProperty("description")
    private String description;
    @JsonProperty("status")
    private OderStatus status;
    @JsonProperty("created_at")
    private Date createdAt;
    @JsonProperty("updated_at")
    private Date updatedAt;

}
