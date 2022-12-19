package com.vincent.huobi.client.req.account;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;

import com.vincent.huobi.constant.enums.BalanceModeEnum;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
@ToString
public class SubAccountChangeRequest {

  private BalanceModeEnum balanceMode;

}
