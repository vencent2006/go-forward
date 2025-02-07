package com.jiawa.nls.business.resp;

import lombok.Data;

@Data
public class DemoQueryResp {
    private Long id;

    private String mobile;

    // 表实体Demo里的属性不能修改，使用resp类，可以增加或减少属性，让返回值更灵活
    // private String password;
    
}