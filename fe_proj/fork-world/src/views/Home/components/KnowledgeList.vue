<script setup lang="ts">
import { ref } from 'vue'
import KnowledgeCard from './KnowledgeCard.vue'
import type { KnowledgeList, KnowledgeParams, KnowledgeType } from '@/types/consult'
import { getKnowledgePage } from '@/services/consult'

const props = defineProps<{
  type: KnowledgeType
}>()

// 加载中状态
const loading = ref(false)
// 是否完整加载完毕数据
const finished = ref(false)
// 列表数据
const list = ref<KnowledgeList>([])
// 查询参数
const params = ref<KnowledgeParams>({
  type: props.type,
  current: 1,
  pageSize: 5,
})
// 滚动到底部 加载更多数据
const onLoad = async () => {
  const res = await getKnowledgePage(params.value)
  list.value.push(...res.data.rows)
  // 判断已经加载完成
  if (params.value.current >= res.data.total) {
    finished.value = true
  } else {
    params.value.current++
  }
  loading.value = false
}
</script>

<template>
  <div class="knowledge-list">
    <van-list
      v-model:loading="loading"
      :finished="finished"
      finished-text="没有更多了"
      @load="onLoad"
    >
      <knowledge-card v-for="(item, i) in list" :key="i" :item="item"></knowledge-card>
    </van-list>
  </div>
</template>

<style lang="scss" scoped>
.knowledge-list {
  padding: 0 15px;
}
</style>
