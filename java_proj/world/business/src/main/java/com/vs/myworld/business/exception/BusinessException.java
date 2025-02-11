package com.vs.myworld.business.exception;

import lombok.Data;

@Data
public class BusinessException extends RuntimeException {
    private com.vs.myworld.business.exception.BusinessExceptionEnum e;

    public BusinessException(com.vs.myworld.business.exception.BusinessExceptionEnum e) {
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
