package com.jingdianjichi.redis.util;

import com.jingdianjichi.redis.exception.ShareLockException;
import org.apache.commons.lang.StringUtils;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import java.util.concurrent.TimeUnit;

@Component
public class RedisShareLockUtil {
    @Resource
    private RedisUtil redisUtil;

    // 自旋超时时间
    private long spinTimeout = 1000L;

    /**
     * 加锁
     *
     * @param lockKey
     * @param requestId
     * @param timeout
     * @return
     */
    public boolean lock(String lockKey, String requestId, long timeout) {
        // 1. 参数校验
        if (StringUtils.isBlank(lockKey) || StringUtils.isBlank(requestId) || timeout <= 0) {
            throw new ShareLockException("分布式锁-加锁参数异常");
        }
        long currentTime = System.currentTimeMillis();
        long outTime = currentTime + spinTimeout;
        Boolean result = false;
        // 2. 加锁可以自旋
        while (currentTime < outTime) {
            // 3. 借助redis的setnx进行锁的设置
            result = redisUtil.setNx(lockKey, requestId, timeout, TimeUnit.MILLISECONDS);
            if (result) {
                return true;
            }
            // 4. 自旋
            try {
                Thread.sleep(100);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            currentTime = System.currentTimeMillis();
        }
        return false;
    }

    /**
     * 解锁
     *
     * @param lockKey
     * @param requestId
     * @return
     */
    public boolean unlock(String lockKey, String requestId) {
        // 1. 参数校验
        if (StringUtils.isBlank(lockKey) || StringUtils.isBlank(requestId)) {
            throw new ShareLockException("分布式锁-解锁参数异常");
        }
        try {
            // 2. 解锁
            String value = redisUtil.get(lockKey);
            if (requestId.equals(value)) {
                redisUtil.del(lockKey);
                return true;
            }
        } catch (Exception e) {
            e.printStackTrace();
            // todo 记录日志
        }

        return false;
    }

    /**
     * 尝试加锁
     *
     * @param lockKey
     * @param requestId
     * @param timeout
     * @return
     */
    public boolean tryLock(String lockKey, String requestId, long timeout) {
        if (StringUtils.isBlank(lockKey) || StringUtils.isBlank(requestId) || timeout <= 0) {
            throw new ShareLockException("分布式锁-尝试加锁参数异常");
        }
        return redisUtil.setNx(lockKey, requestId, timeout, TimeUnit.MILLISECONDS);
    }
}
