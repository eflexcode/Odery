package com.ifeanyi.ProductService.service.impl;

import com.ifeanyi.ProductService.entity.Product;
import com.ifeanyi.ProductService.exception.NotFoundExceptionHandler;
import com.ifeanyi.ProductService.model.ProductModel;
import com.ifeanyi.ProductService.model.StandardResponse;
import com.ifeanyi.ProductService.repository.ProductRepository;
import com.ifeanyi.ProductService.service.CategoryService;
import com.ifeanyi.ProductService.service.ProductService;
import com.ifeanyi.ProductService.service.impl.OtherServices.model.User;
import com.ifeanyi.ProductService.util.Util;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.BeanUtils;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.*;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.util.Date;
import java.util.Optional;

@Service
@RequiredArgsConstructor
public class ProductServiceImpl implements ProductService {

    private final ProductRepository repository;
    private final CategoryService categoryService;

    private final RestTemplate restTemplate;

    @Override
    public Product create(ProductModel productModel) throws NotFoundExceptionHandler {

        categoryService.get(productModel.getCategoryId());
        User user = getUserFromUserService(productModel.getUserId());

        if (user != null && user.getId().equals(productModel.getUserId())){

        }

        Product product = new Product();
        //TODO check if user is an admin
        BeanUtils.copyProperties(productModel, product);
        Date date = new Date();
        product.setCreatedAt(date);
        product.setUpdatedAt(date);

        Product savedProduct = repository.save(product);

        return savedProduct;
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

    public User getUserFromUserService(String id) {
        String endpoint = "" + id;

        ResponseEntity<User> userResponseEntity = restTemplate.getForEntity(Util.USER_SERVICE_BASE_URL + endpoint, User.class);
        if (userResponseEntity.getStatusCode() != HttpStatus.OK) {
            return null;
        }
        return userResponseEntity.getBody();
    }

}
