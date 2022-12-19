package com.vincent.huobi.client;

import java.util.List;

import com.vincent.huobi.client.req.etf.ETFSwapListRequest;
import com.vincent.huobi.client.req.etf.ETFSwapRequest;
import com.vincent.huobi.model.etf.ETFConfig;
import com.vincent.huobi.model.etf.ETFSwapRecord;

public interface ETFClient {

  ETFConfig getConfig(String etfName);

  void etfSwap(ETFSwapRequest request);

  List<ETFSwapRecord> getEtfSwapList(ETFSwapListRequest request);

}
