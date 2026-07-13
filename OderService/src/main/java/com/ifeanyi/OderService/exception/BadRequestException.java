package com.ifeanyi.OderService.exception;

import org.springframework.web.bind.annotation.ControllerAdvice;

public class BadRequestException extends Exception{

    public BadRequestException(String message) {
        super(message);
    }
}
