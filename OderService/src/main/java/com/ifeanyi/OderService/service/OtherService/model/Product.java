package com.ifeanyi.OderService.service.OtherService.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;
import org.springframework.data.annotation.Id;

import java.util.Date;

@Data
public class Product {

    @Id
    private String id;
    @JsonProperty("category_id")
    private String categoryId;
    @JsonProperty("product_name")
    private String productName;
    @JsonProperty("user_id")
    private String userId;
    @JsonProperty("product_description")
    private String productDescription;
    @JsonProperty("product_img_url")
    private String productImgUrl;
    @JsonProperty("price")
    private int price;
    @JsonProperty("in_stock")
    private int inStock;
    @JsonProperty("create_at")
    private Date createdAt;
    @JsonProperty("updated_at")
    private Date updatedAt;

}
