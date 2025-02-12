package com.vs.myworld.business.resp;

import lombok.Data;

import java.util.Date;

@Data
public class WorldQueryResp {
    private Long id;

    private String name;

    private String image;

    private String desc;

    private Long creatorUid;

    private Byte status;

    private Long roleCount;

    private Long followCount;

    private Date beginAt;

    private Date endAt;

    private Integer elapseSpeed;

    // private Date createdAt;
    //
    // private Date updatedAt;
}