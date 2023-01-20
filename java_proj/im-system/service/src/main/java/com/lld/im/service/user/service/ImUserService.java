package com.lld.im.service.user.service;

import com.lld.im.common.ResponseVO;
import com.lld.im.service.user.dao.ImUserDataEntity;
import com.lld.im.service.user.model.req.DeleteUserReq;
import com.lld.im.service.user.model.req.GetUserInfoReq;
import com.lld.im.service.user.model.req.ImportUserReq;
import com.lld.im.service.user.model.req.ModifyUserInfoReq;
import com.lld.im.service.user.model.resp.GetUserInfoResp;

public interface ImUserService {

    /**
     * 导入用户
     * @param req ImportUser
     * @return ResponseVO
     */
    public ResponseVO importUser(ImportUserReq req);

    /**
     * 批量获取用户信息
     * @param req GetUserInfoReq(包含一个userId的list)
     * @return GetUserInfoResp
     */
    public ResponseVO<GetUserInfoResp> getUserInfo(GetUserInfoReq req);


    /**
     * 获取单个用户信息
     * @param userId 用户id
     * @param appId 业务id
     * @return ImUserDataEntity
     */
    public ResponseVO<ImUserDataEntity> getSingleUserInfo(String userId , Integer appId);

    /**
     * 设置为删除状态
     * @param req DeleteUserReq(包含为一个列表)
     * @return ImportUserResp
     */
    public ResponseVO deleteUser(DeleteUserReq req);

    /**
     * 修改某一个用户的信息
     * @param req ModifyUserInfoReq
     * @return
     */
    public ResponseVO modifyUserInfo(ModifyUserInfoReq req);
}
