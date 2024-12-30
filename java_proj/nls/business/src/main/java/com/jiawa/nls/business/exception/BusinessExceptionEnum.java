package com.jiawa.nls.business.exception;

import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor
public enum BusinessExceptionEnum {
    DEMO_MOBILE_NOT_NULL("手机号不能为空"),
    SMS_CODE_TOO_FREQUENT("短信请求过于频繁");

    @Getter
    private final String desc;
}
