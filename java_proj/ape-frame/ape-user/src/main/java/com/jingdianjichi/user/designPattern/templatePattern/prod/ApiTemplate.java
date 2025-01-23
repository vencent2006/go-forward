package com.jingdianjichi.user.designPattern.templatePattern.prod;

import com.jingdianjichi.bean.Result;
import com.jingdianjichi.bean.ResultCode;

/**
 * api模板
 */
public class ApiTemplate {
    public void execute(Result result, final Action action) {
        try {
            action.validate();
            action.execute();
            action.after();
            result.setSuccess(true);
            result.setCode(1024);
        } catch (Exception e) {
            result.setSuccess(false);
            result.setCode(ResultCode.ERROR);
        }
    }
}
