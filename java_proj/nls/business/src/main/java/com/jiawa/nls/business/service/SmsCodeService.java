package com.jiawa.nls.business.service;

import cn.hutool.core.date.DateUtil;
import cn.hutool.core.util.IdUtil;
import cn.hutool.core.util.RandomUtil;
import com.jiawa.nls.business.domain.SmsCode;
import com.jiawa.nls.business.domain.SmsCodeExample;
import com.jiawa.nls.business.enums.SmsCodeStatusEnum;
import com.jiawa.nls.business.exception.BusinessException;
import com.jiawa.nls.business.exception.BusinessExceptionEnum;
import com.jiawa.nls.business.mapper.SmsCodeMapper;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.Date;

@Slf4j
@Service
public class SmsCodeService {
    @Resource
    private SmsCodeMapper smsCodeMapper;

    /**
     * 发送验证码
     * 校验：如果1分钟内有同手机号同用途发送记录，则报错：短信请求过于频繁
     *
     * @param mobile 手机号
     * @param use    用途
     */
    public void sendCode(String mobile, String use) {
        Date now = new Date();
        String code = RandomUtil.randomNumbers(6);

        log.info("当前时间：{}", new Date());
        log.info("1分钟前时间：{}", DateUtil.offsetMinute(new Date(), -1));

        // 校验：如果1分钟内有同手机号同用途发送记录，则报错：短信请求过于频繁
        SmsCodeExample smsCodeExample = new SmsCodeExample();
        SmsCodeExample.Criteria criteria = smsCodeExample.createCriteria();
        criteria.andMobileEqualTo(mobile)
                .andUseEqualTo(use)
                .andCreateAtGreaterThan(DateUtil.offsetMinute(new Date(), -1));
        long count = smsCodeMapper.countByExample(smsCodeExample);
        if (count > 0) {
            throw new BusinessException(BusinessExceptionEnum.SMS_CODE_TOO_FREQUENT);
        }


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
