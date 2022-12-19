package com.vincent.huobi.client;

import java.util.List;

import com.vincent.huobi.client.req.trade.BatchCancelOpenOrdersRequest;
import com.vincent.huobi.client.req.trade.CreateOrderRequest;
import com.vincent.huobi.client.req.trade.FeeRateRequest;
import com.vincent.huobi.client.req.trade.MatchResultRequest;
import com.vincent.huobi.client.req.trade.OpenOrdersRequest;
import com.vincent.huobi.client.req.trade.OrderHistoryRequest;
import com.vincent.huobi.client.req.trade.OrdersRequest;
import com.vincent.huobi.client.req.trade.ReqOrderListRequest;
import com.vincent.huobi.client.req.trade.SubOrderUpdateRequest;
import com.vincent.huobi.client.req.trade.SubOrderUpdateV2Request;
import com.vincent.huobi.client.req.trade.SubTradeClearingRequest;
import com.vincent.huobi.constant.Options;
import com.vincent.huobi.constant.enums.ExchangeEnum;
import com.vincent.huobi.exception.SDKException;
import com.vincent.huobi.model.trade.BatchCancelOpenOrdersResult;
import com.vincent.huobi.model.trade.BatchCancelOrderResult;
import com.vincent.huobi.model.trade.FeeRate;
import com.vincent.huobi.model.trade.MatchResult;
import com.vincent.huobi.model.trade.Order;
import com.vincent.huobi.model.trade.OrderDetailReq;
import com.vincent.huobi.model.trade.OrderListReq;
import com.vincent.huobi.model.trade.OrderUpdateEvent;
import com.vincent.huobi.model.trade.OrderUpdateV2Event;
import com.vincent.huobi.model.trade.TradeClearingEvent;
import com.vincent.huobi.service.huobi.HuobiTradeService;
import com.vincent.huobi.utils.ResponseCallback;

public interface TradeClient {

  Long createOrder(CreateOrderRequest request);

  Long cancelOrder(Long orderId);

  Integer cancelOrder(String clientOrderId);

  BatchCancelOpenOrdersResult batchCancelOpenOrders(BatchCancelOpenOrdersRequest request);

  BatchCancelOrderResult batchCancelOrder(List<Long> ids);

  List<Order> getOpenOrders(OpenOrdersRequest request);

  Order getOrder(Long orderId);

  Order getOrder(String clientOrderId);

  List<Order> getOrders(OrdersRequest request);

  List<Order> getOrdersHistory(OrderHistoryRequest request);

  List<MatchResult> getMatchResult(Long orderId);

  List<MatchResult> getMatchResults(MatchResultRequest request);

  List<FeeRate> getFeeRate(FeeRateRequest request);

  void subOrderUpdateV2(SubOrderUpdateV2Request request, ResponseCallback<OrderUpdateV2Event> callback);

  void subTradeClearing(SubTradeClearingRequest request, ResponseCallback<TradeClearingEvent> callback);

  static TradeClient create(Options options) {

    if (options.getExchange().equals(ExchangeEnum.HUOBI)) {
      return new HuobiTradeService(options);
    }
    throw new SDKException(SDKException.INPUT_ERROR, "Unsupport Exchange.");
  }

}
