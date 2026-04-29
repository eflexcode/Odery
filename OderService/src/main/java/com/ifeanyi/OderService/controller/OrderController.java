package com.ifeanyi.OderService.controller;

import com.ifeanyi.OderService.entity.Order;
import com.ifeanyi.OderService.exception.NotFoundException;
import com.ifeanyi.OderService.model.OrderModel;
import com.ifeanyi.OderService.service.OrderService;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/")
@RequiredArgsConstructor
public class OrderController {

    private final OrderService orderService;

    @PostMapping("create")
    public ResponseEntity<Order> create(@RequestBody OrderModel orderModel) {
        return new ResponseEntity<>(orderService.create(orderModel), HttpStatus.CREATED);
    }

    @PutMapping("update/{id}")
    public ResponseEntity<Order> update(@RequestBody OrderModel orderModel, @PathVariable String id) throws NotFoundException {
        return new ResponseEntity<>(orderService.update(orderModel, id), HttpStatus.OK);
    }

    @GetMapping("get")
    @ResponseStatus(HttpStatus.OK)
    public Page<Order> get(@PathVariable String userId, @PathVariable String productId, Pageable pageable) {
        return orderService.get(userId, productId, pageable);
    }

    @GetMapping("getById")
    @ResponseStatus(HttpStatus.OK)
    public Order getById(String id) throws NotFoundException {
        return orderService.getById(id);
    }

}
