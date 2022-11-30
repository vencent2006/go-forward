package com.vincent.controller;

import com.vincent.pojo.DbStu;
import com.vincent.pojo.bo.DbStuBO;
import com.vincent.service.StuService;
import com.vincent.utils.JSONResult;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.validation.BindingResult;
import org.springframework.validation.FieldError;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import javax.validation.Valid;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.UUID;

@RestController
@RequestMapping("dbstu")
@Slf4j
public class DbStuController {

    @Autowired
    private StuService stuService;


    @GetMapping("{stuId}/get")
    public Object get(@PathVariable("stuId") String stuId,
                      @RequestParam Integer id,
                      @RequestParam String name){
        /**
         * @RequestParam 用于获取url中的请求参数，如果参数变量名保持一致，该注解可以省略
         */
        log.info("stuId = " + stuId);
        log.info("id = " + id);
        log.info("name = " + name);

        return "查询";
    }

    @GetMapping("get")
    public JSONResult get(@RequestParam String stuId){
        DbStu stu = stuService.queryById(stuId);
        return JSONResult.ok(stu);
    }

    @GetMapping("list")
    public JSONResult list(@RequestParam String name, @RequestParam Integer sex){
        List<DbStu> stuList = stuService.queryByCondition(name, sex);
        return JSONResult.ok(stuList);
    }

    @PostMapping("create")
    public JSONResult createStu(){
        String sid = UUID.randomUUID().toString();
        DbStu stu = new DbStu();
        stu.setId(sid);
        stu.setName("mike");
        stu.setSex(1);

        stuService.saveStu(stu);

        return JSONResult.ok();
    }

//curl --location --request POST 'http://localhost:8090/dbstu/create2' \
// --header 'Content-Type: application/json' \
// --header 'Cookie: clientId=123456' \
// --data-raw '{
//     "name": "lucas",
//     "sex": 0,
//     "grade": "0",
//     "classroom": "19",
//     "skill": ["java", "html", "ios", "大数据", "jsp", "html"],
//     "englishName": "ab",
//     "email": "dfasfdsfdsa"
// }'

    @PostMapping("create2")
    public JSONResult createStu2(@Valid @RequestBody DbStuBO stuBO,
                                 BindingResult result){
        // 判断BindingResult是否有错误，错误信息会包含在里面，如果有则直接return
        if (result.hasErrors()){
            Map<String, String> map = getErrors(result);
            return JSONResult.errorMap(map);
        }
        DbStu stu = new DbStu();
        BeanUtils.copyProperties(stuBO, stu);
        String sid = UUID.randomUUID().toString();
        stu.setId(sid);

        stuService.saveStu(stu);

        return JSONResult.ok();
    }


    @PostMapping("update")
    public JSONResult update(){
        DbStu stu = new DbStu();
        stu.setId("af789438-b2d0-4679-9693-a2459f7f79d1");
        stu.setName("michelle");
        stu.setSex(0);

        stuService.updateStu(stu);

        return JSONResult.ok();
    }

    @PostMapping("delete")
    public JSONResult delete(){
        DbStu stu = new DbStu();
        // stu.setId("fdc21cf8-7908-434c-91b5-227b3dcb1adc");
        // stu.setSex(2);
        stu.setName("lewis");
        stuService.deleteStu(stu);
        return JSONResult.ok();
    }


    public Map<String, String> getErrors(BindingResult result) {
        Map<String, String> map = new HashMap<>();
        List<FieldError> errorList = result.getFieldErrors();
        for (FieldError error : errorList){
            String field = error.getField();
            String msg = error.getDefaultMessage();
            map.put(field, msg);
        }
        return map;
    }
}
