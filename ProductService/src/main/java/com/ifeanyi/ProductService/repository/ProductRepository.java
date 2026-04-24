package com.ifeanyi.ProductService.repository;

import com.ifeanyi.ProductService.entity.Product;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface ProductRepository extends MongoRepository<Product,String> {

    Page<Product> findByProductNameContaining(String productName, Pageable pageable);
    Page<Product> findByUserId(String userId,Pageable pageable);
    Page<Product> findByInStockBetween(int minZero,int max,Pageable pageable);

}
