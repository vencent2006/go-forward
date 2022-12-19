package com.vincent.huobi.service.huobi.parser.subuser;

import java.util.List;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import com.vincent.huobi.model.subuser.SubUserDeposit;
import com.vincent.huobi.service.huobi.parser.HuobiModelParser;

public class SubUserDepositParser implements HuobiModelParser<SubUserDeposit> {

  @Override
  public SubUserDeposit parse(JSONObject json) {
    return null;
  }

  @Override
  public SubUserDeposit parse(JSONArray json) {
    return null;
  }

  @Override
  public List<SubUserDeposit> parseArray(JSONArray jsonArray) {
    return jsonArray.toJavaList(SubUserDeposit.class);
  }
}
