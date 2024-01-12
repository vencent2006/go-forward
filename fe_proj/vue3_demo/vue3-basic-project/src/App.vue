<script setup>
import Edit from './components/Edit.vue'
import {onMounted, ref} from "vue";
import axios from "axios";

// TODO: 列表渲染
const list = ref([])
const getList = async () => {
  const res = await axios.get('/list')
  list.value = res.data
}
onMounted(()=>getList())


// TODO: 删除功能
const onDelete = async (id) => {
  console.log(id)
  await axios.delete(`/del/${id}`)
  await getList()
}

// TODO: 编辑功能
// 思路: 打开弹框 -> 回填数据 -> 更新数据
// 1. 打开弹框（获取子组件实例 调用方法或者修改属性）
const editRef = ref(null)
const onEdit = (row) => {
  editRef.value.open(row)
}
// 2. 回填数据 （调用详情接口 / 当前行的静态数据）

</script>

<template>
  <div class="app">
    <el-table :data="list">
      <el-table-column label="ID" prop="id"></el-table-column>
      <el-table-column label="姓名" prop="name" width="150"></el-table-column>
      <el-table-column label="籍贯" prop="place"></el-table-column>
      <el-table-column label="操作" width="150">
        <template #default="{ row }">
          <el-button type="primary" @click="onEdit(row)" link>编辑</el-button>
          <el-button type="danger" @click="onDelete(row.id)" link>删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
  <Edit ref="editRef" @on-update="getList"/>
</template>

<style scoped>
.app {
  width: 980px;
  margin: 100px auto 0;
}
</style>
