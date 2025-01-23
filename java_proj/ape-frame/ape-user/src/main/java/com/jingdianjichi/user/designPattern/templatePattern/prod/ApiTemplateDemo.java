package com.jingdianjichi.user.designPattern.templatePattern.prod;

import com.jingdianjichi.bean.Result;

public class ApiTemplateDemo {
    public static void main(String[] args) {
        ApiTemplate apiTemplate = new ApiTemplate();
        Result result = Result.ok();
        apiTemplate.execute(result, new Action() {
            @Override
            public void validate() {
                System.out.println("参数校验");
            }

            @Override
            public void execute() {
                System.out.println("执行");
            }

            @Override
            public void after() {
                System.out.println("后续执行");
            }
        });
        System.out.println(result);
    }
}
