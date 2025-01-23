package com.jingdianjichi.user.designPattern.abstractFactoryPattern.easy;

public class Square implements Shape {

    @Override
    public void draw() {
        System.out.println("Inside Square::draw() method.");
    }
}