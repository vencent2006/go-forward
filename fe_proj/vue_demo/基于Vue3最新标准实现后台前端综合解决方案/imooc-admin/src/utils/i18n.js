import i18n from '@/i18n'

/**
 * 将title转化为国际化内容
 * @param title
 * @returns {TranslateResult}
 */
export function generateTitle(title) {
  return i18n.global.t('msg.route.' + title)
}
