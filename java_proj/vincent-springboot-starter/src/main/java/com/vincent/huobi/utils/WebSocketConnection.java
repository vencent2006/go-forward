package com.vincent.huobi.utils;

import com.vincent.huobi.constant.enums.ConnectionStateEnum;

public interface WebSocketConnection {

  ConnectionStateEnum getState();

  Long getConnectionId();

  void reConnect();

  void reConnect(int delayInSecond);

  long getLastReceivedTime();

  void send(String str);
}
