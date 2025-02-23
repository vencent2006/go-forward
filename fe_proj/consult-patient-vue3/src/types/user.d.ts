export type User = {
  // token令牌
  token: string
  // 用户id
  id: string
  // 用户名称
  account: string
  // 手机号
  mobile: string
  // 头像
  avatar: string
}

export type CodeType = 'login' | 'register' | 'changeMobile' | 'forgetPassword' | 'bindMobile'
