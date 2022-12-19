package com.vincent.huobi.constant;

import com.vincent.huobi.constant.enums.ExchangeEnum;

public interface Options {

  String getApiKey();

  String getSecretKey();

  ExchangeEnum getExchange();

  String getRestHost();

  String getWebSocketHost();

  boolean isWebSocketAutoConnect();

}
