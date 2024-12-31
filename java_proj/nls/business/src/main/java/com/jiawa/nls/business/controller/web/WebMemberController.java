package com.jiawa.nls.business.controller.web;

import cn.hutool.crypto.digest.DigestUtil;
import com.jiawa.nls.business.req.MemberRegisterReq;
import com.jiawa.nls.business.resp.CommonResp;
import com.jiawa.nls.business.service.MemberService;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import javax.validation.Valid;

@RestController
@RequestMapping("/web/member")
public class WebMemberController {
    @Resource
    private MemberService memberService;


    /**
     * 注册
     *
     * @param req MemberRegisterReq
     * @return CommonResp
     */
    @PostMapping("/register")
    public CommonResp<String> register(@Valid @RequestBody MemberRegisterReq req) {
        req.setPassword(DigestUtil.md5Hex(req.getPassword()));// 密码加密
        memberService.register(req);
        return new CommonResp<>();
    }

}
