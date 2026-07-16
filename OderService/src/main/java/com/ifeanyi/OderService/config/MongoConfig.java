package com.ifeanyi.OderService.config;

import com.ifeanyi.OderService.util.Util;
import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.mongodb.config.AbstractMongoClientConfiguration;

@Configuration
public class MongoConfig extends AbstractMongoClientConfiguration {
    @Override
    protected String getDatabaseName() {
        return Util.DatabaseName;
    }

    @Override
    public MongoClient mongoClient() {
        return MongoClients.create(Util.MongoDBUrl);
    }
}
