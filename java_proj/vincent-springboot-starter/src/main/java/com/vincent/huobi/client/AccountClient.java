package com.vincent.huobi.client;

import java.util.List;

import com.vincent.huobi.client.req.account.AccountAssetValuationRequest;
import com.vincent.huobi.client.req.account.AccountBalanceRequest;
import com.vincent.huobi.client.req.account.AccountFuturesTransferRequest;
import com.vincent.huobi.client.req.account.AccountHistoryRequest;
import com.vincent.huobi.client.req.account.AccountLedgerRequest;
import com.vincent.huobi.client.req.account.AccountTransferRequest;
import com.vincent.huobi.client.req.account.PointRequest;
import com.vincent.huobi.client.req.account.PointTransferRequest;
import com.vincent.huobi.client.req.account.SubAccountUpdateRequest;
import com.vincent.huobi.constant.Options;
import com.vincent.huobi.constant.enums.ExchangeEnum;
import com.vincent.huobi.exception.SDKException;
import com.vincent.huobi.model.account.Account;
import com.vincent.huobi.model.account.AccountAssetValuationResult;
import com.vincent.huobi.model.account.AccountBalance;
import com.vincent.huobi.model.account.AccountFuturesTransferResult;
import com.vincent.huobi.model.account.AccountHistory;
import com.vincent.huobi.model.account.AccountLedgerResult;
import com.vincent.huobi.model.account.AccountTransferResult;
import com.vincent.huobi.model.account.AccountUpdateEvent;
import com.vincent.huobi.model.account.Point;
import com.vincent.huobi.model.account.PointTransferResult;
import com.vincent.huobi.model.subuser.SubUserState;
import com.vincent.huobi.service.huobi.HuobiAccountService;
import com.vincent.huobi.utils.ResponseCallback;

public interface AccountClient {

  /**
   * Get User Account List
   * @return
   */
  List<Account> getAccounts();

  /**
   * Get User Account Balance
   * @param request
   * @return
   */
  AccountBalance getAccountBalance(AccountBalanceRequest request);

  List<AccountHistory> getAccountHistory(AccountHistoryRequest request);

  AccountLedgerResult getAccountLedger(AccountLedgerRequest request);

  AccountTransferResult accountTransfer(AccountTransferRequest request);

  AccountFuturesTransferResult accountFuturesTransfer(AccountFuturesTransferRequest request);

  Point getPoint(PointRequest request);

  PointTransferResult pointTransfer(PointTransferRequest request);

  AccountAssetValuationResult accountAssetValuation(AccountAssetValuationRequest request);

  void subAccountsUpdate(SubAccountUpdateRequest request, ResponseCallback<AccountUpdateEvent> callback);

  static AccountClient create(Options options) {

    if (options.getExchange().equals(ExchangeEnum.HUOBI)) {
      return new HuobiAccountService(options);
    }
    throw new SDKException(SDKException.INPUT_ERROR, "Unsupport Exchange.");
  }
}
