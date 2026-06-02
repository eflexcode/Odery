package com.ifeanyi.OderService.Repository;

import com.ifeanyi.OderService.entity.Order;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface OrderRepository extends MongoRepository<Order, String> {
    Page<Order> findAllByUserIdOrProductId(String userId, String productId, Pageable pageable);
}
