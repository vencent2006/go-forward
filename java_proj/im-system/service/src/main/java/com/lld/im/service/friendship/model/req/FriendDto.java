package com.lld.im.service.friendship.model.req;

import lombok.Data;

@Data
public class FriendDto {

    private String toId;

    private String remark; // 备注

    private String addSource; // 来源

    private String extra; // 扩展

}
