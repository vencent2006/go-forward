package com.vincent.huobi.client.req.account;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import com.vincent.huobi.constant.enums.AccountTypeEnum;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class AccountAssetValuationRequest {

  AccountTypeEnum accountType;

  String valuationCurrency;

  Long subUid;

}
