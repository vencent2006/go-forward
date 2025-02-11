package com.vs.myworld.business.resp;

import lombok.Data;

@Data
public class MemberLoginResp {
    private String name;
    private String token;
}