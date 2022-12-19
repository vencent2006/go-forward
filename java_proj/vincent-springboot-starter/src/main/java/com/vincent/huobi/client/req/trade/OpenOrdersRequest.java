package com.vincent.huobi.client.req.trade;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import com.vincent.huobi.constant.enums.OrderSideEnum;
import com.vincent.huobi.constant.enums.QueryDirectionEnum;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class OpenOrdersRequest {

  private  String symbol;

  private Long accountId;

  private Integer size;

  private OrderSideEnum side;

  private QueryDirectionEnum direct;

  private Long from;

}
