package com.jingdianjichi.user.event;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.event.TransactionalEventListener;

@Service
@Slf4j
public class PersonEventListener {
    // 这个注解必须打上
    @TransactionalEventListener(fallbackExecution = true)
    public void listenCreateEvent(PersonChangeEvent personChangeEvent) {
        switch (personChangeEvent.getOperateType()) {
            case "create":
                // todo 大家应该写自己的逻辑
                log.info("执行创建事件:{}", JSON.toJSONString(personChangeEvent.getPerson()));
                break;
            default:
                break;
        }
    }
}
