package com.vincent.huobi.client.req.algo;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import com.vincent.huobi.constant.enums.QuerySortEnum;
import com.vincent.huobi.constant.enums.algo.AlgoOrderSideEnum;
import com.vincent.huobi.constant.enums.algo.AlgoOrderTypeEnum;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class GetOpenAlgoOrdersRequest {

  private Long accountId;

  private String symbol;

  private AlgoOrderSideEnum orderSide;

  private AlgoOrderTypeEnum orderType;

  private QuerySortEnum sort;

  private Integer limit;

  private Long fromId;


}
