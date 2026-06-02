package com.ifeanyi.UserService.service.impl;

import com.ifeanyi.UserService.entity.User;
import com.ifeanyi.UserService.exception.NotFoundException;
import com.ifeanyi.UserService.model.UserModel;
import com.ifeanyi.UserService.repository.UserRepository;
import com.ifeanyi.UserService.service.UserService;
import lombok.AllArgsConstructor;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.BeanUtils;
import org.springframework.stereotype.Service;

import java.util.Date;

@Service
@RequiredArgsConstructor
public class UserServiceImpl implements UserService {

    private final UserRepository userRepository;

    @Override
    public User insert(UserModel userModel) {

        User user = new User();
        BeanUtils.copyProperties(userModel, user);
        user.setCreatedAt(new Date());
        user.setUpdatedAt(new Date());

        return userRepository.save(user);
    }

    @Override
    public User update(UserModel userModel,String id) throws NotFoundException {

        User user = get(id);
        BeanUtils.copyProperties(userModel, user);
        user.setUpdatedAt(new Date());

        return userRepository.save(user);
    }

    @Override
    public User get(String id) throws NotFoundException {
        return userRepository.findById(id).orElseThrow(() -> new NotFoundException("No user found with id: "+id));
    }

    @Override
    public void delete(String id) {
        userRepository.deleteById(id);
    }

}
