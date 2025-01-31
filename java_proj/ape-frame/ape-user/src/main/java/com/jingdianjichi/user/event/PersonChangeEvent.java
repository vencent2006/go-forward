package com.jingdianjichi.user.event;

import lombok.Getter;
import org.springframework.context.ApplicationEvent;

@Getter
public class PersonChangeEvent extends ApplicationEvent {
    private Person person;
    private String operateType;

    public PersonChangeEvent(Person person, String operateType) {
        super(person);
        this.person = person;
        this.operateType = operateType;
    }
}
