<template>
  <!-- 1. 创建 DOM 容器，指定大小 -->
  <div class="trend-chart-container" ref="target"></div>
</template>

<script setup>
// 2. 导入 Echarts 模块
import * as echarts from 'echarts' // echarts是按需导入，所以这里是导入*
import { onMounted, ref } from 'vue'

// 3. 利用 echarts.init(target), 获取到 mCharts实例
const target = ref(null)
let mChart
onMounted(() => {
  mChart = echarts.init(target.value)
  renderChart()
})

const renderChart = () => {
  // 4. 构建 options 配置对象 ( echarts 渲染的核心 )
  const options = {
    // 图表标题配置
    title: {
      // 标题文本
      text: 'Echarts 入门实例'
    },
    // 图例配置
    legend: {
      // 图例数据
      data: ['销量']
    },
    // X 轴配置
    xAxis: {
      // 数据
      data: ['衬衫', '羊毛衫', '雪纺衫', '裤子', '高跟鞋', '袜子']
    },
    // Y 轴配置
    yAxis: {},
    // 图表类型
    series: [
      // 数组中的每一个对象，都是一个图表
      {
        // 图表名称，与图例数据的名字要对应上
        name: '销量',
        // 图表的类型
        type: 'bar',
        // 图表的数据
        data: [5, 20, 36, 10, 10, 20]
      }
    ]
  }

  // 5. 利用 mChart.setOption 方法配置 options
  mChart.setOption(options)
}
</script>

<style lang="scss" scoped>
.trend-chart-container {
  height: 100%;
}
</style>
