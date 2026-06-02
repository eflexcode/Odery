package com.ifeanyi.ProductService.exception;

import com.ifeanyi.ProductService.model.StandardResponse;
import org.springframework.data.crossstore.ChangeSetPersister;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

import java.util.Date;

@ControllerAdvice
public class GlobalException extends ResponseEntityExceptionHandler {

    @ExceptionHandler(NotFoundExceptionHandler.class)
    public ResponseEntity<StandardResponse> NotFoundException(NotFoundExceptionHandler notFoundExceptionHandler){
        return new ResponseEntity<>(new StandardResponse("Not Found",HttpStatus.NOT_FOUND.value(),new Date()),HttpStatus.NOT_FOUND);
    }

}
