package com.lld.im.service.friendship.model.req;

import com.lld.im.common.enums.FriendShipStatusEnum;
import com.lld.im.common.model.RequestBase;
import lombok.Data;

import javax.validation.constraints.NotBlank;
import java.util.List;

@Data
public class AddFriendReq extends RequestBase {

    @NotBlank(message = "fromId不能为空")
    private String fromId;

    private List<ImportFriendDto> friendItem;

    @Data
    public static class ImportFriendDto {

        private String toId;

        private String remark;

        private String addSource;

        private Integer status = FriendShipStatusEnum.FRIEND_STATUS_NO_FRIEND.getCode(); // 默认：不是好友

        private Integer black = FriendShipStatusEnum.BLACK_STATUS_NORMAL.getCode(); // 默认: 正常

    }
}
