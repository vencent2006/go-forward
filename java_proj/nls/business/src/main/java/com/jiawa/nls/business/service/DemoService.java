package com.jiawa.nls.business.service;

import cn.hutool.core.bean.BeanUtil;
import com.jiawa.nls.business.domain.Demo;
import com.jiawa.nls.business.domain.DemoExample;
import com.jiawa.nls.business.mapper.DemoMapper;
import com.jiawa.nls.business.req.DemoQueryReq;
import com.jiawa.nls.business.resp.DemoQueryResp;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.List;

@Service
public class DemoService {
    @Resource
    private DemoMapper demoMapper;

    public int count() {
        return Math.toIntExact(demoMapper.countByExample(null));
    }

    public List<DemoQueryResp> query(DemoQueryReq req) {
        String mobile = req.getMobile();
        DemoExample demoExample = new DemoExample();
        DemoExample.Criteria criteria = demoExample.createCriteria();
        if (mobile != null) {
            criteria.andMobileEqualTo(mobile);
        }
        List<Demo> list = demoMapper.selectByExample(demoExample);
        return BeanUtil.copyToList(list, DemoQueryResp.class);
    }
}
