package com.example.demo.service.impl;

import com.example.demo.service.SparrowGodService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class SparrowGodServiceImpl implements SparrowGodService {

    @Override
    public List<Integer> claimRoom() {
        return null;
    }

    @Override
    public boolean joinRoom() {
        return false;
    }

    @Override
    public boolean sit() {
        return false;
    }
}
