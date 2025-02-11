package com.vs.myworld.business.util;

import cn.hutool.core.date.DateField;
import cn.hutool.core.date.DateTime;
import cn.hutool.crypto.GlobalBouncyCastleProvider;
import cn.hutool.json.JSONObject;
import cn.hutool.jwt.JWT;
import cn.hutool.jwt.JWTPayload;
import cn.hutool.jwt.JWTUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.HashMap;
import java.util.Map;

public class JwtUtil {
    private static final Logger LOG = LoggerFactory.getLogger(JwtUtil.class);

    /**
     * 盐值很重要，不能泄露，且每个项目都应该不一样，可以放在配置文件中
     */
    private static final String key = "vsmyworld";

    public static String createLoginToken(Map<String, Object> map) {
        LOG.info("准备生成登录token, map: {}", map);
        return createToken(map, 24 * 60);
    }

    public static String createToken(Map<String, Object> map, Integer expMinute) {
        LOG.info("开始生成JWT token, map: {}", map);
        GlobalBouncyCastleProvider.setUseBouncyCastle(false);
        DateTime now = DateTime.now();
        DateTime expTime = now.offsetNew(DateField.MINUTE, expMinute);
        Map<String, Object> payload = new HashMap<>();
        // 签发时间
        payload.put(JWTPayload.ISSUED_AT, now);
        // 过期时间
        payload.put(JWTPayload.EXPIRES_AT, expTime);
        // 生效时间
        payload.put(JWTPayload.NOT_BEFORE, now);
        // 载荷(内容)
        payload.putAll(map);
        String token = JWTUtil.createToken(payload, key.getBytes());
        LOG.info("生成成功JWT token: {}", token);
        return token;
    }

    public static boolean validate(String token) {
        LOG.info("开始JWT token校验, token: {}", token);
        GlobalBouncyCastleProvider.setUseBouncyCastle(false);
        JWT jwt = JWTUtil.parseToken(token).setKey(key.getBytes());
        // validate包含了verify
        boolean validate = jwt.validate(0);
        LOG.info("JWT token校验结果: {}", validate);
        return validate;
    }

    public static JSONObject getJsonObject(String token) {
        LOG.info("开始JWT token解析, token: {}", token);
        GlobalBouncyCastleProvider.setUseBouncyCastle(false);
        JWT jwt = JWTUtil.parseToken(token).setKey(key.getBytes());
        JSONObject payloads = jwt.getPayloads();
        payloads.remove(JWTPayload.ISSUED_AT);
        payloads.remove(JWTPayload.EXPIRES_AT);
        payloads.remove(JWTPayload.NOT_BEFORE);
        LOG.info("根据token获取原始内容: {}", payloads);
        return payloads;
    }

    public static void main(String[] args) {
        Map<String, Object> map = new HashMap<>();
        map.put("name", "136****8888");
        String token = createLoginToken(map);

        // String token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJuYmYiOjE3MzU4MzUwMDEsIm5hbWUiOiIxMzYqKioqODg4OCIsImV4cCI6MTczNTkyMTQwMSwiaWF0IjoxNzM1ODM1MDAxfQ.qQfTWD66uLkcxtZlWRdJ2ac_dPZO1gfoytFaLdsakI0";
        validate(token);

        getJsonObject(token);
    }
}
