package com.vs.myworld.business.controller.web;

import com.vs.myworld.business.req.SmsCodeRegisterReq;
import com.vs.myworld.business.req.SmsCodeResetReq;
import com.vs.myworld.business.resp.CommonResp;
import com.vs.myworld.business.service.SmsCodeService;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import javax.validation.Valid;

@RestController
@RequestMapping("/web/sms-code")
public class WebSmsCodeController {
    @Resource
    private SmsCodeService smsCodeService;

    /**
     * 为注册发送短信
     *
     * @param req RegisterSmsCodeReq
     * @return null
     */
    @PostMapping("/send-for-register")
    public CommonResp<Object> sendForRegister(@Valid @RequestBody SmsCodeRegisterReq req) {
        smsCodeService.sendCodeForRegister(req.getMobile());
        return new CommonResp<>();
    }


    @PostMapping("/send-for-reset")
    public CommonResp<Object> sendForReset(@Valid @RequestBody SmsCodeResetReq req) {
        smsCodeService.sendCodeForReset(req.getMobile());
        return new CommonResp<>();
    }
}
