package com.lld.im.service.user.controller;

import com.lld.im.common.ResponseVO;
import com.lld.im.service.user.model.req.ImportUserReq;
import com.lld.im.service.user.service.ImUserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("v1/user")
public class ImUserController {

    @Autowired
    private ImUserService imUserService;


    // TODO 这里应该用post
    @RequestMapping("/importUser")
    public ResponseVO importUser(@RequestBody ImportUserReq req, Integer appId) {

        req.setAppId(appId);

        return imUserService.importUser(req);
    }
}
