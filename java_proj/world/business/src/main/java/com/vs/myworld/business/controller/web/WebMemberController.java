package com.vs.myworld.business.controller.web;

import cn.hutool.crypto.digest.DigestUtil;
import com.vs.myworld.business.enums.SmsCodeUseEnum;
import com.vs.myworld.business.req.MemberLoginReq;
import com.vs.myworld.business.req.MemberRegisterReq;
import com.vs.myworld.business.req.MemberResetReq;
import com.vs.myworld.business.resp.CommonResp;
import com.vs.myworld.business.resp.MemberLoginResp;
import com.vs.myworld.business.service.MemberService;
import com.vs.myworld.business.service.SmsCodeService;
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
    public CommonResp<Object> register(@Valid @RequestBody MemberRegisterReq req) {
        req.setPassword(DigestUtil.md5Hex(req.getPassword().toLowerCase()).toLowerCase());// 密码加密

        log.info("会员注册开始: {}", req.getMobile());
        smsCodeService.validCode(req.getMobile(), SmsCodeUseEnum.REGISTER.getCode(), req.getCode());
        log.info("注册验证码校验通过: {}", req.getMobile());

        memberService.register(req);
        return new CommonResp<>();
    }

    /**
     * 登录
     *
     * @param req MemberRegisterReq
     * @return CommonResp
     */
    @PostMapping("/login")
    public CommonResp<MemberLoginResp> login(@Valid @RequestBody MemberLoginReq req) {
        req.setPassword(DigestUtil.md5Hex(req.getPassword().toLowerCase()));// 密码加密

        log.info("会员登录开始: {}", req.getMobile());
        MemberLoginResp memberLoginResp = memberService.login(req);
        log.info("会员登录结束: {}", req.getMobile());

        return new CommonResp<>(memberLoginResp);
    }

    @PostMapping("/reset")
    public CommonResp<Object> reset(@Valid @RequestBody MemberResetReq req) {
        req.setPassword(DigestUtil.md5Hex(req.getPassword().toLowerCase()).toLowerCase());// 密码加密

        log.info("重置密码, 开始: {}", req.getMobile());
        smsCodeService.validCode(req.getMobile(), SmsCodeUseEnum.RESET.getCode(), req.getCode());
        log.info("重置密码, 验证码校验通过: {}", req.getMobile());

        memberService.reset(req);
        return new CommonResp<>();
    }

}
