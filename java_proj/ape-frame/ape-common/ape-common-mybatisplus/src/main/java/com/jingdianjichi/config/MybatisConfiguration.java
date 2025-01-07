package com.jingdianjichi.config;

import com.jingdianjichi.inteceptor.SqlBeautyInterceptor;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class MybatisConfiguration {
    @Bean
    @ConditionalOnProperty(name = "sql.beauty.show", havingValue = "true", matchIfMissing = true)
    // 只有配置了sql.beauty.show=true才会生效
    public SqlBeautyInterceptor sqlBeautyInterceptor() {
        return new SqlBeautyInterceptor();
    }
}
