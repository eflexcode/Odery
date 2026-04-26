package com.ifeanyi.ProductService.service;

import com.ifeanyi.ProductService.entity.Category;
import com.ifeanyi.ProductService.exception.NotFoundExceptionHandler;
import com.ifeanyi.ProductService.model.CategoryModel;

import java.util.List;

public interface CategoryService {

    Category create(CategoryModel model);
    Category update(CategoryModel model,String id) throws NotFoundExceptionHandler;
    Category get(String id) throws NotFoundExceptionHandler;
    List<Category> getAll();

}
