package com.lld.im.service.friendship.service;

import com.lld.im.common.ResponseVO;
import com.lld.im.service.friendship.model.req.AddFriendReq;
import com.lld.im.service.friendship.model.req.ImportFriendShipReq;

public interface ImFriendService {

    public ResponseVO importFriendShip(ImportFriendShipReq req);

    public ResponseVO AddFriend(AddFriendReq req);
}
