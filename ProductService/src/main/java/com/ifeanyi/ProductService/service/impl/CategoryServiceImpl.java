package com.ifeanyi.ProductService.service.impl;

import com.ifeanyi.ProductService.entity.Category;
import com.ifeanyi.ProductService.exception.NotFoundExceptionHandler;
import com.ifeanyi.ProductService.model.CategoryModel;
import com.ifeanyi.ProductService.repository.CategoryRepository;
import com.ifeanyi.ProductService.service.CategoryService;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.BeanUtils;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.List;

@Service
@RequiredArgsConstructor
public class CategoryServiceImpl implements CategoryService {

    private final CategoryRepository repository;

    @Override
    public Category create(CategoryModel model) {

        Category category = new Category();
        BeanUtils.copyProperties(model, category);
        Date date = new Date();
        category.setCreatedAt(date);
        category.setUpdatedAt(date);

        return repository.save(category);
    }

    @Override
    public Category update(CategoryModel model, String id) throws NotFoundExceptionHandler {

        Category category = get(id);
        BeanUtils.copyProperties(model, category);
        Date date = new Date();
        category.setUpdatedAt(date);

        return repository.save(category);
    }

    @Override
    public Category get(String id) throws NotFoundExceptionHandler {
        return repository.findById(id).orElseThrow(() -> new NotFoundExceptionHandler("No category found with id: " + id));
    }

    @Override
    public List<Category> getAll() {
        return repository.findAll();
    }
}
