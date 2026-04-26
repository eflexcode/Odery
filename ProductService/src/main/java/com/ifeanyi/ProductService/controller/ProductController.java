package com.ifeanyi.ProductService.controller;

import com.ifeanyi.ProductService.entity.Product;
import com.ifeanyi.ProductService.exception.NotFoundExceptionHandler;
import com.ifeanyi.ProductService.model.ProductModel;
import com.ifeanyi.ProductService.service.ProductService;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequiredArgsConstructor
@RequestMapping("/")
public class ProductController {

    private final ProductService productService;

    @PostMapping("create")
    public ResponseEntity<Product> create(@RequestBody ProductModel productModel) {
        return new ResponseEntity<>(productService.create(productModel), HttpStatus.CREATED);
    }

    @GetMapping("{id}")
    public ResponseEntity<Product> get(@PathVariable String id) throws NotFoundExceptionHandler {
        return new ResponseEntity<>(productService.get(id), HttpStatus.OK);
    }

    @PutMapping("update/{id}")
    public ResponseEntity<Product> update(@PathVariable String id, @RequestBody ProductModel productModel) throws NotFoundExceptionHandler {
        return new ResponseEntity<>(productService.update(id, productModel), HttpStatus.OK);
    }

    @GetMapping("get-all")
    public Page<Product> getAll(Pageable pageable) {
        return productService.getAll(pageable);
    }

    @GetMapping("search/{name}")
    @ResponseStatus(HttpStatus.OK)
    public Page<Product> getByProductNameContaining(@PathVariable(name = "name") String productName, Pageable pageable) {
        return productService.findByProductNameContaining(productName, pageable);
    }

    @GetMapping("get-by-userId/{id}")
    @ResponseStatus(HttpStatus.OK)
    public Page<Product> getByUserId(@PathVariable(name = "id") String userId, Pageable pageable) {
        return productService.findByUserId(userId, pageable);
    }

    @GetMapping("available")
    @ResponseStatus(HttpStatus.OK)
    public Page<Product> getByInStockBetween( @PathVariable(name = "max") int max, Pageable pageable) {
        return productService.findByInStockBetween(1, max, pageable);
    }


}
