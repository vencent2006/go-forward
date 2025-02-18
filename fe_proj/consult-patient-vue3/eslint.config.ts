import pluginVue from 'eslint-plugin-vue'
import { defineConfigWithVueTs, vueTsConfigs } from '@vue/eslint-config-typescript'
import skipFormatting from '@vue/eslint-config-prettier/skip-formatting'

// To allow more languages other than `ts` in `.vue` files, uncomment the following lines:
// import { configureVueProject } from '@vue/eslint-config-typescript'
// configureVueProject({ scriptLangs: ['ts', 'tsx'] })
// More info at https://github.com/vuejs/eslint-config-typescript/#advanced-setup

export default defineConfigWithVueTs(
  {
    name: 'app/files-to-lint',
    files: ['**/*.{ts,mts,tsx,vue}'],
  },

  {
    name: 'app/files-to-ignore',
    ignores: ['**/dist/**', '**/dist-ssr/**', '**/coverage/**'],
  },

  {
    rules: {
      'prettier/prettier': [
        'warn',
        {
          singleQuote: true, // 单引号
          semi: false, // 不使用分号
          printWidth: 80, // 单行字符数 80
          trailingComma: 'none', // 结尾不加逗号
          endOfLine: 'auto', // 自动识别换行符 换行符合不限制 （win mac 不一致）
        },
      ],
      'vue/multi-word-component-names': [
        'warn',
        {
          ignores: ['index'], // 忽略 index 组件, vue 组件名称多单词组成
        },
      ],
      'vue/no-setup-props-destructure': ['off'], // 关闭 setup 语法糖中 props 解构的警告
    },
  },

  pluginVue.configs['flat/essential'],
  vueTsConfigs.recommended,
  skipFormatting,
)
