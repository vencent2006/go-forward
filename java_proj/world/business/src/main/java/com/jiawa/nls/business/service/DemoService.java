package com.jiawa.nls.business.service;

import cn.hutool.core.bean.BeanUtil;
import cn.hutool.core.util.StrUtil;
import com.jiawa.nls.business.domain.Demo;
import com.jiawa.nls.business.domain.DemoExample;
import com.jiawa.nls.business.exception.BusinessException;
import com.jiawa.nls.business.exception.BusinessExceptionEnum;
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
        // if (mobile != null) {
        //     criteria.andMobileEqualTo(mobile);
        // }
        // StrUtil.isBlank() isBlank方法会判断字符串是否为空或仅包含空白字符‌。如果字符串为null、空字符串("")或仅包含空白字符（如空格、制表符、换行符等），则返回true；否则返回false。
        // StrUtil.isEmpty() isEmpty方法仅检查字符串是否为null或空字符串‌，不考虑空白字符。如果字符串为null或空字符串("")，则返回true；
        if (StrUtil.isBlank(mobile)) {
            throw new BusinessException(BusinessExceptionEnum.DEMO_MOBILE_NOT_NULL);
        }
        criteria.andMobileEqualTo(mobile);
        List<Demo> list = demoMapper.selectByExample(demoExample);
        return BeanUtil.copyToList(list, DemoQueryResp.class);
    }
}
