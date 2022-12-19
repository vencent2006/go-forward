package com.vincent.huobi.service.huobi.parser.subuser;

import java.util.List;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import com.vincent.huobi.model.subuser.GetSubUserDepositResult;
import com.vincent.huobi.service.huobi.parser.HuobiModelParser;

public class GetSubUserDepositResultParser implements HuobiModelParser<GetSubUserDepositResult> {

  @Override
  public GetSubUserDepositResult parse(JSONObject json) {
    return GetSubUserDepositResult.builder()
        .nextId(json.getLong("nextId"))
        .list(new SubUserDepositParser().parseArray(json.getJSONArray("data")))
        .build();
  }

  @Override
  public GetSubUserDepositResult parse(JSONArray json) {
    return null;
  }

  @Override
  public List<GetSubUserDepositResult> parseArray(JSONArray jsonArray) {
    return null;
  }
}
