package com.vs.myworld.business.req;

import lombok.Data;

import javax.validation.constraints.NotBlank;

@Data
public class SmsCodeResetReq {
    @NotBlank(message = "[手机号] 不能为空")
    private String mobile;
}
