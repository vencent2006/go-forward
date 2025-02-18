package com.vs.myworld.business.cardlist;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@JsonInclude(JsonInclude.Include.NON_NULL)
public class EnterInfo {
    private String title;
    private String icon;
    private String desc;
    @JsonProperty("is_satisfy")
    private int isSatisfy;
    @JsonProperty("certification_title")
    private String certificationTitle;
    @JsonProperty("certification_link")
    private String certificationLink;
}
