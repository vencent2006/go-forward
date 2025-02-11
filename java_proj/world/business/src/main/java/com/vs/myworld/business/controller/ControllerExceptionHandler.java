package com.vs.myworld.business.controller;

import com.vs.myworld.business.exception.BusinessException;
import com.vs.myworld.business.resp.CommonResp;
import lombok.extern.slf4j.Slf4j;
import org.springframework.validation.BindException;
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

    /**
     * 校验异常统一处理 处理validation的消息返回
     *
     * @param e
     * @return
     */
    @ExceptionHandler(value = BindException.class) // 处理是什么异常
    @ResponseBody
    public CommonResp<Object> exceptionHandler(BindException e) {
        CommonResp<Object> commonResp = new CommonResp<>();
        log.error("校验异常: {}", e.getBindingResult().getAllErrors().get(0).getDefaultMessage());
        commonResp.setSuccess(false);
        commonResp.setMessage(e.getBindingResult().getAllErrors().get(0).getDefaultMessage());
        return commonResp;
    }
}
