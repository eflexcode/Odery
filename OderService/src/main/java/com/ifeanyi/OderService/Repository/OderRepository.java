package com.ifeanyi.OderService.Repository;

import com.ifeanyi.OderService.entity.Oder;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface OderRepository extends MongoRepository<Oder,String> {



}
