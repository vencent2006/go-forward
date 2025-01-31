package com.jingdianjichi.user.event;

import lombok.extern.slf4j.Slf4j;
import org.springframework.context.ApplicationEventPublisher;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;

@Service
@Slf4j
public class PersonEventService {
    @Resource
    private ApplicationEventPublisher applicationEventPublisher;

    public void createPerson(Person person) {
        PersonChangeEvent personChangeEvent = new PersonChangeEvent(person, "create");
        applicationEventPublisher.publishEvent(personChangeEvent);
    }
}
