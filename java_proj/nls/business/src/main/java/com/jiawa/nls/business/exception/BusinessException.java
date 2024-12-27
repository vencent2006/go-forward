package com.jiawa.nls.business.exception;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class BusinessException extends RuntimeException {
    private BusinessExceptionEnum e;
}
