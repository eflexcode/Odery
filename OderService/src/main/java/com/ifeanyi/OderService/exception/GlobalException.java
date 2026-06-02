package com.ifeanyi.OderService.exception;

import com.ifeanyi.OderService.model.StandardResp;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

import java.util.Date;

@ControllerAdvice
public class GlobalException extends ResponseEntityExceptionHandler {

    @ExceptionHandler(NotFoundException.class)
    public  ResponseEntity<StandardResp> notFound(NotFoundException exception){
        return new ResponseEntity<>(new StandardResp("Not found",HttpStatus.NOT_FOUND.value(),new Date()), HttpStatus.NOT_FOUND);
    }

}

