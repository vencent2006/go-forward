package com.vincent.huobi.client.req.subuser;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import com.vincent.huobi.constant.enums.TradableMarketAccountTypeEnum;
import com.vincent.huobi.constant.enums.TradableMarketActivationEnums;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class SubUserTradableMarketRequest {

  private String subUids;

  private TradableMarketAccountTypeEnum accountType;

  private TradableMarketActivationEnums activation;

}
