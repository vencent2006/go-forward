package com.jingdianjichi.user;

import lombok.extern.slf4j.Slf4j;
import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;

@Slf4j
@SpringBootApplication
@MapperScan("com.jingdianjichi.*.mapper")
@ComponentScan("com.jingdianjichi")
public class UserApplication {
    public static void main(String[] args) {
        try {
            SpringApplication.run(UserApplication.class);
        } catch (Exception e) {
            log.error("Application failed to start", e);
        }
    }
}
