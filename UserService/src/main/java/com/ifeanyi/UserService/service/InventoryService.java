package com.ifeanyi.UserService.service;

import com.ifeanyi.UserService.entity.Inventory;
import com.ifeanyi.UserService.model.InventoryModel;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

public interface InventoryService {

    Inventory add (InventoryModel model);
    Inventory get (String id);
    Page<Inventory> getAll(String userId, Pageable pageable);

}
