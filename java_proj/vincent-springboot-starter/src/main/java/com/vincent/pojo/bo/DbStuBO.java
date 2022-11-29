package com.vincent.pojo.bo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.ToString;


// BO means: business object

@Data
@ToString
@AllArgsConstructor
public class DbStuBO {
    private String id;
    private String name;
    private Integer sex;
}
