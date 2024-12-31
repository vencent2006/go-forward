package com.jiawa.nls.business.service;

import cn.hutool.core.collection.CollUtil;
import cn.hutool.core.util.IdUtil;
import com.jiawa.nls.business.domain.Member;
import com.jiawa.nls.business.domain.MemberExample;
import com.jiawa.nls.business.exception.BusinessException;
import com.jiawa.nls.business.exception.BusinessExceptionEnum;
import com.jiawa.nls.business.mapper.MemberMapper;
import com.jiawa.nls.business.req.MemberRegisterReq;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.Date;
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

    /**
     * 注册
     */
    public void register(MemberRegisterReq req) {
        String mobile = req.getMobile();
        Date now = new Date();

        // 校验：手机号是否已注册
        Member memberDB = selectByMobile(req.getMobile());
        if (memberDB != null) {
            throw new BusinessException(BusinessExceptionEnum.MEMBER_MOBILE_HAD_REGISTER);
        }
        
        // 插入用户
        Member member = new Member();
        member.setId(IdUtil.getSnowflakeNextId());
        member.setMobile(mobile);
        member.setPassword(req.getPassword());
        member.setName(mobile.substring(0, 3) + "****" + mobile.substring(7));// 手机号中间四位为*
        member.setCreateAt(now);
        member.setUpdateAt(now);
        memberMapper.insert(member);


    }
}
