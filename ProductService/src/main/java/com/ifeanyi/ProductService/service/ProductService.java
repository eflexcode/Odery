package com.ifeanyi.ProductService.service;

import com.ifeanyi.ProductService.entity.Product;
import com.ifeanyi.ProductService.exception.NotFoundExceptionHandler;
import com.ifeanyi.ProductService.model.ProductModel;
import com.ifeanyi.ProductService.model.StandardResponse;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;

public interface ProductService {

    Product create(MultipartFile file_img, MultipartFile productFile, ProductModel productModel) throws NotFoundExceptionHandler, IOException;
    Product update(String id,ProductModel productModel) throws NotFoundExceptionHandler;
    Product get(String id) throws NotFoundExceptionHandler;
    StandardResponse delete(String id);
    Page<Product> getAll(Pageable pageable);
    Page<Product> findByProductNameContaining(String productName, Pageable pageable);
    Page<Product> findByUserId(String userId,Pageable pageable);
    Page<Product> findByInStockBetween(int minZero,int max,Pageable pageable);

    byte[] getProductFile(String fileName) throws IOException;
}
