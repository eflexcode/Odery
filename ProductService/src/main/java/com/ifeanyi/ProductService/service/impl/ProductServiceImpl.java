package com.ifeanyi.ProductService.service.impl;

import com.ifeanyi.ProductService.entity.Product;
import com.ifeanyi.ProductService.exception.NotFoundExceptionHandler;
import com.ifeanyi.ProductService.model.ProductModel;
import com.ifeanyi.ProductService.model.StandardResponse;
import com.ifeanyi.ProductService.repository.ProductRepository;
import com.ifeanyi.ProductService.service.CategoryService;
import com.ifeanyi.ProductService.service.ProductService;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.BeanUtils;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.Optional;

@Service
@RequiredArgsConstructor
public class ProductServiceImpl implements ProductService {

    private final ProductRepository repository;
    private final CategoryService categoryService;

    @Override
    public Product create(ProductModel productModel) throws NotFoundExceptionHandler {

        categoryService.get(productModel.getCategoryId());

        Product product = new Product();

        BeanUtils.copyProperties(productModel, product);
        Date date = new Date();
        product.setCreatedAt(date);
        product.setUpdatedAt(date);

        return repository.save(product);
    }

    @Override
    public Product update(String id, ProductModel productModel) throws NotFoundExceptionHandler {

        Product product = get(id);
        BeanUtils.copyProperties(productModel, product);
        product.setUpdatedAt(new Date());

        return repository.save(product);
    }

    @Override
    public Product get(String id) throws NotFoundExceptionHandler {
        return repository.findById(id).orElseThrow(() -> new NotFoundExceptionHandler("No product found with id: " + id));
    }

    @Override
    public StandardResponse delete(String id) {
        repository.deleteById(id);
        return new StandardResponse("Product deleted successfully", 200, new Date());
    }

    @Override
    public Page<Product> getAll(Pageable pageable) {
        return repository.findAll(pageable);
    }

    @Override
    public Page<Product> findByProductNameContaining(String productName, Pageable pageable) {
        return repository.findByProductNameContaining(productName, pageable);
    }

    @Override
    public Page<Product> findByUserId(String userId, Pageable pageable) {
        return repository.findByUserId(userId, pageable);
    }

    @Override
    public Page<Product> findByInStockBetween(int minZero, int max, Pageable pageable) {
        return repository.findByInStockBetween(minZero, max, pageable);
    }

}
