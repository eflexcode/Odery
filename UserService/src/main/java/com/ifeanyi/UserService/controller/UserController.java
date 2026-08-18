package com.ifeanyi.UserService.controller;

import com.ifeanyi.UserService.entity.Inventory;
import com.ifeanyi.UserService.entity.User;
import com.ifeanyi.UserService.exception.NotFoundException;
import com.ifeanyi.UserService.model.InventoryModel;
import com.ifeanyi.UserService.model.UserModel;
import com.ifeanyi.UserService.service.UserService;
import com.ifeanyi.UserService.service.InventoryService;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.BeanUtils;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.Date;

@RestController
@RequiredArgsConstructor
@RequestMapping("/")
public class UserController {

    private final UserService userService;
    private final InventoryService inventoryService;

    @PostMapping("create")
    public ResponseEntity<User> create(@RequestBody UserModel userModel) {
        return new ResponseEntity<User>(userService.insert(userModel), HttpStatus.CREATED);
    }

    @PostMapping("create/admin")
    public ResponseEntity<User> createAdmin(@RequestBody UserModel userModel) {
        return new ResponseEntity<User>(userService.insertAdmin(userModel), HttpStatus.CREATED);
    }

    @PutMapping("{id}")
    public ResponseEntity<User> update(@RequestBody UserModel userModel, @PathVariable String id) throws NotFoundException {
        return new ResponseEntity<>(userService.update(userModel, id), HttpStatus.OK);
    }

    @GetMapping("{id}")
    public ResponseEntity<User> get(@PathVariable String id) throws NotFoundException {
        return new ResponseEntity<>(userService.get(id), HttpStatus.OK);
    }

    @DeleteMapping("{id}")
    public ResponseEntity<Void> delete(@PathVariable String id) {
        userService.delete(id);//on order service
        return new ResponseEntity<>(HttpStatus.OK);
    }

    @PostMapping("inv/add")
    public ResponseEntity<Inventory> addInventory(@RequestBody InventoryModel inventoryModel) {
        return new ResponseEntity<>(inventoryService.add(inventoryModel), HttpStatus.OK);
    }

    @DeleteMapping("inv/del/{order_id}")
    public ResponseEntity<Void> delInventory(@PathVariable("order_id") String orderId) {
        return new ResponseEntity<>(inventoryService.del(orderId), HttpStatus.OK);
    }

    @PostMapping("inv/get/{id}")
    public ResponseEntity<Inventory> getInventory(@PathVariable String id) {
        return new ResponseEntity<>(inventoryService.get(id), HttpStatus.OK);
    }

    @PostMapping("inv/get-all/{userId}")
    public ResponseEntity<Page<Inventory>> getAllUserInventory(@PathVariable String userId, Pageable pageable) {
        return new ResponseEntity<>(inventoryService.getAll(userId, pageable), HttpStatus.OK);
    }

}
