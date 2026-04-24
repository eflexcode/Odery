package com.ifeanyi.ProductService.service;

import com.ifeanyi.ProductService.entity.Product;
import com.ifeanyi.ProductService.model.ProductModel;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;

import java.util.Optional;

public interface ProductService {

    Product create(ProductModel productModel);
    Product update(ProductModel productModel);
    Page<Product> getAll(Pageable pageable);
    Page<Product> findByProductNameContaining(String productName, Pageable pageable);
    Optional<Product> findByUserId(String userId);
    Page<Product> findByInStockBetween(int minZero,int max,Pageable pageable);

}
