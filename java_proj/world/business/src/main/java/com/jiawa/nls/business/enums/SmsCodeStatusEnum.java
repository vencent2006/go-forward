package com.jiawa.nls.business.enums;

import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor
public enum SmsCodeStatusEnum {
    UNUSED("0", "未使用"),
    USED("1", "已使用");

    @Getter
    private final String code;
    @Getter
    private final String desc;
}
