package com.ifeanyi.ProductService.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

@Data
public class CategoryModel {

    @JsonProperty("name")
    private String name;

}
