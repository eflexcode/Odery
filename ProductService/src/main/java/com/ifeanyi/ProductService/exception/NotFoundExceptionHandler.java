package com.ifeanyi.ProductService.exception;

public class NotFoundExceptionHandler extends Exception{
    public NotFoundExceptionHandler(String message) {
        super(message);
    }
}
