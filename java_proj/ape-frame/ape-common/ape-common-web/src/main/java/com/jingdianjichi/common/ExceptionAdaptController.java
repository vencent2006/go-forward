package com.jingdianjichi.common;

import com.jingdianjichi.bean.Result;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;

@ControllerAdvice
public class ExceptionAdaptController {
    @ExceptionHandler(value = RuntimeException.class)
    @ResponseBody
    public Result runtimeException(RuntimeException e) {
        e.printStackTrace();
        return Result.fail(e.getMessage());
    }

    @ResponseBody
    @ExceptionHandler(value = Exception.class)
    public Result exception(RuntimeException e) {
        e.printStackTrace();
        return Result.fail(e.getMessage());
    }
}
