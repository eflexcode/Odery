package com.ifeanyi.ProductService.service;

import com.ifeanyi.ProductService.entity.Product;
import com.ifeanyi.ProductService.exception.NotFoundExceptionHandler;
import com.ifeanyi.ProductService.model.ProductModel;
import com.ifeanyi.ProductService.model.StandardResponse;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.ResponseEntity;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.Optional;

public interface ProductService {

    Product create(MultipartFile file_img,ProductModel productModel) throws NotFoundExceptionHandler, IOException;
    Product update(String id,ProductModel productModel) throws NotFoundExceptionHandler;
    Product get(String id) throws NotFoundExceptionHandler;
    StandardResponse delete(String id);
    Page<Product> getAll(Pageable pageable);
    Page<Product> findByProductNameContaining(String productName, Pageable pageable);
    Page<Product> findByUserId(String userId,Pageable pageable);
    Page<Product> findByInStockBetween(int minZero,int max,Pageable pageable);

    byte[] getProductImg(String fileName) throws IOException;
}
