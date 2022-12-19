package com.vincent.huobi.client;

import java.util.List;

import com.vincent.huobi.client.req.margin.IsolatedMarginAccountRequest;
import com.vincent.huobi.client.req.margin.IsolatedMarginApplyLoanRequest;
import com.vincent.huobi.client.req.margin.IsolatedMarginLoanInfoRequest;
import com.vincent.huobi.client.req.margin.IsolatedMarginLoanOrdersRequest;
import com.vincent.huobi.client.req.margin.IsolatedMarginRepayLoanRequest;
import com.vincent.huobi.client.req.margin.IsolatedMarginTransferRequest;
import com.vincent.huobi.constant.Options;
import com.vincent.huobi.constant.enums.ExchangeEnum;
import com.vincent.huobi.exception.SDKException;
import com.vincent.huobi.model.isolatedmargin.IsolatedMarginAccount;
import com.vincent.huobi.model.isolatedmargin.IsolatedMarginLoadOrder;
import com.vincent.huobi.model.isolatedmargin.IsolatedMarginSymbolInfo;
import com.vincent.huobi.service.huobi.HuobiIsolatedMarginService;

public interface IsolatedMarginClient {

  Long transfer(IsolatedMarginTransferRequest request);

  Long applyLoan(IsolatedMarginApplyLoanRequest request);

  Long repayLoan(IsolatedMarginRepayLoanRequest request);

  List<IsolatedMarginLoadOrder> getLoanOrders(IsolatedMarginLoanOrdersRequest request);

  List<IsolatedMarginAccount> getLoanBalance(IsolatedMarginAccountRequest request);

  List<IsolatedMarginSymbolInfo> getLoanInfo(IsolatedMarginLoanInfoRequest request);

  static IsolatedMarginClient create(Options options) {

    if (options.getExchange().equals(ExchangeEnum.HUOBI)) {
      return new HuobiIsolatedMarginService(options);
    }
    throw new SDKException(SDKException.INPUT_ERROR, "Unsupport Exchange.");
  }
}
