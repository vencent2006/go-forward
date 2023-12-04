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
  const lineOption = {
    title: {
      text: '网站访问渠道统计'
    },
    tooltip: {
      trigger: 'axis'
    },
    legend: {
      data: []
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    toolbox: {
      feature: {
        saveAsImage: {}
      }
    },
    xAxis: {
      type: 'time', // type为time时,不要传xAxis.data的值,x轴坐标的数据会根据传入的时间自动展示
      boundaryGap: false // false横坐标两边不需要留白
    },
    yAxis: {
      type: 'value',
      name: '访问人数'
    },
    series: []
  }
  // 获取动态数据的函数
  getData()
  function getData() {
    const data = [
      {
        type: 'email',
        name: '邮件营销',
        data: [
          ['2020-10-1', 450],
          ['2020-10-2', 350],
          ['2020-10-3', 290],
          ['2020-10-4', 380],
          ['2020-10-5', 540],
          ['2020-10-6', null],
          ['2020-10-7', null],
          ['2020-10-8', 430],
          ['2020-10-9', 330],
          ['2020-10-10', 280],
          ['2020-10-11', 340],
          ['2020-10-12', 455],
          ['2020-10-13', 330]
        ]
      },
      {
        type: 'ad',
        name: '联盟广告',
        data: [
          ['2020-10-1', 50],
          ['2020-10-2', 150],
          ['2020-10-3', 100],
          ['2020-10-4', 140],
          ['2020-10-5', 141],
          ['2020-10-6', 66],
          ['2020-10-7', 78],
          ['2020-10-8', 67],
          ['2020-10-9', 55],
          ['2020-10-10', 80],
          ['2020-10-11', 40],
          ['2020-10-12', 120],
          ['2020-10-13', 130]
        ]
      },
      {
        type: 'video',
        name: '视频广告',
        data: [
          ['2020-10-1', 234],
          ['2020-10-2', 254],
          ['2020-10-3', 260],
          ['2020-10-4', 270],
          ['2020-10-5', 250],
          ['2020-10-6', 277],
          ['2020-10-7', 289],
          ['2020-10-8', 240],
          ['2020-10-9', 230],
          ['2020-10-10', 222],
          ['2020-10-11', 244],
          ['2020-10-12', 254],
          ['2020-10-13', 279]
        ]
      },
      {
        type: 'offline',
        name: '直接访问',
        data: [
          ['2020-10-1', null],
          ['2020-10-2', null],
          ['2020-10-3', null],
          ['2020-10-4', 40],
          ['2020-10-5', 80],
          ['2020-10-6', 87],
          ['2020-10-7', 98],
          ['2020-10-8', 33],
          ['2020-10-9', 35],
          ['2020-10-10', 78],
          ['2020-10-11', 48],
          ['2020-10-12', 67],
          ['2020-10-13', 30],
          ['2020-10-14 08:20:33', 77]
          // ['2021-2-13', 88]
        ]
      }
    ]
    const series = []
    const legendData = []
    data.forEach((item) => {
      const obj = {
        name: item.name,
        type: 'line',
        data: item.data
      }
      legendData.push(item.name)
      series.push(obj)
    })
    lineOption.legend.data = legendData
    lineOption.series = series
    // 第二个参数true表示清空之前的echarts设置,重新设置option,适用于legend等数据变化
    // 如果只是渲染数据变化可以不传,lenged等不会重新渲染
    mChart.setOption(lineOption, true)
  }
}
</script>

<style lang="scss" scoped>
.trend-chart-container {
  height: 100%;
}
</style>
