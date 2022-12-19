package com.vincent.huobi.constant.enums;

import lombok.AllArgsConstructor;
import lombok.Getter;

/**
 * buy, sell, both.
 */
@Getter
@AllArgsConstructor
public enum OrderSideEnum {
  BUY("buy"),
  SELL("sell");

  private final String code;
}
