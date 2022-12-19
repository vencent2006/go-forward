package com.vincent.huobi.constant.enums;

import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public enum  AccountTransferAccountTypeEnum {

  SPOT("spot"),
  MARGIN("margin"),

  ;
  private String accountType;

}
