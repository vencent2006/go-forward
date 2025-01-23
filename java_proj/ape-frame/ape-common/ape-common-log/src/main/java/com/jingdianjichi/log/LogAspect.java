package com.jingdianjichi.log;

import com.google.gson.Gson;
import lombok.extern.slf4j.Slf4j;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Pointcut;
import org.aspectj.lang.reflect.MethodSignature;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.stereotype.Component;

/**
 * 日志切面
 */
@Aspect
@Slf4j
@Component
@ConditionalOnProperty(name = {"log.aspect.enable"}, havingValue = "true", matchIfMissing = true) // 配置开关
public class LogAspect {
    /**
     * 切入点
     */
    @Pointcut("execution(* com.jingdianjichi.*.controller.*Controller.*(..)) || execution(* com.jingdianjichi.*.service.*Service.*(..))")
    private void pointCut() {

    }

    @Around("pointCut()")
    public void around(ProceedingJoinPoint pjp) throws Throwable {
        Object[] reqArgs = pjp.getArgs();
        String req = new Gson().toJson(reqArgs);
        MethodSignature methodSignature = (MethodSignature) pjp.getSignature();
        // 类名 + 方法名
        String methodName = methodSignature.getDeclaringType().getName() + "." + methodSignature.getName();
        log.info("{}.req:{}", methodName, req);
        Long startTime = System.currentTimeMillis();
        Object responseObj = pjp.proceed();// 不能catch业务自己的异常
        String resp = new Gson().toJson(responseObj);
        Long endTime = System.currentTimeMillis();
        log.info("{}.response:{},costTime:{}ms", methodName, resp, endTime - startTime);
    }
}
