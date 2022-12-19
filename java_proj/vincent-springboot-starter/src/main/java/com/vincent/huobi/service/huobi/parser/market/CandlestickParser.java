package com.vincent.huobi.service.huobi.parser.market;

import java.util.List;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import com.vincent.huobi.model.market.Candlestick;
import com.vincent.huobi.service.huobi.parser.HuobiModelParser;

public class CandlestickParser implements HuobiModelParser<Candlestick> {

  @Override
  public Candlestick parse(JSONObject json) {
    return json.toJavaObject(Candlestick.class);
  }

  @Override
  public Candlestick parse(JSONArray json) {
    return null;
  }

  @Override
  public List<Candlestick> parseArray(JSONArray jsonArray) {
    return jsonArray.toJavaList(Candlestick.class);
  }
}
