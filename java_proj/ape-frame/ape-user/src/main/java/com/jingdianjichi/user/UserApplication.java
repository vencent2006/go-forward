package com.jingdianjichi.user;

import lombok.extern.slf4j.Slf4j;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@Slf4j
@SpringBootApplication
public class UserApplication {
    public static void main(String[] args) {
        try {
            SpringApplication.run(UserApplication.class);
        } catch (Exception e) {
            log.error("Application failed to start", e);
        }
    }
}
