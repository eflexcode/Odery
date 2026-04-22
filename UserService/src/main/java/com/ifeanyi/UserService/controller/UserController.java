package com.ifeanyi.UserService.controller;

import com.ifeanyi.UserService.entity.User;
import com.ifeanyi.UserService.exception.NotFoundException;
import com.ifeanyi.UserService.model.StandardRes;
import com.ifeanyi.UserService.model.UserModel;
import com.ifeanyi.UserService.service.UserService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequiredArgsConstructor
@RequestMapping("/")
public class UserController {

    private final UserService userService;

    @PostMapping("create")
    public ResponseEntity<User> create(@RequestBody UserModel userModel) {
        return new ResponseEntity<User>(userService.insert(userModel), HttpStatus.CREATED);
    }

    @PutMapping("update/{id}")
    public ResponseEntity<User> update(@RequestBody UserModel userModel, @PathVariable String id) throws NotFoundException {
        return new ResponseEntity<>(userService.update(userModel, id), HttpStatus.OK);
    }

    @GetMapping("{id}")
    public ResponseEntity<User> get(@PathVariable String id) throws NotFoundException {
        return new ResponseEntity<>(userService.get(id), HttpStatus.OK);
    }

    @DeleteMapping("{id}")
    public ResponseEntity<Void> delete(@PathVariable String id) {
        userService.delete(id);
        return new ResponseEntity<>(HttpStatus.OK);
    }

}
