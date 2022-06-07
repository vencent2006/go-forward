<template>
  <div class="mixin-components-container">
    <el-row>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>创建HTTP服务</span>
        </div>
        <el-form ref="form" :model="form" label-width="140px">
          <el-form-item label="服务名称" class="is-required">
            <el-input v-model="form.service_name" placeholder="6-128位数字字母下划线" />
          </el-form-item>
          <el-form-item label="服务描述" class="is-required">
            <el-input v-model="form.service_desc" placeholder="最多255字符必填" />
          </el-form-item>
          <el-form-item label="接入类型" class="is-required">
            <el-input v-model="form.rule" placeholder="路径格式:/user/，域名格式:www.test.com" class="input-with-select">
              <el-select slot="prepend" v-model="form.rule_type" placeholder="请选择" style="width:80px">
                <el-option label="路径" :value="0" />
                <el-option label="域名" :value="1" />
              </el-select>
              <el-button slot="append" icon="el-icon-search" />
            </el-input>
          </el-form-item>
          <el-form-item label="支持HTTPS">
            <el-switch
              v-model="form.need_https"
              :active-value="1"
              :inactive-value="0"
            />
            <span style="color: #606626;font-weight: 700;">&nbsp;&nbsp;支持strip_uri&nbsp;&nbsp;</span>
            <el-switch
              v-model="form.need_strip_uri"
              :active-value="1"
              :inactive-value="0"
            />
            <span style="color: #606626;font-weight: 700;">&nbsp;&nbsp;支持websocket&nbsp;&nbsp;</span>
            <el-switch
              v-model="form.need_websocket"
              :active-value="1"
              :inactive-value="0"
            />
          </el-form-item>

          <el-form-item label="URL重写">
            <el-input v-model="form.url_rewrite" type="textarea" autosize placeholder="格式:^/gateway/test_service(.*) $1 多条换行" />
          </el-form-item>
          <el-form-item label="HEADER转换">
            <el-input v-model="form.header_transfor" type="textarea" autosize placeholder="header转换 支持add(增加)/del(删除)/edit(修改) 格式: add headerName headerValue" />
          </el-form-item>
          <el-form-item label="开启验证">
            <el-switch
              v-model="form.open_auth"
              :active-value="1"
              :inactive-value="0"
            />
          </el-form-item>
          <el-form-item label="IP白名单">
            <el-input v-model="form.white_list" type="textarea" autosize placeholder="格式: 127.0.0.1 多条换行，白名单优先于黑名单" />
          </el-form-item>
          <el-form-item label="IP黑名单">
            <el-input v-model="form.black_list" type="textarea" autosize placeholder="格式: 127.0.0.1 多条换行，白名单优先于黑名单" />
          </el-form-item>
          <el-form-item label="客户端限流">
            <el-input v-model="form.clientip_flow_limit" placeholder="0表示无限制" />
          </el-form-item>
          <el-form-item label="服务端限流">
            <el-input v-model="form.service_flow_limit" placeholder="0表示无限制" />
          </el-form-item>
          <el-form-item label="轮询方式">
            <el-radio-group v-model="form.round_type">
              <el-radio :label="0">random</el-radio>
              <el-radio :label="1">round-robin</el-radio>
              <el-radio :label="2">weight-round-robin</el-radio>
              <el-radio :label="3">ip_hash</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="IP列表" class="is-required">
            <el-input v-model="form.ip_list" type="textarea" autosize placeholder="格式: 127.0.0.1:80 多条换行" />
          </el-form-item>
          <el-form-item label="权重列表" class="is-required">
            <el-input v-model="form.weight_list" type="textarea" autosize placeholder="格式: 50 多条换行" />
          </el-form-item>
          <el-form-item label="建立连接超时">
            <el-input v-model="form.upstream_connect_timeout" placeholder="单位:s, 0表示无限制" />
          </el-form-item>
          <el-form-item label="获取header超时">
            <el-input v-model="form.upstream_header_timeout" placeholder="单位:s, 0表示无限制" />
          </el-form-item>
          <el-form-item label="连接最大空闲时间">
            <el-input v-model="form.upstream_idle_timeout" placeholder="0表示无限制" />
          </el-form-item>
          <el-form-item label="最大空闲连接数">
            <el-input v-model="form.upstream_max_idle" placeholder="0表示无限制" />
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              :disabled="submitButtonDisabled"
              @click="handleCreate"
            >立即创建</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </el-row>
  </div>
</template>

<script>
import { serviceAddHttp } from '@/api/service'
export default {
  name: 'ServiceCreateHttp',
  data() {
    return {
      submitButtonDisabled: false,
      form: {
        service_name: '',
        service_desc: '',
        rule_type: 0,
        rule: '',
        need_https: 0,
        need_websocket: 0,
        need_strip_uri: 1,
        url_rewrite: '',
        header_transfor: '',
        round_type: 2,
        ip_list: '',
        weight_list: '',
        upstream_connect_timeout: '',
        upstream_header_timeout: '',
        upstream_idle_timeout: '',
        upstream_max_idle: '',
        open_auth: 0,
        white_list: '',
        black_list: '',
        clientip_flow_limit: '',
        service_flow_limit: ''
      }
    }
  },
  methods: {
    handleCreate() {
      // 禁用按钮
      this.submitButtonDisabled = true

      const query = Object.assign({}, this.form) // {}表示浅拷贝
      console.log(query)
      query.url_rewrite = query.url_rewrite.replace(/\n/g, ',')
      query.header_transfor = query.header_transfor.replace(/\n/g, ',')
      query.ip_list = query.ip_list.replace(/\n/g, ',')
      query.black_list = query.black_list.replace(/\n/g, ',')
      query.weight_list = query.weight_list.replace(/\n/g, ',')
      query.clientip_flow_limit = Number(query.clientip_flow_limit)
      query.upstream_connect_timeout = Number(query.upstream_connect_timeout)
      query.upstream_header_timeout = Number(query.upstream_header_timeout)
      query.upstream_idle_timeout = Number(query.upstream_idle_timeout)
      query.upstream_max_idle = Number(query.upstream_max_idle)
      query.service_flow_limit = Number(query.service_flow_limit)
      serviceAddHttp(query).then(response => {
        this.submitButtonDisabled = false
        this.$notify({
          title: 'Success',
          message: '添加成功',
          type: 'success',
          duration: 2000
        })
      }).catch(() => {
        this.submitButtonDisabled = false
      })
    }
  }
}
</script>

<style scoped>
.mixin-components-container {
  background-color: #f0f2f5;
  padding: 30px;
  min-height: calc(100vh - 84px);
}
.component-item{
  min-height: 100px;
}
</style>
