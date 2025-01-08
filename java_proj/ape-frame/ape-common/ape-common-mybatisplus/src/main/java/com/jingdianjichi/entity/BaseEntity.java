package com.jingdianjichi.entity;

import com.baomidou.mybatisplus.annotation.FieldFill;
import com.baomidou.mybatisplus.annotation.TableField;
import com.baomidou.mybatisplus.annotation.TableLogic;
import lombok.Data;

import java.io.Serializable;
import java.util.Date;

// BaseEntity implements Serializable
// 这句话的意思是，BaseEntity 类实现了 Serializable 接口。
// 在 Java 中，Serializable 接口是一个标记接口，它没有任何方法或字段，其目的是标识一个类的实例可以被序列化。
// 序列化是将对象的状态转换为字节流的过程，以便可以将其存储在文件中或通过网络传输。
// 当一个类实现了 Serializable 接口时，
// 它的对象可以被序列化和反序列化，这在许多应用程序中是非常有用的，例如在网络通信、数据持久化和远程方法调用中。
@Data
public class BaseEntity implements Serializable {
    @TableField(fill = FieldFill.INSERT)
    private String createBy;

    @TableField(fill = FieldFill.INSERT)
    private Date createTime;

    @TableField(fill = FieldFill.UPDATE)
    private String updateBy;

    @TableField(fill = FieldFill.UPDATE)
    private Date updateTime;

    @TableField(fill = FieldFill.INSERT)
    @TableLogic
    private Integer deleteFlag;

    @TableField(fill = FieldFill.INSERT)
    private Integer version;
}
