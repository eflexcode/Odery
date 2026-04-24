package com.ifeanyi.ProductService.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.RequiredArgsConstructor;

import java.util.Date;

@Data
@AllArgsConstructor
public class StandardResponse {

    private String message;
    private int status;
    private Date date;

}
