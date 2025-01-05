package com.jingdianjichi.handler;

import com.baomidou.mybatisplus.core.handlers.MetaObjectHandler;
import lombok.extern.slf4j.Slf4j;
import org.apache.ibatis.reflection.MetaObject;
import org.springframework.stereotype.Component;

import java.util.Date;

@Slf4j
@Component
public class MyMetaObjectHandler implements MetaObjectHandler {
    @Override
    public void insertFill(MetaObject metaObject) {
        Date now = new Date();
        this.strictInsertFill(metaObject, "createBy", String.class, "jingdianjichi");
        this.strictInsertFill(metaObject, "createTime", Date.class, now);
        this.strictUpdateFill(metaObject, "updateBy", String.class, "jingdianjichi");
        this.strictInsertFill(metaObject, "updateTime", Date.class, now);
        this.strictInsertFill(metaObject, "deleteFlag", Integer.class, 0);
        this.strictInsertFill(metaObject, "version", Integer.class, 0);
    }

    @Override
    public void updateFill(MetaObject metaObject) {
        this.strictUpdateFill(metaObject, "updateBy", String.class, "jingdianjichi");
        this.strictUpdateFill(metaObject, "updateTime", Date.class, new Date());
    }
}
