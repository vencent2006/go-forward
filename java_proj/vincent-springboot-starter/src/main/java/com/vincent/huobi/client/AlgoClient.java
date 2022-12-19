package com.vincent.huobi.client;

import com.vincent.huobi.client.req.algo.CancelAlgoOrderRequest;
import com.vincent.huobi.client.req.algo.CreateAlgoOrderRequest;
import com.vincent.huobi.client.req.algo.GetHistoryAlgoOrdersRequest;
import com.vincent.huobi.client.req.algo.GetOpenAlgoOrdersRequest;
import com.vincent.huobi.constant.Options;
import com.vincent.huobi.constant.enums.ExchangeEnum;
import com.vincent.huobi.exception.SDKException;
import com.vincent.huobi.model.algo.AlgoOrder;
import com.vincent.huobi.model.algo.CancelAlgoOrderResult;
import com.vincent.huobi.model.algo.CreateAlgoOrderResult;
import com.vincent.huobi.model.algo.GetHistoryAlgoOrdersResult;
import com.vincent.huobi.model.algo.GetOpenAlgoOrdersResult;
import com.vincent.huobi.service.huobi.HuobiAlgoService;

public interface AlgoClient {

  CreateAlgoOrderResult createAlgoOrder(CreateAlgoOrderRequest request);

  CancelAlgoOrderResult cancelAlgoOrder(CancelAlgoOrderRequest request);

  GetOpenAlgoOrdersResult getOpenAlgoOrders(GetOpenAlgoOrdersRequest request);

  GetHistoryAlgoOrdersResult getHistoryAlgoOrders(GetHistoryAlgoOrdersRequest request);

  AlgoOrder getAlgoOrdersSpecific(String clientOrderId);


  static AlgoClient create(Options options) {

    if (options.getExchange().equals(ExchangeEnum.HUOBI)) {
      return new HuobiAlgoService(options);
    }
    throw new SDKException(SDKException.INPUT_ERROR, "Unsupport Exchange.");
  }
}
