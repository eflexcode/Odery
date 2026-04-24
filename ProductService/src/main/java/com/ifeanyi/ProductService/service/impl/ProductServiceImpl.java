package com.ifeanyi.ProductService.service.impl;

import com.ifeanyi.ProductService.entity.Product;
import com.ifeanyi.ProductService.model.ProductModel;
import com.ifeanyi.ProductService.repository.ProductRepository;
import com.ifeanyi.ProductService.service.ProductService;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import java.util.Optional;

@Service
@RequiredArgsConstructor
public class ProductServiceImpl implements ProductService {

    private final ProductRepository repository;

    @Override
    public Product create(ProductModel productModel) {
        return null;
    }

    @Override
    public Product update(ProductModel productModel) {
        return null;
    }

    @Override
    public Page<Product> getAll(Pageable pageable) {
        return repository.findAll(pageable);
    }

    @Override
    public Page<Product> findByProductNameContaining(String productName, Pageable pageable) {
        return null;
    }

    @Override
    public Optional<Product> findByUserId(String userId) {
        return Optional.empty();
    }

    @Override
    public Page<Product> findByInStockBetween(int minZero, int max, Pageable pageable) {
        return null;
    }
}
