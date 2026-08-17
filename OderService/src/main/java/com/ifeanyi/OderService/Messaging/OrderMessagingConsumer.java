package com.ifeanyi.OderService.Messaging;

import com.ifeanyi.OderService.entity.OrderStatus;
import com.ifeanyi.OderService.exception.NotFoundException;
import com.ifeanyi.OderService.service.OrderService;
import com.ifeanyi.OderService.service.OtherService.model.Inventory;
import com.ifeanyi.OderService.service.OtherService.model.Payment;
import com.ifeanyi.OderService.service.OtherService.model.Product;
import com.ifeanyi.OderService.util.Util;
import lombok.RequiredArgsConstructor;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.http.*;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;
import tools.jackson.databind.ObjectMapper;

import java.util.Objects;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;

@Component
@RequiredArgsConstructor
public class OrderMessagingConsumer {

    private final OrderService orderService;
    private final RestTemplate restTemplate;

    @RabbitListener(queues = "order")
    public void receiveMessage(String message) throws NotFoundException {
        System.out.println("Received message: " + message);
        String[] typeAndBody = message.split(" rabbitmqIfy ");
        String[] type = typeAndBody[0].split(":");

        ScheduledExecutorService scheduled = Executors.newSingleThreadScheduledExecutor();

        if (type[0].equals("Payment")) {
            ObjectMapper objectMapper = new ObjectMapper();
            Payment payment = objectMapper.readValue(typeAndBody[1], Payment.class);
            //refund,paid,-
            if (payment.getType().equals("paid")) {
                //update order status add to inventory
                orderService.updateStatus(OrderStatus.DONE, payment.getOrderId());
                addInventory(payment.getUserId(), payment.getProductId(),payment.getItemCount());
            } else if (payment.getType().equals("refund")) {
                //update order status remove from inventory
                //order already canceled before request for refund a just update inventory
                deleteInventory(payment.getOrderId());
            } else if (payment.getType().equals("-")) {
                //payment failed:  eg insufficient funds might delete order in case of price change TODO use delay of one day auto delete
                Runnable task = () ->{
                    deleteInventory(payment.getOrderId());
                    scheduled.shutdown();
                };

                scheduled.schedule(task,1, TimeUnit.DAYS);

            }
        } else if (type[0].equals("Order")) {

        }

    }

    private boolean addInventory(String userId, String productId,int count) {

        String endpoint = "inv/add";
        Inventory inventory = new Inventory();
        inventory.setUserId(userId);
        inventory.setProductId(productId);
        inventory.setCount(count);

        HttpHeaders httpHeaders = new HttpHeaders();
        HttpEntity<Inventory> httpEntity = new HttpEntity<>(inventory, httpHeaders);
        ResponseEntity<Inventory> responseEntity = restTemplate.exchange(Util.ProductServiceBaseUrl + endpoint, HttpMethod.POST, httpEntity, Inventory.class);

        if (responseEntity.getStatusCode() != HttpStatus.OK) {
            return false;
        }

        return responseEntity.getBody() != null;
    }

    private boolean deleteInventory(String orderId) {

        String endpoint = "inv/del/"+orderId;

        HttpHeaders httpHeaders = new HttpHeaders();
        HttpEntity<Inventory> httpEntity = new HttpEntity<>(httpHeaders);
        ResponseEntity<Void> responseEntity = restTemplate.exchange(Util.ProductServiceBaseUrl + endpoint, HttpMethod.DELETE, httpEntity, Void.class);

        return responseEntity.getStatusCode() == HttpStatus.OK;
    }

//    @RabbitListener(queues = "payment")
//    public void receiveMessagePayment(String message){
//        //       Handle the received message here
//        System.out.println("Received message: " + message);
//    }

}
