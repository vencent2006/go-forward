package com.vs.myworld.business.aspect;

import cn.hutool.core.util.IdUtil;
import com.alibaba.fastjson.JSONObject;
import com.alibaba.fastjson.support.spring.PropertyPreFilters;
import lombok.extern.slf4j.Slf4j;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.Signature;
import org.aspectj.lang.annotation.*;
import org.slf4j.MDC;
import org.springframework.stereotype.Component;
import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;
import org.springframework.web.multipart.MultipartFile;

import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServletRequest;

@Slf4j
@Aspect
@Component
public class LogAspect {

    /*
    定义一个切点pointcut

    public * com.vs..*Controller.*(..)
    public：public方法
    第一个*: 表示所有的返回值
    com.vs..*Controller.*(..)：
    - ..*Controller：com.vs下的所有以Controller结尾
    - Controller后的.*：所有方法
    - (..)：所有参数
     */
    @Pointcut("execution(public * com.vs..*Controller.*(..))")
    public void pointcut() {

    }


    // 前置通知
    @Before("pointcut()")
    public void doBefore(JoinPoint joinPoint) {
        // log.info("前置通知");
    }

    // 环绕通知
    @Around("pointcut()")
    public Object doAround(ProceedingJoinPoint proceedingJoinPoint) throws Throwable {
        // 添加 LOG_ID
        MDC.put("LOG_ID", IdUtil.getSnowflakeNextIdStr()); // 雪花算法
        // MDC.put("LOG_ID", String.valueOf(System.currentTimeMillis()));

        log.info("------------ 环绕通知开始 ------------");
        long startTime = System.currentTimeMillis();
        // 开始打印请求日志
        ServletRequestAttributes attributes = (ServletRequestAttributes) RequestContextHolder.getRequestAttributes();
        HttpServletRequest request = attributes.getRequest();
        Signature signature = proceedingJoinPoint.getSignature();
        String name = signature.getName();

        // 打印请求信息
        log.info("请求地址: {} {}", request.getRequestURL().toString(), request.getMethod());
        log.info("类名方法: {} {}", signature.getDeclaringTypeName(), name);
        log.info("远程地址: {}", request.getRemoteAddr());

        // 打印请求参数
        Object[] args = proceedingJoinPoint.getArgs();

        // 排除特殊类型的参数，如文件类型
        Object[] arguments = new Object[args.length];
        for (int i = 0; i < args.length; i++) {
            if (args[i] instanceof ServletRequest
                    || args[i] instanceof ServletResponse
                    || args[i] instanceof MultipartFile) {
                continue;
            }
            arguments[i] = args[i];
        }
        // 排除字段，敏感字段或太长的字段不显示：身份证、手机号、邮箱、密码等
        String[] excludeProperties = {"cvv2", "idCard"};
        PropertyPreFilters filters = new PropertyPreFilters();
        PropertyPreFilters.MySimplePropertyPreFilter excludeFilter = filters.addFilter();
        excludeFilter.addExcludes(excludeProperties);
        log.info("请求参数: {}", JSONObject.toJSONString(arguments, excludeFilter));
        Object result = proceedingJoinPoint.proceed();
        // 排除字段，敏感字段或太长的字段不显示：身份证、手机号、邮箱、密码等
        log.info("返回结果: {}", JSONObject.toJSONString(result, excludeFilter));
        log.info("------------ 环绕通知结束 耗时: {} ms ------------", System.currentTimeMillis() - startTime);
        return result;
    }

    // 后置通知
    @After("pointcut()")
    public void doAfter(JoinPoint joinPoint) {
        // log.info("后置通知");
    }
}
