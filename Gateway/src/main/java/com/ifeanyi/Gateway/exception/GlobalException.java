package com.ifeanyi.Gateway.exception;

import com.ifeanyi.Gateway.exception.model.ErrorModel;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.reactive.result.method.annotation.ResponseEntityExceptionHandler;

import java.util.Date;

@ControllerAdvice
public class GlobalException extends ResponseEntityExceptionHandler {

    @ExceptionHandler(Unauthorized.class)
    public ResponseEntity<ErrorModel> unauthorizedExceptionHandler(Unauthorized unauthorized) {
        return new ResponseEntity<ErrorModel>(new ErrorModel("Auth failed", new Date(), HttpStatus.UNAUTHORIZED.value()), HttpStatus.UNAUTHORIZED);
    }

}
