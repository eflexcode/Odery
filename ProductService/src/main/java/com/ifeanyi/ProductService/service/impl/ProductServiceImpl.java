package com.ifeanyi.ProductService.service.impl;

import com.ifeanyi.ProductService.entity.Category;
import com.ifeanyi.ProductService.entity.Product;
import com.ifeanyi.ProductService.exception.NotFoundExceptionHandler;
import com.ifeanyi.ProductService.model.ProductModel;
import com.ifeanyi.ProductService.model.StandardResponse;
import com.ifeanyi.ProductService.repository.ProductRepository;
import com.ifeanyi.ProductService.service.CategoryService;
import com.ifeanyi.ProductService.service.ProductService;
import com.ifeanyi.ProductService.service.impl.OtherServices.model.Role;
import com.ifeanyi.ProductService.service.impl.OtherServices.model.User;
import com.ifeanyi.ProductService.util.Util;
import com.ifeanyi.ProductService.service.impl.OtherServices.User.UserService;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.BeanUtils;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.*;
import org.springframework.stereotype.Service;
import org.springframework.web.client.HttpClientErrorException;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.multipart.MultipartFile;
import org.springframework.web.server.ResponseStatusException;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

import java.io.File;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Date;
import java.util.Optional;

@Service
@RequiredArgsConstructor
public class ProductServiceImpl implements ProductService {

    private final ProductRepository repository;
    private final CategoryService categoryService;
    private final UserService userService;

    @Override
    public Product create(MultipartFile file_img, ProductModel productModel) throws IOException {

        User user = null;
        Category category;

        //validations
        try {
            category = categoryService.get(productModel.getCategoryId());
        } catch (NotFoundExceptionHandler e) {
            throw new ResponseStatusException(HttpStatus.NOT_FOUND, "Invalid category id");
        }
        try {
            user = userService.getUserFromUserService(productModel.getUserId());
        } catch (HttpClientErrorException httpClientErrorException) {
            throw new ResponseStatusException(HttpStatus.BAD_REQUEST, "Invalid user id");
        }

        if (user == null) {
            throw new ResponseStatusException(HttpStatus.BAD_REQUEST, "User id is no valid");
        }

        if (user.getRole() != Role.ADMIN) {
            throw new ResponseStatusException(HttpStatus.UNAUTHORIZED, "User role cannot perform this activity");
        }

        //upload file s3 or disk save
        String imageDownloadUrl = uploadSlashSaveFile(file_img);

        Product product = new Product();
        BeanUtils.copyProperties(productModel, product);
        product.setProductImgUrl(imageDownloadUrl);
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

    @Override
    public byte[] getProductImg(String fileName) throws IOException {
        //for s3
        //byte[] bytes = s3.getObject(GetObjectRequest.builder().bucket(buketName).key(key).build()).readAllBytes();

        return Files.readAllBytes(Path.of(Util.FILE_DIR + fileName));
    }

    public String uploadSlashSaveFile(MultipartFile img_file) throws IOException {

        String fileName = System.currentTimeMillis() + img_file.getOriginalFilename();
        File file = new File(Util.FILE_DIR + fileName);

        img_file.transferTo(file);
        //        for aws s3
//        s3.putObject(PutObjectRequest.builder()
//                        .bucket(buketName)
//                        .key(key)
//                        .acl("public-read")
//                        .build(),
//                RequestBody.fromBytes(file.getBytes()));
//        downloadUrl = Util.endpoint + "/" + buketName + "/" + key;

        return "http://localhost:8092/img/" + fileName;
    }


}
