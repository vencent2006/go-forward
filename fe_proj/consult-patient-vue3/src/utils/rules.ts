// 提供校验规则

import type { FieldRule } from 'vant'

// 校验手机号
const mobileRules: FieldRule[] = [
  { required: true, message: '请输入手机号' },
  { pattern: /^1[3-9]\d{9}$/, message: '手机号格式不正确' },
]

//
const passwordRules: FieldRule[] = [
  { required: true, message: '请输入密码' },
  { pattern: /^\w{8,24}$/, message: '密码必须是8-24位' },
]

// 校验验证码
const codeRules: FieldRule[] = [
  { required: true, message: '请输入验证码' },
  { pattern: /^\d{6}$/, message: '验证码必须是6位数字' },
]

export { mobileRules, passwordRules, codeRules }
