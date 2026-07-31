package com.ifeanyi.UserService.repository;

import com.ifeanyi.UserService.entity.Inventory;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface InventoryRepository extends JpaRepository<Inventory,String> {
    Page<Inventory> findByUserId(String userId, Pageable pageable);
}
