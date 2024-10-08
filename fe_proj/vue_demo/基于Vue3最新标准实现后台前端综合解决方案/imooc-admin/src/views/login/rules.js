import i18n from '@/i18n'
export const validatePassword = () => {
  // 返回的是一个函数
  return (rule, value, callback) => {
    if (value.length < 6) {
      callback(new Error(i18n.global.t('msg.login.passwordRule')))
    } else {
      callback()
    }
  }
}
