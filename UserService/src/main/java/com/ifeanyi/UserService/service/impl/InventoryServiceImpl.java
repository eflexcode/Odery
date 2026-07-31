package com.ifeanyi.UserService.service.impl;

import com.ifeanyi.UserService.entity.Inventory;
import com.ifeanyi.UserService.model.InventoryModel;
import com.ifeanyi.UserService.repository.InventoryRepository;
import com.ifeanyi.UserService.service.InventoryService;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.BeanUtils;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import org.springframework.web.server.ResponseStatusException;
import org.springframework.web.server.handler.ResponseStatusExceptionHandler;

import java.util.Date;

@Service
@RequiredArgsConstructor
public class InventoryServiceImpl implements InventoryService {

    private final InventoryRepository inventoryRepository;

    @Override
    public Inventory add(InventoryModel model) {
        Inventory inventory = new Inventory();
        BeanUtils.copyProperties(model,inventory);
        inventory.setCreatedAt(new Date());
        return inventoryRepository.save(inventory);
    }

    @Override
    public Inventory get(String id) {
        return inventoryRepository.findById(id).orElseThrow(() ->new ResponseStatusException(HttpStatus.NOT_FOUND,"No inventory found with id:: "+id));
    }

    @Override
    public Page<Inventory> getAll(String userId, Pageable pageable) {
        return inventoryRepository.findByUserId(userId, pageable);
    }

}
