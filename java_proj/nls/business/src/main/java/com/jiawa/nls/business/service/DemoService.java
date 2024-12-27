package com.jiawa.nls.business.service;

import com.jiawa.nls.business.domain.Demo;
import com.jiawa.nls.business.domain.DemoExample;
import com.jiawa.nls.business.mapper.DemoMapper;
import com.jiawa.nls.business.req.DemoQueryReq;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.List;

@Service
public class DemoService {
    @Resource
    private DemoMapper demoMapper;

    public int count(){
        return Math.toIntExact(demoMapper.countByExample(null));
    }

    public List<Demo> query(DemoQueryReq req){
        String mobile = req.getMobile();
        DemoExample demoExample = new DemoExample();
        DemoExample.Criteria criteria = demoExample.createCriteria();
        if(mobile != null){
            criteria.andMobileEqualTo(mobile);
        }
        return demoMapper.selectByExample(demoExample);
    }
}
