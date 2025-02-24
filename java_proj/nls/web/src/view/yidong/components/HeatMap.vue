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
const getVirtualData = (year) => {
    const date = +echarts.time.parse(year + '-01-01');
    const end = +echarts.time.parse(+year + 1 + '-01-01');
    const dayTime = 3600 * 24 * 1000;
    const data = [];
    for (let time = date; time < end; time += dayTime) {
        data.push([
            echarts.time.format(time, '{yyyy}-{MM}-{dd}', false),
            Math.floor(Math.random() * 10000)
        ]);
    }
    return data;
}


// 图表配置项
const initChart = () => {
    const year = '2025'
    const option = {
        title: {
            top: 30,
            left: 'center',
            text: year + ' 每日异动热力图'
        },
        tooltip: {},
        visualMap: {
            min: 0,
            max: 10000,
            type: 'piecewise',
            orient: 'horizontal',
            left: 'center',
            top: 65
        },
        calendar: {
            top: 120,
            left: 30,
            right: 30,
            cellSize: ['auto', 13],
            range: year,
            itemStyle: {
                borderWidth: 0.5
            },
            yearLabel: {show: false}
        },
        series: {
            type: 'heatmap',
            coordinateSystem: 'calendar',
            data: getVirtualData(year)
        }
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