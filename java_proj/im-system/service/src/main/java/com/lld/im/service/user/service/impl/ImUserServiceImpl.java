package com.lld.im.service.user.service.impl;

import com.lld.im.common.ResponseVO;
import com.lld.im.service.user.dao.mapper.ImUserDataMapper;
import com.lld.im.service.user.model.req.ImportUserReq;
import com.lld.im.service.user.model.resp.ImportUserRes;
import com.lld.im.service.user.service.ImUserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;

import java.util.ArrayList;
import java.util.List;

@Service
public class ImUserServiceImpl implements ImUserService {

    @Autowired
    private ImUserDataMapper imUserDataMapper;

    @Transactional(propagation = Propagation.REQUIRED)
    @Override
    public ResponseVO importUser(ImportUserReq req) {

        // 检查数量
        if (req.getUserData().size() > 100) {
            // TODO 返回数量过多
        }

        List<String> successId = new ArrayList<>();
        List<String> errorId = new ArrayList<>();

        req.getUserData().forEach(item->{

            try {

                item.setAppId(req.getAppId());

                int insert = imUserDataMapper.insert(item);
                if (1 == insert) {
                    successId.add(item.getUserId());
                }

            } catch (Exception e) {
                e.printStackTrace();
                errorId.add(item.getUserId());
            }
        });


        ImportUserRes importUserRes = new ImportUserRes();
        importUserRes.setSuccessId(successId);
        importUserRes.setErrorId(errorId);


        return ResponseVO.successResponse(importUserRes);
    }
}
