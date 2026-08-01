package com.ifeanyi.OderService.Messaging;

import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Component;

@Component
public class OrderMessagingConsumer {

    @RabbitListener(queues = "order")
    public void receiveMessage(String message){
        //       Handle the received message here
        System.out.println("Received message: " + message);
    }

//    @RabbitListener(queues = "payment")
//    public void receiveMessagePayment(String message){
//        //       Handle the received message here
//        System.out.println("Received message: " + message);
//    }

}
