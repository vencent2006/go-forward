package com.jingdianjichi.user.service.impl;

import com.jingdianjichi.bean.PageResponse;
import com.jingdianjichi.user.dao.SysUserDao;
import com.jingdianjichi.user.entity.po.SysUser;
import com.jingdianjichi.user.entity.req.SysUserReq;
import com.jingdianjichi.user.service.SysUserService;
import org.springframework.beans.BeanUtils;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.List;

/**
 * (SysUser)表服务实现类
 *
 * @author makejava
 * @since 2025-01-11 00:35:41
 */
@Service("sysUserService")
public class SysUserServiceImpl implements SysUserService {
    @Resource
    private SysUserDao sysUserDao;

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    @Override
    public SysUser queryById(Long id) {
        return this.sysUserDao.queryById(id);
    }


    /**
     * 分页查询
     *
     * @param sysUserReq 请求
     * @return 查询结果
     */
    @Override
    public PageResponse<SysUser> queryByPage(SysUserReq sysUserReq) {
        SysUser sysUser = new SysUser();
        BeanUtils.copyProperties(sysUserReq, sysUser);
        PageResponse<SysUser> pageResponse = new PageResponse<>();
        pageResponse.setCurrent(sysUserReq.getPageNo());
        pageResponse.setPageSize(sysUserReq.getPageSize());
        Long pageStart = (sysUserReq.getPageNo() - 1) * sysUserReq.getPageSize();
        long total = this.sysUserDao.count(sysUser);
        List<SysUser> sysUserList = this.sysUserDao.queryAllByLimit(sysUser, pageStart, sysUserReq.getPageSize());
        pageResponse.setTotal(total);
        pageResponse.setRecords(sysUserList);
        return pageResponse;
    }

    /**
     * 新增数据
     *
     * @param sysUser 实例对象
     * @return 实例对象
     */
    @Override
    public SysUser insert(SysUser sysUser) {
        this.sysUserDao.insert(sysUser);
        return sysUser;
    }

    /**
     * 修改数据
     *
     * @param sysUser 实例对象
     * @return 实例对象
     */
    @Override
    public SysUser update(SysUser sysUser) {
        this.sysUserDao.update(sysUser);
        return this.queryById(sysUser.getId());
    }

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    @Override
    public boolean deleteById(Long id) {
        return this.sysUserDao.deleteById(id) > 0;
    }
}
