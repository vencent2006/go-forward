package com.vs.myworld.business.cardlist;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@NoArgsConstructor
@JsonInclude(JsonInclude.Include.NON_NULL)
public class CardStatus {
    @JsonProperty("card_status")
    private int cardStatus;
    private String title;
    private String desc;
    @JsonProperty("satisfy_show")
    private int satisfyShow;
    @JsonProperty("satisfy_statue")
    private int satisfyStatue;
    @JsonProperty("recommend_text")
    private String recommendText;
    private List<Settle> settle;
}
