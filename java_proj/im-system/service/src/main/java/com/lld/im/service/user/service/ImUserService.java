package com.lld.im.service.user.service;

import com.lld.im.common.ResponseVO;
import com.lld.im.service.user.model.req.ImportUserReq;

public interface ImUserService {

    /**
     * 导入用户
     * @param req ImportUser
     * @return ResponseVO
     */
    public ResponseVO importUser(ImportUserReq req);
}
