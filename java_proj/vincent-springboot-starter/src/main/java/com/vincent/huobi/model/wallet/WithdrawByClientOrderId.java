package com.vincent.huobi.model.wallet;

import com.alibaba.fastjson.annotation.JSONField;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.math.BigDecimal;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class WithdrawByClientOrderId {
    private Long id;

    @JSONField(name = "client-order-id")
    private String clientOrderId;

    private String type;

    private String currency;

    private String txHash;

    private String chain;

    private BigDecimal amount;

    private String address;

    private String addressTag;

    private BigDecimal fee;

    private String state;

    private String errorCode;

    private String errorMessage;

    private Long createdAt;

    private Long updatedAt;
}

/*
 {
    "status": "ok",
    "data": {
        "id": 101123262,
        "client-order-id": "1113",
        "type": "withdraw",
        "sub-type": "FAST",
        "currency": "usdt",
        "chain": "usdt",
        "tx-hash": "",
        "amount": 1.200000000000000000,
        "from-addr-tag": "",
        "address": "1PL24EbWrNNrnMKw1cxAHPsebUz7DdhWTx",
        "address-tag": "",
        "fee": 0E-18,
        "state": "confirmed",
        "created-at": 1637758163686,
        "updated-at": 1637758251559
    }
}
 */
