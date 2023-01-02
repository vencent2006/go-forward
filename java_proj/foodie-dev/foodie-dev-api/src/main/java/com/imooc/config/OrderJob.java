package com.imooc.config;

import com.imooc.service.OrderService;
import com.imooc.utils.DateUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

@Component
public class OrderJob {
    static final Logger logger = LoggerFactory.getLogger(OrderJob.class);

    private OrderService orderService;

    // https://cron.qqe2.com 查询cron配置
    // @Scheduled(cron = "0/3 * * * * ?") // 每3秒运行一次
    @Scheduled(cron = "* * 0/1 * * ?") // 每小时运行一次
    public void autoCloseOrder() {
        logger.info("执行定时任务，当前时间为:" + DateUtil.getCurrentDateString(DateUtil.DATETIME_PATTERN));
        orderService.closeOrder();
    }
}
