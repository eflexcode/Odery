package com.ifeanyi.UserService.service;

import com.ifeanyi.UserService.entity.User;
import com.ifeanyi.UserService.exception.NotFoundException;
import com.ifeanyi.UserService.model.UserModel;

public interface UserService {

    User insert(UserModel userModel);

    User update(UserModel userModel,String id) throws NotFoundException;

    User get(String id) throws NotFoundException;

    void delete(String id);

}
