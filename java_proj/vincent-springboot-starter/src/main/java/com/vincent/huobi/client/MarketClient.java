package com.vincent.huobi.client;

import java.util.List;

import com.vincent.huobi.client.req.market.CandlestickRequest;
import com.vincent.huobi.client.req.market.MarketDepthRequest;
import com.vincent.huobi.client.req.market.MarketDetailMergedRequest;
import com.vincent.huobi.client.req.market.MarketDetailRequest;
import com.vincent.huobi.client.req.market.MarketHistoryTradeRequest;
import com.vincent.huobi.client.req.market.MarketTradeRequest;
import com.vincent.huobi.client.req.market.ReqCandlestickRequest;
import com.vincent.huobi.client.req.market.ReqMarketDepthRequest;
import com.vincent.huobi.client.req.market.ReqMarketDetailRequest;
import com.vincent.huobi.client.req.market.ReqMarketTradeRequest;
import com.vincent.huobi.client.req.market.SubCandlestickRequest;
import com.vincent.huobi.client.req.market.SubMarketBBORequest;
import com.vincent.huobi.client.req.market.SubMarketDepthRequest;
import com.vincent.huobi.client.req.market.SubMarketDetailRequest;
import com.vincent.huobi.client.req.market.SubMarketTradeRequest;
import com.vincent.huobi.client.req.market.SubMbpIncrementalUpdateRequest;
import com.vincent.huobi.client.req.market.SubMbpRefreshUpdateRequest;
import com.vincent.huobi.constant.Options;
import com.vincent.huobi.constant.enums.ExchangeEnum;
import com.vincent.huobi.exception.SDKException;
import com.vincent.huobi.model.market.Candlestick;
import com.vincent.huobi.model.market.CandlestickEvent;
import com.vincent.huobi.model.market.CandlestickReq;
import com.vincent.huobi.model.market.MarketBBOEvent;
import com.vincent.huobi.model.market.MarketDepth;
import com.vincent.huobi.model.market.MarketDepthEvent;
import com.vincent.huobi.model.market.MarketDepthReq;
import com.vincent.huobi.model.market.MarketDetail;
import com.vincent.huobi.model.market.MarketDetailEvent;
import com.vincent.huobi.model.market.MarketDetailMerged;
import com.vincent.huobi.model.market.MarketDetailReq;
import com.vincent.huobi.model.market.MarketTicker;
import com.vincent.huobi.model.market.MarketTrade;
import com.vincent.huobi.model.market.MarketTradeEvent;
import com.vincent.huobi.model.market.MarketTradeReq;
import com.vincent.huobi.model.market.MbpIncrementalUpdateEvent;
import com.vincent.huobi.model.market.MbpRefreshUpdateEvent;
import com.vincent.huobi.service.huobi.HuobiMarketService;
import com.vincent.huobi.service.huobi.connection.HuobiWebSocketConnection;
import com.vincent.huobi.utils.ResponseCallback;
import com.vincent.huobi.utils.WebSocketConnection;

public interface MarketClient {

  List<Candlestick> getCandlestick(CandlestickRequest request);

  MarketDetailMerged getMarketDetailMerged(MarketDetailMergedRequest request);

  MarketDetail getMarketDetail(MarketDetailRequest request);

  List<MarketTicker> getTickers();

  MarketDepth getMarketDepth(MarketDepthRequest request);

  List<MarketTrade> getMarketTrade(MarketTradeRequest request);

  List<MarketTrade> getMarketHistoryTrade(MarketHistoryTradeRequest request);

  void subCandlestick(SubCandlestickRequest request, ResponseCallback<CandlestickEvent> callback);

  void subMarketDetail(SubMarketDetailRequest request, ResponseCallback<MarketDetailEvent> callback);

  void subMarketDepth(SubMarketDepthRequest request, ResponseCallback<MarketDepthEvent> callback);

  void subMarketTrade(SubMarketTradeRequest request, ResponseCallback<MarketTradeEvent> callback);

  void subMarketBBO(SubMarketBBORequest request, ResponseCallback<MarketBBOEvent> callback);

  void subMbpRefreshUpdate(SubMbpRefreshUpdateRequest request, ResponseCallback<MbpRefreshUpdateEvent> callback);

  WebSocketConnection subMbpIncrementalUpdate(SubMbpIncrementalUpdateRequest request, ResponseCallback<MbpIncrementalUpdateEvent> callback);

  WebSocketConnection reqMbpIncrementalUpdate(SubMbpIncrementalUpdateRequest request, WebSocketConnection connection);

  void reqCandlestick(ReqCandlestickRequest request, ResponseCallback<CandlestickReq> callback);

  void reqMarketDepth(ReqMarketDepthRequest request, ResponseCallback<MarketDepthReq> callback);

  void reqMarketTrade(ReqMarketTradeRequest request, ResponseCallback<MarketTradeReq> callback);

  void reqMarketDetail(ReqMarketDetailRequest request, ResponseCallback<MarketDetailReq> callback);

  static MarketClient create(Options options) {

    if (options.getExchange().equals(ExchangeEnum.HUOBI)) {
      return new HuobiMarketService(options);
    }
    throw new SDKException(SDKException.INPUT_ERROR, "Unsupport Exchange.");
  }


}
