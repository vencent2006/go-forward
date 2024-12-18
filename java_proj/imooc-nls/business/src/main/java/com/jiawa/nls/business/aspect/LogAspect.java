package com.jiawa.nls.business.aspect;

import lombok.extern.slf4j.Slf4j;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.*;
import org.springframework.stereotype.Component;

@Slf4j
@Aspect
@Component
public class LogAspect {

    /*
    定义一个切点pointcut

    public * com.jiawa..*Controller.*(..)
    public：public方法
    第一个*: 表示所有的返回值
    com.jiawa..*Controller.*(..)：
    - ..*Controller：com.jiawa下的所有以Controller结尾
    - Controller后的.*：所有方法
    - (..)：所有参数
     */
    @Pointcut("execution(public * com.jiawa..*Controller.*(..))")
    public void pointcut(){

    }


    // 前置通知
    @Before("pointcut()")
    public void doBefore(JoinPoint joinPoint){
        log.info("前置通知");
    }

    // 环绕通知
    @Around("pointcut()")
    public Object doAround(ProceedingJoinPoint proceedingJoinPoint) throws Throwable {
        log.info("环绕通知开始");
        Object result = proceedingJoinPoint.proceed();
        log.info("环绕通知结束");
        return result;
    }

    // 后置通知
    @After("pointcut()")
    public void doAfter(JoinPoint joinPoint){
        log.info("后置通知");
    }
}
