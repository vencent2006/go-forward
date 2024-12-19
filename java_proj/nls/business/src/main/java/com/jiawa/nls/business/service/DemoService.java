package com.jiawa.nls.business.service;

import com.jiawa.nls.business.mapper.DemoMapper;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;

@Service
public class DemoService {
    @Resource
    private DemoMapper demoMapper;

    public int count(){
        return demoMapper.count();
    }
}
