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

// Omit 排除 | Pick 选择(摘取)
type OmitUser = Omit<User, 'token'> // 排除token的类型

export type UserInfo = OmitUser & {
  likeNumber: number // 关注数量
  collectionNumber: number // 收藏数量
  score: number // 积分
  couponNumber: number // 优惠券数量
  orderInfo: {
    paidNumber: number // 待付款
    receivedNumber: number // 待收货
    shippedNumber: number // 待发货
    finishedNumber: number // 已完成
  }
}

// 患者
export type Patient = {
  // 患者id
  id: string
  // 患者名称
  name: string
  // 患者身份证号
  idCard: string
  // 0 非默认 1 默认
  defaultFlag: 0 | 1
  // 患者性别 0 女 1 男
  gender: 0 | 1
  // 性别文字
  genderValue: string
  // 患者年龄
  age: number
}

// 家庭档案 - 患者信息列表
export type PatientList = Patient[]
