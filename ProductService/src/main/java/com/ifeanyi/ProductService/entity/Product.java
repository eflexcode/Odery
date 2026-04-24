package com.ifeanyi.ProductService.entity;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document("Products")
@Data
public class Product {

    @Id
    private String id;
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

}
