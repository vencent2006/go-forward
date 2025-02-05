package com.jingdianjichi.user.enums;

import lombok.Getter;

import java.util.Map;
import java.util.stream.Collectors;
import java.util.stream.Stream;

@Getter
public enum ChannelEnum {
    DOU_YIN(0, "抖音渠道"),
    BILI_BILI(1, "B站渠道");
    private int code;
    private String desc;

    ChannelEnum(int code, String desc) {
        this.code = code;
        this.desc = desc;
    }

    public static final Map<Integer, ChannelEnum> clientChannelMap = Stream.of(ChannelEnum.values())
            .collect(Collectors.toMap(e -> e.getCode(), e -> e));

    public static ChannelEnum getChannel(Integer code) {
        return clientChannelMap.get(code);
    }

}
