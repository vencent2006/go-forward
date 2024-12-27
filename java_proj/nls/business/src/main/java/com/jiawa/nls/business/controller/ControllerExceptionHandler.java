package com.jiawa.nls.business.controller;

import com.jiawa.nls.business.exception.BusinessException;
import com.jiawa.nls.business.resp.CommonResp;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;

/**
 * 统一异常处理，数据预处理等
 */
@Slf4j
@ControllerAdvice
public class ControllerExceptionHandler {
    /**
     * 所有异常统一处理
     *
     * @param e
     * @return
     * @throws Exception
     */
    @ExceptionHandler(value = Exception.class) // 处理是什么异常
    @ResponseBody
    public CommonResp<Object> exceptionHandler(Exception e) {
        CommonResp<Object> commonResp = new CommonResp<>();
        log.error("系统异常: ", e);
        commonResp.setSuccess(false);
        commonResp.setMessage("系统出现异常, 请联系管理员");
        return commonResp;
    }

    /**
     * 所有业务异常统一处理
     *
     * @param e
     * @return
     */
    @ExceptionHandler(value = BusinessException.class) // 处理是什么异常
    @ResponseBody
    public CommonResp<Object> exceptionHandler(BusinessException e) {
        CommonResp<Object> commonResp = new CommonResp<>();
        log.error("系统异常: ", e);
        commonResp.setSuccess(false);
        commonResp.setMessage(e.getE().getDesc());
        return commonResp;
    }
}
