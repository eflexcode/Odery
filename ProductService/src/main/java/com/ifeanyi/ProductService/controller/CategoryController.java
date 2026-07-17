package com.ifeanyi.ProductService.controller;

import com.ifeanyi.ProductService.entity.Category;
import com.ifeanyi.ProductService.exception.NotFoundExceptionHandler;
import com.ifeanyi.ProductService.model.CategoryModel;
import com.ifeanyi.ProductService.service.CategoryService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/category/")
@RequiredArgsConstructor
public class CategoryController {

    private final CategoryService categoryService;

    @PostMapping("create")
    public ResponseEntity<Category> create(@RequestBody CategoryModel categoryModel) {
        return new ResponseEntity<>(categoryService.create(categoryModel), HttpStatus.CREATED);
    }

    @PutMapping("{id}")
    public ResponseEntity<Category> update(@RequestBody CategoryModel categoryModel,@PathVariable(name = "id") String id) throws NotFoundExceptionHandler {
        return new ResponseEntity<>(categoryService.update(categoryModel,id),HttpStatus.OK);
    }

    @GetMapping("{id}")
    public ResponseEntity<Category> get(@PathVariable(name = "id") String id) throws NotFoundExceptionHandler {
        return new ResponseEntity<>(categoryService.get(id),HttpStatus.OK);
    }

    @GetMapping("all")
    public ResponseEntity<List<Category>> getAll(){
        return new ResponseEntity<>(categoryService.getAll(),HttpStatus.OK);
    }

}
