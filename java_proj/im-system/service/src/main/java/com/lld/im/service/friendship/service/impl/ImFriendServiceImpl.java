package com.lld.im.service.friendship.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.lld.im.common.ResponseVO;
import com.lld.im.common.enums.FriendShipStatusEnum;
import com.lld.im.service.friendship.dao.ImFriendShipEntity;
import com.lld.im.service.friendship.dao.mapper.ImFriendShipMapper;
import com.lld.im.service.friendship.model.req.AddFriendReq;
import com.lld.im.service.friendship.model.req.FriendDto;
import com.lld.im.service.friendship.model.req.ImportFriendShipReq;
import com.lld.im.service.friendship.model.resp.ImportFriendShipResp;
import com.lld.im.service.friendship.service.ImFriendService;
import com.lld.im.service.user.dao.ImUserDataEntity;
import com.lld.im.service.user.service.ImUserService;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.ArrayList;
import java.util.List;

@Service
public class ImFriendServiceImpl implements ImFriendService {

    @Autowired
    private ImFriendShipMapper imFriendShipMapper;

    @Autowired
    private ImUserService imUserService;


    @Override
    public ResponseVO importFriendShip(ImportFriendShipReq req) {

        // 检查数量
        if (req.getFriendItem().size() > 100) {
            // TODO 返回数量过多
        }

        List<String> successId = new ArrayList<>();
        List<String> errorId = new ArrayList<>();


        for (ImportFriendShipReq.ImportFriendDto dto : req.getFriendItem()) {

            ImFriendShipEntity entity = new ImFriendShipEntity();
            BeanUtils.copyProperties(dto, entity);
            entity.setAppId(req.getAppId());
            entity.setFromId(req.getFromId());
            try {
                int insert = imFriendShipMapper.insert(entity);
                if (1 == insert) {
                    successId.add(dto.getToId());
                } else {
                    errorId.add(dto.getToId());
                }
            } catch (Exception e) {
                e.printStackTrace();
                errorId.add(dto.getToId());
            }
        }


        // resp
        ImportFriendShipResp resp = new ImportFriendShipResp();
        resp.setSuccessId(successId);
        resp.setErrorId(errorId);


        return ResponseVO.successResponse(resp);
    }


    @Override
    public ResponseVO AddFriend(AddFriendReq req) {


        // 判断from和to，是否存在
        ResponseVO<ImUserDataEntity> fromInfo = imUserService.getSingleUserInfo(req.getFromId(), req.getAppId());
        if (!fromInfo.isOk()) {
            return fromInfo;
        }

        ResponseVO<ImUserDataEntity> toInfo = imUserService.getSingleUserInfo(req.getToItem().getToId(), req.getAppId());
        if (!fromInfo.isOk()) {
            return toInfo;
        }

        return null;
    }


    @Transactional
    public ResponseVO doAddFriend(String fromId, FriendDto dto, Integer appId) {

        // A - B
        // Friend表插入 A 和 B 两条记录
        // 查询是否有记录存在
        // 如果存在则判断状态，如果已经添加，则提示已添加；如果是未添加，则修改状态

        QueryWrapper<ImFriendShipEntity> query = new QueryWrapper<>();
        query.eq("app_id", appId);
        query.eq("from_id", fromId);
        query.eq("to_id", dto.getToId());
        ImFriendShipEntity entity = imFriendShipMapper.selectOne(query);
        if (null == entity) {
            // 走添加逻辑
            entity = new ImFriendShipEntity();
            entity.setFromId(fromId);
            BeanUtils.copyProperties(dto, entity);
            entity.setStatus(FriendShipStatusEnum.FRIEND_STATUS_NORMAL.getCode());
            entity.setCreateTime(System.currentTimeMillis());

            imFriendShipMapper.insert(entity);
        } else {
            // 走修改逻辑
        }


        return ResponseVO.successResponse();
    }
}
