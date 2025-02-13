package com.vs.myworld.business.resp;

import lombok.Data;

import java.util.List;

@Data
public class WorldQueryInfo {
    private Long total;

    private List<WorldQueryResp> list;
}