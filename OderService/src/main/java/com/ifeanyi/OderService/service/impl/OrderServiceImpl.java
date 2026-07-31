package com.ifeanyi.OderService.service.impl;

import com.ifeanyi.OderService.Messaging.OrderMessagingProducer;
import com.ifeanyi.OderService.Repository.OrderRepository;
import com.ifeanyi.OderService.entity.Order;
import com.ifeanyi.OderService.entity.OrderStatus;
import com.ifeanyi.OderService.exception.BadRequestException;
import com.ifeanyi.OderService.exception.NotFoundException;
import com.ifeanyi.OderService.model.OrderModel;
import com.ifeanyi.OderService.service.OrderService;
import com.ifeanyi.OderService.service.OtherService.model.Product;
import com.ifeanyi.OderService.util.Util;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.BeanUtils;
import org.springframework.boot.json.JsonParser;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.*;
import org.springframework.stereotype.Service;
import org.springframework.web.client.HttpClientErrorException;
import org.springframework.web.client.RestTemplate;
//import org.springframework.web.reactive.result.method.annotation.ResponseEntityExceptionHandler;
import org.springframework.web.server.ResponseStatusException;
import tools.jackson.databind.ObjectMapper;

import java.time.Instant;
import java.time.InstantSource;
import java.util.Date;

@Service
@RequiredArgsConstructor
public class OrderServiceImpl implements OrderService {

    private final OrderRepository repository;
    private final OrderMessagingProducer messagingProducer;
    private final RestTemplate restTemplate;

    @Override
    public Order create(OrderModel orderModel) throws BadRequestException {

        Order order = new Order();
        Product product;

        //call product to validate product id
        try {
            product = validateProduct(orderModel.getProductId());
        } catch (HttpClientErrorException httpClientErrorException) {
            throw new ResponseStatusException(HttpStatus.BAD_REQUEST, "Invalid product id");
        }

        if (product == null) {
            throw new BadRequestException("Invalid product id");
        }

        BeanUtils.copyProperties(orderModel, order);
        order.setStatus(OrderStatus.SUBMITTED);// changed to done on message received from RabbitMQ
        int amount = product.getPrice() * orderModel.getCount();
        order.setAmount(amount);

        Date date = new Date();
        order.setCreatedAt(date);
        order.setUpdatedAt(date);

        Order savedOrder = repository.save(order);

        ObjectMapper objectMapper = new ObjectMapper();

        messagingProducer.sendMessage("QueueType:Order rabbitmqIfy "+objectMapper.writeValueAsString(savedOrder));// rabbitmq message

        return order;
    }

    @Override
    public Order update(OrderModel orderModel, String id) throws NotFoundException {

        Order order = getById(id);

        BeanUtils.copyProperties(orderModel, order);

        Date date = new Date();
        order.setUpdatedAt(date);

        return repository.save(order);
    }

    @Override
    public Order cancel(String id) throws NotFoundException {
        Order order = getById(id);

        int twoDays = 60 * 60 * 24 * 2;
        Date dateTwoDaysFromOrderPlaced = new Date(order.getCreatedAt().getTime() + twoDays);

        if (order.getCreatedAt().after(dateTwoDaysFromOrderPlaced)) {
            throw new ResponseStatusException(HttpStatus.UNPROCESSABLE_CONTENT, "Cannot cancel order after 2 days");
        }

        order.setStatus(OrderStatus.CANCELED);
        Date date = new Date();
        order.setUpdatedAt(date);

        Order savedOrder = repository.save(order);

        ObjectMapper objectMapper = new ObjectMapper();

        messagingProducer.sendMessage(objectMapper.writeValueAsString(savedOrder));// rabbitmq message

        return repository.save(savedOrder);
    }

    @Override
    public Page<Order> get(String userId, OrderStatus status, Pageable pageable) {
        return repository.findAllByUserId(userId, pageable);
    }

    @Override
    public Order getById(String id) throws NotFoundException {
        return repository.findById(id).orElseThrow(() -> new NotFoundException("no order found with id: " + id));
    }

    public Product validateProduct(String productId) {
        String endpoint = "" + productId;
        HttpHeaders httpHeaders = new HttpHeaders();

        HttpEntity<Product> httpEntity = new HttpEntity<>(httpHeaders);
        ResponseEntity<Product> responseEntity = restTemplate.exchange(Util.ProductServiceBaseUrl + endpoint, HttpMethod.GET, httpEntity, Product.class);

        if (responseEntity.getStatusCode() != HttpStatus.OK) {
            return null;
        }

        if (responseEntity.getBody() == null) {
            return null;
        }

        return responseEntity.getBody();

//        String gottenId = responseEntity.getBody().getId();
//        if (gottenId.equals(productId)) {
//            return false;
//        }
//        return true;
    }

}

