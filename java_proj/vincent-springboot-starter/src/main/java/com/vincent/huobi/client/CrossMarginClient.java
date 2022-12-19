package com.vincent.huobi.client;

import java.util.List;

import com.vincent.huobi.client.req.crossmargin.CrossMarginApplyLoanRequest;
import com.vincent.huobi.client.req.crossmargin.CrossMarginLoanOrdersRequest;
import com.vincent.huobi.client.req.crossmargin.CrossMarginRepayLoanRequest;
import com.vincent.huobi.client.req.crossmargin.CrossMarginTransferRequest;
import com.vincent.huobi.client.req.crossmargin.GeneralLoanOrdersRequest;
import com.vincent.huobi.client.req.crossmargin.GeneralRepayLoanRequest;
import com.vincent.huobi.constant.Options;
import com.vincent.huobi.constant.enums.ExchangeEnum;
import com.vincent.huobi.exception.SDKException;
import com.vincent.huobi.model.crossmargin.CrossMarginAccount;
import com.vincent.huobi.model.crossmargin.CrossMarginCurrencyInfo;
import com.vincent.huobi.model.crossmargin.CrossMarginLoadOrder;
import com.vincent.huobi.model.crossmargin.GeneralRepayLoanRecord;
import com.vincent.huobi.model.crossmargin.GeneralRepayLoanResult;
import com.vincent.huobi.service.huobi.HuobiCrossMarginService;

public interface CrossMarginClient {

  Long transfer(CrossMarginTransferRequest request);

  Long applyLoan(CrossMarginApplyLoanRequest request);

  void repayLoan(CrossMarginRepayLoanRequest request);

  List<CrossMarginLoadOrder> getLoanOrders(CrossMarginLoanOrdersRequest request);

  CrossMarginAccount getLoanBalance();

  List<CrossMarginCurrencyInfo> getLoanInfo();

  static CrossMarginClient create(Options options) {

    if (options.getExchange().equals(ExchangeEnum.HUOBI)) {
      return new HuobiCrossMarginService(options);
    }
    throw new SDKException(SDKException.INPUT_ERROR, "Unsupport Exchange.");
  }

  List<GeneralRepayLoanResult> repayLoan(GeneralRepayLoanRequest request);

  List<GeneralRepayLoanRecord> getRepaymentLoanRecords(GeneralLoanOrdersRequest request);

}
