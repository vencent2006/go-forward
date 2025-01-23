package com.jingdianjichi.user.designPattern.abstractFactoryPattern.easy;

public class Green implements Color {

    @Override
    public void fill() {
        System.out.println("Inside Green::fill() method.");
    }
}