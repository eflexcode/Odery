package com.ifeanyi.OderService.service.impl;

import com.ifeanyi.OderService.Repository.OrderRepository;
import com.ifeanyi.OderService.entity.Order;
import com.ifeanyi.OderService.exception.NotFoundException;
import com.ifeanyi.OderService.model.OrderModel;
import com.ifeanyi.OderService.service.OrderService;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.BeanUtils;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import java.util.Date;

@Service
@RequiredArgsConstructor
public class OrderServiceImpl implements OrderService {

    private final OrderRepository repository;

    @Override
    public Order create(OrderModel orderModel) {

        Order order = new Order();

        BeanUtils.copyProperties(orderModel, order);

        Date date = new Date();
        order.setCreatedAt(date);
        order.setUpdatedAt(date);

        return repository.save(order);
    }

    @Override
    public Order update(OrderModel orderModel, String id) throws NotFoundException {

        Order order = getById(id);

        BeanUtils.copyProperties(orderModel, order);

        Date date = new Date();
        order.setCreatedAt(date);
        order.setUpdatedAt(date);

        return repository.save(order);
    }

    @Override
    public Page<Order> get(String userId, String productId, Pageable pageable) {
        return repository.findAllByUserIdOrProductId(userId, productId, pageable);
    }

    @Override
    public Order getById(String id) throws NotFoundException {
        return repository.findById(id).orElseThrow(()-> new NotFoundException("no order found with id: "+id));
    }

}
