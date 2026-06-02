package com.ifeanyi.OderService.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.Date;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class StandardResp {
    private String message;
    private int status;
    private Date date;
}
