// 1. 导入所有的svg图标
// 2. 完成SvgIcon的全局注册
import SvgIcon from '@/components/SvgIcon'

// https://webpack.docschina.org/guides/dependency-management/#requirecontext
// require.context的参数
// 1.要搜索的目录
// 2.是否还有子目录
// 3.匹配的文件
const svgRequire = require.context('./svg', false, /\.svg$/)
// 此时返回了一个Require函数，可以接收一个request的参数，用于require的导入
// 该函数提供了三个属性，可以通过 svgRequire.keys() 获得所有的svg图标
// 遍历图标，把图标作为 request 参数导入 svgRequire 导入函数中，完成本地svg图标导入
// console.log(svgRequire.keys())
svgRequire.keys().forEach((svgIcon) => svgRequire(svgIcon))

export default (app) => {
  app.component('svg-icon', SvgIcon)
}
