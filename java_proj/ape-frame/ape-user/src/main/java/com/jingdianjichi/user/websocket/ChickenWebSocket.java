package com.jingdianjichi.user.websocket;

import com.jingdianjichi.websocket.config.WebSocketServerConfig;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

import javax.websocket.*;
import javax.websocket.server.ServerEndpoint;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicInteger;

@Slf4j
@Component
@ServerEndpoint(value = "/chicken/socket", configurator = WebSocketServerConfig.class)
public class ChickenWebSocket {

    /**
     * 当前在线的连接数
     */
    private static AtomicInteger onlineCount = new AtomicInteger(0);
    /**
     * 存放所有在线的客户端
     */
    private static Map<String, ChickenWebSocket> clients = new ConcurrentHashMap<>();
    /**
     * 某个客户端连接的session会话
     */
    private Session session;
    /**
     * 标识当前会话的key
     */
    private String erp = "";


    @OnOpen
    public void onOpen(Session session, EndpointConfig config) throws Exception {
        Map<String, Object> userProperties = config.getUserProperties();
        String erp = (String) userProperties.get("erp");
        this.erp = erp;
        this.session = session;
        if (clients.containsKey(erp)) {
            clients.get(this.erp).session.close();// 关闭之前的连接
            clients.remove(this.erp);// 移除之前的连接
            onlineCount.decrementAndGet();// 在线数减1
        }
        clients.put(this.erp, this);// 加入map中
        onlineCount.incrementAndGet();// 在线数加1
        log.info("有新连接加入:{}, 当前在线人数为：{}", this.erp, onlineCount.get());
        sendMessage("连接成功", this.session);
    }

    @OnMessage
    public void onMessage(String message, Session session) throws Exception {
        log.info("服务端收到：来自客户端:{}, 的消息:{}", this.erp, message);
        // todo 心跳检测
        if (message.equals("ping")) {
            sendMessage("pong", session);
        }
    }

    @OnClose
    public void onClose() throws Exception {
        if (clients.containsKey(erp)) {
            clients.get(erp).session.close();
            clients.remove(erp);
            onlineCount.decrementAndGet();
        }
        log.info("有一处连接关闭:{}, 当前在线人数:{}", this.erp, onlineCount.get());
    }

    @OnError
    public void onError(Session session, Throwable throwable) throws Exception {
        // todo 为啥是3个输入项呢
        log.error("websocket:{}, 发送错误, 错误原因:{}", this.erp, throwable.getMessage(), throwable);
        session.close();
    }

    public void sendMessage(String message, Session session) throws Exception {
        log.info("服务端给客户端:{}, 发送消息:{}", this.erp, message);
        session.getBasicRemote().sendText(message);
    }

    public void sendMessage(String message) throws Exception {
        for (Map.Entry<String, ChickenWebSocket> sessionEntry : clients.entrySet()) {
            String erp = sessionEntry.getKey();
            ChickenWebSocket webSocket = sessionEntry.getValue();
            Session session = webSocket.session;
            sendMessage(message, session);
        }
    }

}
