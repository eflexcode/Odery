package com.ifeanyi.UserService.exception.model;

import lombok.Data;
import lombok.RequiredArgsConstructor;

import java.util.Date;

@Data
@RequiredArgsConstructor
public class ErrorModel {

    private String message;
    private Date date;

}
