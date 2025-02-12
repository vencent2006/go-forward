package com.vs.myworld.business.service;

import cn.hutool.core.bean.BeanUtil;
import com.vs.myworld.business.domain.World;
import com.vs.myworld.business.domain.WorldExample;
import com.vs.myworld.business.mapper.WorldMapper;
import com.vs.myworld.business.resp.WorldQueryResp;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.List;

@Slf4j
@Service
public class WorldService {
    @Resource
    private WorldMapper worldMapper;

    public List<WorldQueryResp> query() {
        WorldExample worldExample = new WorldExample();
        WorldExample.Criteria criteria = worldExample.createCriteria();
        criteria.andStatusIsNotNull();
        List<World> list = worldMapper.selectByExample(worldExample);
        return BeanUtil.copyToList(list, WorldQueryResp.class);
    }
}
