package com.vs.myworld.business.cardlist;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@NoArgsConstructor
@JsonInclude(JsonInclude.Include.NON_NULL)
public class Settle {
    @JsonProperty("settle_title")
    private String settleTitle;
    @JsonProperty("settle_type")
    private String settleType;
    @JsonProperty("settle_colour")
    private int settleColour;
    @JsonProperty("settle_desc")
    private String settleDesc;
    @JsonProperty("settle_btn_text")
    private String settleBtnText;
    @JsonProperty("settle_link")
    private String settleLink;
    @JsonProperty("settle_btn_text_statue")
    private int settleBtnTextStatue;
    @JsonProperty("recruit_info")
    private List<Object> recruitInfo;
    @JsonProperty("enter_info")
    private List<EnterInfo> enterInfo;
}
