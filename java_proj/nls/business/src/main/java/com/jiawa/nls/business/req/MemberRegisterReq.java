package com.jiawa.nls.business.req;

import lombok.Data;

import javax.validation.constraints.NotBlank;

@Data
public class MemberRegisterReq {
    /**
     * 手机号
     */
    @NotBlank(message = "【手机号】不能为空")
    private String mobile;
    /**
     * 密码
     */
    @NotBlank(message = "【密码】不能为空")
    private String password;
    /**
     * 验证码
     */
    @NotBlank(message = "【验证码】不能为空")
    private String code;
}
