package com.vincent.huobi.service.huobi.parser.market;

import java.util.List;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import com.vincent.huobi.model.market.MarketBBO;
import com.vincent.huobi.service.huobi.parser.HuobiModelParser;

public class MarketBBOParser implements HuobiModelParser<MarketBBO> {

  @Override
  public MarketBBO parse(JSONObject json) {
    return json.toJavaObject(MarketBBO.class);
  }

  @Override
  public MarketBBO parse(JSONArray json) {
    return null;
  }

  @Override
  public List<MarketBBO> parseArray(JSONArray jsonArray) {
    return null;
  }
}
