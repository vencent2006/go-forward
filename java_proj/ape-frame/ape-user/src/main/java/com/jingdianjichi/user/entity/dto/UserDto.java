package com.jingdianjichi.user.entity.dto;

import lombok.Data;

@Data
public class UserDto {
    private String name;
    private Integer age;

    private Integer pageIndex;
    private Integer pageSize;
}
