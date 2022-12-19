package com.vincent.huobi.service.huobi.parser.subuser;

import java.util.List;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import com.vincent.huobi.model.subuser.SubUserManagementResult;
import com.vincent.huobi.service.huobi.parser.HuobiModelParser;

public class SubUserManagementResultParser implements HuobiModelParser<SubUserManagementResult> {

  @Override
  public SubUserManagementResult parse(JSONObject json) {
    return json.toJavaObject(SubUserManagementResult.class);
  }

  @Override
  public SubUserManagementResult parse(JSONArray json) {
    return null;
  }

  @Override
  public List<SubUserManagementResult> parseArray(JSONArray jsonArray) {
    return null;
  }
}
