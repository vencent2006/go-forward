package com.jiawa.nls.business.enums;

import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor
public enum SmsCodeUseEnum {
    REGISTER("0", "注册"),
    RESET("1", "重置密码");

    @Getter
    private final String code;
    @Getter
    private final String desc;
}
