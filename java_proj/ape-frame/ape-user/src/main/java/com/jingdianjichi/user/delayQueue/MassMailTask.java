package com.jingdianjichi.user.delayQueue;

import lombok.Data;

import java.util.Date;

@Data
public class MassMailTask {
    private Long taskId;
    private Date startTime;
}
