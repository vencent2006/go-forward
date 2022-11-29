package com.vincent.pojo.bo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.ToString;

import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotNull;


// BO means: business object

@Data
@ToString
@AllArgsConstructor
public class DbStuBO {
    private String id;
    @NotBlank
    private String name;
    @NotNull
    private Integer sex;
}
