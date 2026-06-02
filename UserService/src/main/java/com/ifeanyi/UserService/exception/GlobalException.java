package com.ifeanyi.UserService.exception;

import com.ifeanyi.UserService.exception.model.ErrorModel;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;

@ControllerAdvice
public class GlobalException extends Exception{

    @ExceptionHandler(NotFoundException.class)
    public ResponseEntity<ErrorModel> notFoundException(NotFoundException notFoundException){
        return new ResponseEntity<ErrorModel>(new ErrorModel(),HttpStatus.NOT_FOUND);
    }

}
