<template>
    <div ref="chartRef" class="chart-container"></div>
</template>

<script setup>
import {onMounted, onUnmounted, ref} from 'vue'
import * as echarts from 'echarts'

// 定义图表容器 ref
const chartRef = ref(null)

// 存储 ECharts 实例
let chartInstance = null

// 图表配置项
const initChart = () => {
    const option = {
        title: {
            text: '分sku销量',
            subtext: '纯属虚构',
            left: 'center'
        },
        tooltip: {},
        xAxis: {
            data: ['衬衫', '羊毛衫', '雪纺衫', '裤子', '高跟鞋', '袜子']
        },
        yAxis: {},
        series: [
            {
                name: '销量',
                type: 'bar',
                data: [5, 20, 36, 10, 10, 20]
            }
        ]
    }

    chartInstance.setOption(option)
}

// 初始化图表
onMounted(() => {
    if (chartRef.value) {
        chartInstance = echarts.init(chartRef.value)
        initChart()

        // 响应窗口大小变化
        window.addEventListener('resize', () => {
            chartInstance.resize()
        })
    }
})

// 组件卸载时销毁实例
onUnmounted(() => {
    chartInstance.dispose()
    window.removeEventListener('resize', () => {
        chartInstance.resize()
    })
})
</script>

<style scoped>
.chart-container {
    width: 600px;
    height: 400px;
}
</style>