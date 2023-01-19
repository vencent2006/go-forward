package com.lld.im.service.user.model.resp;

import lombok.Data;

import java.util.List;

@Data
public class ImportUserRes {

    private List<String> successId;
    private List<String> errorId;
}
