package com.vincent.huobi.service.huobi.parser.market;

import java.util.List;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import com.vincent.huobi.model.market.MarketTradeEvent;
import com.vincent.huobi.service.huobi.parser.HuobiModelParser;
import com.vincent.huobi.service.huobi.utils.DataUtils;

public class MarketTradeEventParser implements HuobiModelParser<MarketTradeEvent> {

  @Override
  public MarketTradeEvent parse(JSONObject json) {
    String dataKey = DataUtils.getDataKey(json);

    JSONObject data = json.getJSONObject(dataKey);
    JSONArray dataArray = data.getJSONArray("data");

    return MarketTradeEvent.builder()
        .ch(json.getString("ch"))
        .ts(json.getLong("ts"))
        .list(new MarketTradeParser().parseArray(dataArray))
        .build();
  }

  @Override
  public MarketTradeEvent parse(JSONArray json) {
    return null;
  }

  @Override
  public List<MarketTradeEvent> parseArray(JSONArray jsonArray) {
    return null;
  }
}
