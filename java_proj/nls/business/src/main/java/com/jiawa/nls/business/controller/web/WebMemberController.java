package com.jiawa.nls.business.controller.web;

import cn.hutool.crypto.digest.DigestUtil;
import com.jiawa.nls.business.enums.SmsCodeUseEnum;
import com.jiawa.nls.business.req.MemberRegisterReq;
import com.jiawa.nls.business.resp.CommonResp;
import com.jiawa.nls.business.service.MemberService;
import com.jiawa.nls.business.service.SmsCodeService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import javax.validation.Valid;

@Slf4j
@RestController
@RequestMapping("/web/member")
public class WebMemberController {
    @Resource
    private MemberService memberService;
    @Resource
    private SmsCodeService smsCodeService;


    /**
     * 注册
     *
     * @param req MemberRegisterReq
     * @return CommonResp
     */
    @PostMapping("/register")
    public CommonResp<String> register(@Valid @RequestBody MemberRegisterReq req) {
        req.setPassword(DigestUtil.md5Hex(req.getPassword()));// 密码加密
        
        log.info("会员注册开始: {}", req.getMobile());
        smsCodeService.validCode(req.getMobile(), SmsCodeUseEnum.REGISTER.getCode(), req.getCode());
        log.info("注册验证码校验通过: {}", req.getMobile());

        memberService.register(req);
        return new CommonResp<>();
    }

}
