package com.vs.myworld.business.common;

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.PropertySource;
import org.springframework.context.annotation.PropertySources;

@Configuration
@ConfigurationProperties(prefix = "other")
@Data

@PropertySources({
        @PropertySource(value = "classpath:other.yml", factory = YamlPropertySourceFactory.class)
})
public class OtherYml {
    private String title;
    private String blog;
}
