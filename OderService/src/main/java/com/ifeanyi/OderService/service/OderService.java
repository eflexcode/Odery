package com.ifeanyi.OderService.service;

import com.ifeanyi.OderService.entity.Oder;
import com.ifeanyi.OderService.model.OderModel;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;

import java.util.Date;

public interface OderService {

    Oder create(OderModel oderModel);

    Oder update(OderModel oderModel);

    Page<Oder> get(String userId, String productId, Pageable pageable);

    Oder getById(String id);
}
