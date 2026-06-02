package com.ifeanyi.ProductService.repository;

import com.ifeanyi.ProductService.entity.Category;
import com.ifeanyi.ProductService.entity.Product;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface CategoryRepository extends MongoRepository<Category,String> {
}
