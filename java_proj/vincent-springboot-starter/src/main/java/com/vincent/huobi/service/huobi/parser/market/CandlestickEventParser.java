package com.vincent.huobi.service.huobi.parser.market;

import java.util.List;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import com.vincent.huobi.model.market.CandlestickEvent;
import com.vincent.huobi.service.huobi.parser.HuobiModelParser;
import com.vincent.huobi.service.huobi.utils.DataUtils;

public class CandlestickEventParser implements HuobiModelParser<CandlestickEvent> {

  @Override
  public CandlestickEvent parse(JSONObject json) {

    String dataKey = DataUtils.getDataKey(json);

    return CandlestickEvent.builder()
        .ch(json.getString("ch"))
        .ts(json.getLong("ts"))
        .candlestick(new CandlestickParser().parse(json.getJSONObject(dataKey)))
        .build();
  }

  @Override
  public CandlestickEvent parse(JSONArray json) {
    return null;
  }

  @Override
  public List<CandlestickEvent> parseArray(JSONArray jsonArray) {
    return null;
  }
}
