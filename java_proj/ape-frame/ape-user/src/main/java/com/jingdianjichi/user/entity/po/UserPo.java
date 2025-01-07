package com.jingdianjichi.user.entity.po;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import com.jingdianjichi.entity.BaseEntity;
import lombok.Data;

@TableName("user")
@Data
public class UserPo extends BaseEntity {
    @TableId(value = "id", type = IdType.AUTO)
    private Long id;

    private String name;
    private Integer age;

}
