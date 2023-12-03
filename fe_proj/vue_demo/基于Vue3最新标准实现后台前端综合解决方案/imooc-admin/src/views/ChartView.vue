<template>
  <!-- 1. 创建 DOM 容器，指定大小 -->
  <div class="trend-chart-container" ref="target"></div>
</template>

<script setup>
// 2. 导入 Echarts 模块
import * as echarts from 'echarts' // echarts是按需导入，所以这里是导入*
// import { onMounted, ref, defineProps } from 'vue'
import { onMounted, ref } from 'vue'

// const props = defineProps({
//   data: {
//     type: Object,
//     required: true
//   }
// })

// console.log(props.data)
const dailyCurve = [
  { time: '2022-03', amount: '141.95' },
  { time: '2022-04', amount: '54.71' },
  { time: '2022-05', amount: '97.28' },
  { time: '2022-06', amount: '143.00' },
  { time: '2022-07', amount: '89.38' },
  { time: '2022-08', amount: '101.45' }
]
console.log(dailyCurve)
const monthAmountList = [
  { time: '2022-03', amount: '122.84' },
  { time: '2022-04', amount: '80.41' },
  { time: '2022-05', amount: '82.42' },
  { time: '2022-06', amount: '122.64' },
  { time: '2022-07', amount: '72.48' },
  { time: '2022-08', amount: '138.92' }
]
console.log(monthAmountList)

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
    // 鼠标移入之后展示提示框
    tooltip: {
      // 鼠标移入到坐标轴，触发提示
      trigger: 'axis',
      // 移入坐标轴时，提示框展示内容的配置
      axisPointer: {
        // 显示十字准心
        type: 'cross',
        // 指示器的样式
        crossStyle: {
          // 色值
          color: '#999'
        }
      }
    },
    // 图例配置
    legend: {
      // 图例数据
      data: ['月累计收益', '日收益曲线'],
      // 图例展示位置
      right: 0
    },
    // echarts 网格绘制的位置
    grid: {
      top: 20,
      right: 0,
      bottom: 0,
      left: 0,
      // 计算间距时，是否包含标签
      containLabel: true
    },
    // X 轴配置
    xAxis: {
      // 坐标轴类型
      type: 'category', // category 适用于 离散值
      // 坐标值的数据
      // data: props.data.monthAmountList.map((item) => item.time),
      data: monthAmountList.map((item) => item.time),
      // 刻度尺
      axisTick: {
        show: false // 关闭刻度尺
      }
    },
    // Y 轴配置
    yAxis: {
      // 坐标轴类型
      type: 'value', // y轴不能再用category？
      // 最小值
      min: 0,
      // 最大值
      max: function (value) {
        console.log('max: ' + value)
        return parseInt(value.max * 1.2)
      },
      // 刻度上展示的文字
      axisLabel: {
        formatter: '{value} 万元'
      }
    },
    // 图表配置
    series: [
      // 柱状图
      {
        // 图表类型
        type: 'bar',
        // 图表的名字，对应图例
        name: '月累计收益',
        // 柱的宽度
        barWidth: 20,
        // 提示框的展示配置
        tooltip: {
          valueFormatter: function (value) {
            return value + '万元'
          }
        },
        // 数据源
        // data: props.data.mouthAmountList.map((item) => item.amount)
        data: monthAmountList.map((item) => item.amount)
      },
      // 曲线图
      {
        // 图表类型
        type: 'line',
        // 颜色
        color: '#6EC6D0',
        // 平滑
        smooth: true,
        // 图表的名字，对应图例
        name: '日收益曲线',
        // 提示框的展示配置
        tooltip: {
          valueFormatter: function (value) {
            return value + '万元'
          }
        },
        // 数据源
        // data: props.data.dailyCurve.map((item) => item.amount)
        data: dailyCurve.map((item) => item.amount)
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
