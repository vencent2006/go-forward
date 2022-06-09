<template>
  <div class="dashboard-editor-container">

    <panel-group :data="panelGroupData" />

    <el-row :gutter="32">
      <el-col :xs="24" :sm="24" :lg="16">
        <div class="chart-wrapper">
          <line-chart :chart-data="lineChartData" />
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <pie-chart />
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import PanelGroup from './components/PanelGroup'
import LineChart from './components/LineChart'
import PieChart from './components/PieChart'
import { panelGroupData, flowStat } from '@/api/dashboard'

export default {
  name: 'DashboardAdmin',
  components: {
    PanelGroup,
    LineChart,
    PieChart
  },
  data() {
    return {
      panelGroupData: {
        'serviceNum': 23,
        'todayRequestNum': 1200,
        'currentQps': 200,
        'appNum': 5
      },
      lineChartData: {
        'title': '今日流量统计',
        'today': [220, 182, 191, 134, 150, 120, 110, 125, 145, 122, 165, 122, 120, 110, 125, 145, 122, 165, 122, 220, 182, 191, 134, 150],
        'yesterday': [120, 110, 125, 145, 122, 165, 122, 220, 182, 191, 134, 150, 220, 182, 191, 134, 150, 120, 110, 125, 145, 122, 165, 122]
      }
    }
  },
  created() {
    this.fetchPanelGroupData()
    this.fetchFlowStat()
  },
  methods: {
    fetchPanelGroupData() {
      panelGroupData({}).then(response => {
        this.panelGroupData = response.data
      }).catch(() => {

      })
    },
    fetchFlowStat() {
      flowStat({}).then(response => {
        this.lineChartData.today = response.data.today
        this.lineChartData.yesterday = response.data.yesterday
      }).catch(() => {

      })
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 32px;
  background-color: rgb(240, 242, 245);
  position: relative;

.github-corner {
  position: absolute;
  top: 0px;
  border: 0;
  right: 0;
}

.chart-wrapper {
  background: #fff;
  padding: 16px 16px 0;
  margin-bottom: 32px;
}
}

@media (max-width:1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}
</style>

