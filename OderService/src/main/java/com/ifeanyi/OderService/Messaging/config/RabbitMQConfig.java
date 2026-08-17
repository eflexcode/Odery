package com.ifeanyi.OderService.Messaging.config;

import org.springframework.amqp.core.*;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class RabbitMQConfig {

    @Bean
    public Queue queue() {
        return new Queue("order", true);
    }

//    @Bean
//    public Exchange exchange() {
//        return new DirectExchange("exchange-order");
//    }
@Bean
    public FanoutExchange paymentExchange() {
        return new FanoutExchange("payment.exchange");
    }
//
//     dont listen to your self
//    @Bean
//    public Binding binding(Queue queue, Exchange exchange) {
//        return BindingBuilder.bind(queue)
//                .to(exchange)
//                .with("order-routing-key")
//                .noargs();
//    }

    @Bean
    public Binding paymentBinding(Queue queue) {
        return BindingBuilder.bind(queue)
                .to(paymentExchange());
    }
}
