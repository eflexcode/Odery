package com.ifeanyi.OderService.service;

import com.ifeanyi.OderService.entity.Order;
import com.ifeanyi.OderService.exception.BadRequestException;
import com.ifeanyi.OderService.exception.NotFoundException;
import com.ifeanyi.OderService.model.OrderModel;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;

public interface OrderService {

    Order create(OrderModel orderModel) throws BadRequestException;

    Order update(OrderModel orderModel, String id) throws NotFoundException;

    Order cancel(String id) throws NotFoundException;

    Page<Order> get(String userId, String productId, Pageable pageable);

    Order getById(String id) throws NotFoundException;

}
