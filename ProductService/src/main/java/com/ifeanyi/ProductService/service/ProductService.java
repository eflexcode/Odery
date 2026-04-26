package com.ifeanyi.ProductService.service;

import com.ifeanyi.ProductService.entity.Product;
import com.ifeanyi.ProductService.exception.NotFoundExceptionHandler;
import com.ifeanyi.ProductService.model.ProductModel;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;

import java.util.Optional;

public interface ProductService {

    Product create(ProductModel productModel) throws NotFoundExceptionHandler;
    Product update(String id,ProductModel productModel) throws NotFoundExceptionHandler;
    Product get(String id) throws NotFoundExceptionHandler;
    Page<Product> getAll(Pageable pageable);
    Page<Product> findByProductNameContaining(String productName, Pageable pageable);
    Page<Product> findByUserId(String userId,Pageable pageable);
    Page<Product> findByInStockBetween(int minZero,int max,Pageable pageable);

}
