package com.jiawa.nls.business.service;

import cn.hutool.core.util.IdUtil;
import cn.hutool.core.util.RandomUtil;
import com.jiawa.nls.business.domain.SmsCode;
import com.jiawa.nls.business.enums.SmsCodeStatusEnum;
import com.jiawa.nls.business.mapper.SmsCodeMapper;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.Date;

@Service
public class SmsCodeService {
    @Resource
    private SmsCodeMapper smsCodeMapper;

    /**
     * 发送验证码
     *
     * @param mobile 手机号
     * @param use    用途
     */
    public void sendCode(String mobile, String use) {
        Date now = new Date();
        String code = RandomUtil.randomNumbers(6);

        // 保存数据库
        SmsCode smsCode = new SmsCode();
        smsCode.setId(IdUtil.getSnowflakeNextId());
        smsCode.setMobile(mobile);
        smsCode.setCode(code);
        smsCode.setUse(use);
        smsCode.setStatus(SmsCodeStatusEnum.UNUSED.getCode());
        smsCode.setCreateAt(now);
        smsCode.setUpdateAt(now);
        smsCodeMapper.insert(smsCode);

        // 发送短信
    }
}
