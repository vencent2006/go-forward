package com.jingdianjichi.user.designPattern.abstractFactoryPattern.easy;

public class Circle implements Shape {

    @Override
    public void draw() {
        System.out.println("Inside Circle::draw() method.");
    }
}