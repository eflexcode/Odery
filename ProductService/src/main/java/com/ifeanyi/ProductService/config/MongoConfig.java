package com.ifeanyi.ProductService.config;

import com.ifeanyi.ProductService.util.Util;
import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.mongodb.config.AbstractMongoClientConfiguration;

@Configuration
public class MongoConfig extends AbstractMongoClientConfiguration {

    @Override
    protected String getDatabaseName() {
        return Util.DATABASE_NAME;
    }

    @Override
    public MongoClient mongoClient() {
        return MongoClients.create(Util.MONGODB_URL);
    }
}
