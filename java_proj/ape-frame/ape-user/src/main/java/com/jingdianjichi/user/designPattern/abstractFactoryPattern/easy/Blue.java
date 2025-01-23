package com.jingdianjichi.user.designPattern.abstractFactoryPattern.easy;

public class Blue implements Color {

    @Override
    public void fill() {
        System.out.println("Inside Blue::fill() method.");
    }
}