package com.jingdianjichi.websocket.config;

import org.springframework.stereotype.Component;
import org.springframework.util.CollectionUtils;
import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;

import javax.servlet.http.HttpServletRequest;
import javax.websocket.HandshakeResponse;
import javax.websocket.server.HandshakeRequest;
import javax.websocket.server.ServerEndpointConfig;
import java.util.List;
import java.util.Map;

@Component
public class WebSocketServerConfig extends ServerEndpointConfig.Configurator {
    @Override
    public boolean checkOrigin(String originHeaderValue) {
        ServletRequestAttributes servletRequestAttributes = (ServletRequestAttributes) RequestContextHolder.getRequestAttributes();
        HttpServletRequest request = servletRequestAttributes.getRequest();
        // todo 校验
        return true;
    }

    @Override
    public void modifyHandshake(ServerEndpointConfig sec, HandshakeRequest request, HandshakeResponse response) {
        Map<String, List<String>> parameterMap = request.getParameterMap();
        List<String> erpList = parameterMap.get("erp");// erp是自己约定的值
        if (!CollectionUtils.isEmpty(erpList)) {
            // 把erp的值放到用户属性中
            sec.getUserProperties().put("erp", erpList.get(0));
        }
    }
}
