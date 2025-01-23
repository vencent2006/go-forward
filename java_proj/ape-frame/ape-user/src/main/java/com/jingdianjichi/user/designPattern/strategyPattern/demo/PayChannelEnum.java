package com.jingdianjichi.user.designPattern.strategyPattern.demo;

public enum PayChannelEnum {

    ZFB_PAY(0, "支付宝支付"),
    WX_PAY(1, "微信支付"),
    BANK_PAY(2, "银行卡支付"),
    ;

    private int code;

    private String desc;

    PayChannelEnum(int code, String desc) {
        this.code = code;
        this.desc = desc;
    }

    /**
     * 根据code值获取渠道枚举
     */
    public static PayChannelEnum getByCode(int codeVal) {
        for (PayChannelEnum payChannelEnum : PayChannelEnum.values()) {
            if (payChannelEnum.code == codeVal) {
                return payChannelEnum;
            }
        }
        return null;
    }

}
