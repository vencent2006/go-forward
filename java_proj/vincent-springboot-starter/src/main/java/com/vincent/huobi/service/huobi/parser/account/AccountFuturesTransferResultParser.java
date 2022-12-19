package com.vincent.huobi.service.huobi.parser.account;

import java.util.List;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import com.vincent.huobi.model.account.AccountFuturesTransferResult;
import com.vincent.huobi.service.huobi.parser.HuobiModelParser;

public class AccountFuturesTransferResultParser implements HuobiModelParser<AccountFuturesTransferResult> {

  @Override
  public AccountFuturesTransferResult parse(JSONObject json) {
    return AccountFuturesTransferResult.builder().data(json.getLong("data")).build();
  }

  @Override
  public AccountFuturesTransferResult parse(JSONArray json) {
    return null;
  }

  @Override
  public List<AccountFuturesTransferResult> parseArray(JSONArray jsonArray) {
    return null;
  }
}
