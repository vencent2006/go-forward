package com.jiawa.nls.business.enums;

import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor
public enum SmsCodeUseEnum {
    REGISTER("0", "未使用"),
    FORGET_PASSWORD("1", "忘记密码");

    @Getter
    private final String code;
    @Getter
    private final String desc;
}
