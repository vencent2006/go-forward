package com.jiawa.nls.business.exception;

import lombok.AllArgsConstructor;
import lombok.Getter;

@AllArgsConstructor
public enum BusinessExceptionEnum {
    DEMO_MOBILE_NOT_NULL("手机号不能为空");
    @Getter
    private final String desc;
}
