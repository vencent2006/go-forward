package com.jingdianjichi.user.designPattern.abstractFactoryPattern.easy;

public class Rectangle implements Shape {

    @Override
    public void draw() {
        System.out.println("Inside Rectangle::draw() method.");
    }
}