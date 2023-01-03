package com.vincent.test;

import com.alibaba.fastjson.JSON;
import com.vincent.pojo.bo.DbStuBO;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.*;


@SpringBootTest
@AutoConfigureMockMvc
public class WebTest {

    @Autowired
    private MockMvc mockMvc;


    @Test
    void should_get_hello() throws Exception {
        mockMvc.perform(get("/hello"))
                .andDo(print())
                .andExpect(status().isOk());
    }


    @Test
    void dbstu_get() throws Exception {

        String stuId = "ccf63f9d-67bf-4955-8de9-98e367d86a91";
        mockMvc.perform(get("/dbstu/get").queryParam("stuId", stuId))
                .andDo(print())
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.data.id").value(stuId));
    }

    @Test
    void dbstu_createStu2() throws Exception {
        DbStuBO dbStuBO = new DbStuBO();
        dbStuBO.setClassroom(1);
        dbStuBO.setSex(1);
        dbStuBO.setGrade(1);
        dbStuBO.setName("will");
        dbStuBO.setEnglishName("william");
        dbStuBO.setEmail("abc@1234.com");

        Object jsonStr = JSON.toJSON(dbStuBO);

        System.out.println(jsonStr.toString());

        mockMvc.perform(post("/dbstu/create2").contentType(MediaType.APPLICATION_JSON).content(jsonStr.toString()))
                .andDo(print())
                .andExpect(status().isOk());

    }

}
