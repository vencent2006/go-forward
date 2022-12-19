package com.vincent.huobi.service.huobi.parser.trade;

import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import com.vincent.huobi.constant.enums.StopOrderOperatorEnum;
import com.vincent.huobi.model.trade.Order;
import com.vincent.huobi.service.huobi.parser.HuobiModelParser;

public class OrderParser implements HuobiModelParser<Order> {

  @Override
  public Order parse(JSONObject json) {

    BigDecimal filledAmount = json.getBigDecimal("field-amount");
    if (filledAmount == null) {
      filledAmount = json.getBigDecimal("filled-amount");
    }

    BigDecimal filledCashAmount = json.getBigDecimal("field-cash-amount");
    if (filledCashAmount == null) {
      filledCashAmount = json.getBigDecimal("filled-cash-amount");
    }

    BigDecimal filledFees = json.getBigDecimal("field-fees");
    if (filledFees == null) {
      filledFees = json.getBigDecimal("filled-fees");
    }

    Order order = json.toJavaObject(Order.class);

    // 添加clientOrderId
    String clientOrderId = json.getString("client-order-id");
    if (null != clientOrderId) {
      order.setClientOrderId(clientOrderId);
    }

    order.setAccountId(json.getLong("account-id"));
    order.setType(json.getString("type"));

    order.setFilledAmount(filledAmount);
    order.setFilledCashAmount(filledCashAmount);
    order.setFilledFees(filledFees);

    order.setSource(json.getString("source"));
    order.setState(json.getString("state"));

    order.setStopPrice(json.getBigDecimal("stop-price"));
    order.setOperator(StopOrderOperatorEnum.find(json.getString("operator")));

    order.setCreatedAt(json.getLong("created-at"));
    order.setCanceledAt(json.getLong("canceled-at"));
    order.setFinishedAt(json.getLong("finished-at"));
    return order;
  }

  @Override
  public Order parse(JSONArray json) {
    return null;
  }

  @Override
  public List<Order> parseArray(JSONArray jsonArray) {

    if (jsonArray == null || jsonArray.size() <= 0) {
      return new ArrayList<>();
    }
    List<Order> list = new ArrayList<>(jsonArray.size());
    for (int i = 0; i < jsonArray.size(); i++) {
      JSONObject jsonObject = jsonArray.getJSONObject(i);
      list.add(parse(jsonObject));
    }
    return list;
  }
}
