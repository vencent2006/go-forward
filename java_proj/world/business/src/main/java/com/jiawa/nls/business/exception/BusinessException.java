package com.jiawa.nls.business.exception;

import lombok.Data;

@Data
public class BusinessException extends RuntimeException {
    private BusinessExceptionEnum e;

    public BusinessException(BusinessExceptionEnum e) {
        super(e.getDesc());
        this.e = e;
    }

    /**
     * 不写入堆栈信息，简化异常日志，提高性能
     */
    @Override
    public Throwable fillInStackTrace() {
        return this;
    }
}
