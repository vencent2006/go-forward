package com.jingdianjichi.user.designPattern.abstractFactoryPattern.easy;

public abstract class AbstractFactory {
    public abstract Color getColor(String color);

    public abstract Shape getShape(String shape);
}