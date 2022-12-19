package com.vincent.huobi.service.huobi.parser.subuser;

import java.util.List;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import com.vincent.huobi.model.subuser.SubUserApiKeyGenerationResult;
import com.vincent.huobi.service.huobi.parser.HuobiModelParser;

public class SubUserApiKeyGenerationResultParser implements HuobiModelParser<SubUserApiKeyGenerationResult> {

  @Override
  public SubUserApiKeyGenerationResult parse(JSONObject json) {
    return json.toJavaObject(SubUserApiKeyGenerationResult.class);
  }

  @Override
  public SubUserApiKeyGenerationResult parse(JSONArray json) {
    return null;
  }

  @Override
  public List<SubUserApiKeyGenerationResult> parseArray(JSONArray jsonArray) {
    return null;
  }
}
