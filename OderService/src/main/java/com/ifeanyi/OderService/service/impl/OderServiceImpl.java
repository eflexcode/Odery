package com.ifeanyi.OderService.service.impl;

import com.ifeanyi.OderService.Repository.OderRepository;
import com.ifeanyi.OderService.entity.Oder;
import com.ifeanyi.OderService.model.OderModel;
import com.ifeanyi.OderService.service.OderService;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.BeanUtils;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.bind.annotation.PostMapping;

import java.util.Date;

@Service
@RequiredArgsConstructor
public class OderServiceImpl implements OderService {

    private final OderRepository repository;

    @Override
    public Oder create(OderModel oderModel) {

        Oder oder = new Oder();

        BeanUtils.copyProperties(oderModel,oder);

        Date date = new Date();
        oder.setCreatedAt(date);
        oder.setUpdatedAt(date);

        return repository.save(oder);
    }

    @Override
    public Oder update(OderModel oderModel) {
        return null;
    }

    @Override
    public Page<Oder> get(String userId, String productId, Pageable pageable) {
        return null;
    }

    @Override
    public Oder getById(String id) {
        return repository.findById(id).orElseThrow(()-> new NotFound);
    }



}
