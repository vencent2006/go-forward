package com.vincent.huobi.client;

import java.util.List;

import com.vincent.huobi.client.req.wallet.CreateWithdrawRequest;
import com.vincent.huobi.client.req.wallet.DepositAddressRequest;
import com.vincent.huobi.client.req.wallet.DepositWithdrawRequest;
import com.vincent.huobi.client.req.wallet.WithdrawAddressRequest;
import com.vincent.huobi.client.req.wallet.WithdrawQuotaRequest;
import com.vincent.huobi.constant.Options;
import com.vincent.huobi.constant.enums.ExchangeEnum;
import com.vincent.huobi.exception.SDKException;
import com.vincent.huobi.model.wallet.DepositAddress;
import com.vincent.huobi.model.wallet.DepositWithdraw;
import com.vincent.huobi.model.wallet.WithdrawAddressResult;
import com.vincent.huobi.model.wallet.WithdrawQuota;
import com.vincent.huobi.service.huobi.HuobiWalletService;

public interface WalletClient {

  List<DepositAddress> getDepositAddress(DepositAddressRequest request);

  WithdrawQuota getWithdrawQuota(WithdrawQuotaRequest request);

  WithdrawAddressResult getWithdrawAddress(WithdrawAddressRequest request);

  Long createWithdraw(CreateWithdrawRequest request);

  Long cancelWithdraw(Long withdrawId);

  List<DepositWithdraw> getDepositWithdraw(DepositWithdrawRequest request);

  static WalletClient create(Options options) {

    if (options.getExchange().equals(ExchangeEnum.HUOBI)) {
      return new HuobiWalletService(options);
    }
    throw new SDKException(SDKException.INPUT_ERROR, "Unsupport Exchange.");
  }
}
