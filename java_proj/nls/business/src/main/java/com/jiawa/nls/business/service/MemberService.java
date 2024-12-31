package com.jiawa.nls.business.service;

import cn.hutool.core.collection.CollUtil;
import com.jiawa.nls.business.domain.Member;
import com.jiawa.nls.business.domain.MemberExample;
import com.jiawa.nls.business.mapper.MemberMapper;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.List;

@Service
public class MemberService {
    @Resource
    private MemberMapper memberMapper;

    /**
     * 按照手机号查询用户
     *
     * @param mobile 手机号
     * @return 用户 Member
     */
    public Member selectByMobile(String mobile) {
        MemberExample memberExample = new MemberExample();
        MemberExample.Criteria criteria = memberExample.createCriteria();
        criteria.andMobileEqualTo(mobile);
        List<Member> list = memberMapper.selectByExample(memberExample);
        if (CollUtil.isNotEmpty(list)) {
            return list.get(0);
        }
        return null;
    }
}
