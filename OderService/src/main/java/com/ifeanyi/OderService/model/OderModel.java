package com.ifeanyi.OderService.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.ifeanyi.OderService.entity.OderStatus;
import lombok.Data;

import java.util.Date;

@Data
public class OderModel {

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
