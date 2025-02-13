<template>
    <div class="world-list-container">
        <!-- 搜索区域 -->
        <div class="search-section">
            <a-input-search
                v-model:value="params.name"
                placeholder="搜索世界"
                @search="handleSearch"
                style="width: 300px"
            />
            <a-button type="primary" @click="handleCreate">
                <PlusOutlined /> 创建
            </a-button>
        </div>

        <!-- 列表区域 -->
        <a-list
            :data-source="worlds"
            :loading="loading"
            item-layout="horizontal"
            class="world-list"
        >
            <template #renderItem="{ item }">
                <a-list-item>
                    <a-card hoverable style="width: 100%">
                        <template #cover>
                            <img :src="item.image || 'default-cover.jpg'" alt="world cover"/>
                        </template>
                        <a-card-meta :title="item.name">
                            <template #description>
                                <div class="world-info">
                                    <p>{{ item.desc }}</p>
                                    <div class="world-stats">
                                        <span><UserOutlined /> {{ item.memberCount }} 成员</span>
                                        <span><ClockCircleOutlined /> {{ formatDate(item.beginAt) }}</span>
                                    </div>
                                </div>
                            </template>
                        </a-card-meta>
                        <template #actions>
                            <EnterOutlined key="enter" @click="handleEnter(item)" />
                            <EditOutlined key="edit" @click="handleEdit(item)" />
                            <DeleteOutlined key="delete" @click="handleDelete(item)" />
                        </template>
                    </a-card>
                </a-list-item>
            </template>
        </a-list>

        <!-- 分页 -->
        <div class="pagination">
            <a-pagination
                v-model:current="params.page"
                :total="total"
                :pageSize="params.size"
                @change="handlePageChange"
            />
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { message } from 'ant-design-vue';
import axios from 'axios';
import { useRouter } from 'vue-router';
import dayjs from 'dayjs';

const router = useRouter();
const loading = ref(false);
const worlds = ref([]);
const total = ref(0);

const params = ref({
    page: 1,
    size: 8,
    name: ''
});

// 获取世界列表
const fetchWorlds = async () => {
    loading.value = true;
    try {
        const response = await axios.get('/myworld/api/world', {
            params: params.value
        });
        const data = response.data;
        console.log(data);
        if (data.success) {
            worlds.value = data.data.list;
            total.value = data.data.total;
        } else {
            message.error(data.message);
        }
    } catch (error) {
        message.error('获取世界列表失败');
    } finally {
        loading.value = false;
    }
};

// 搜索
const handleSearch = () => {
    params.value.page = 1;
    fetchWorlds();
};

// 分页
const handlePageChange = (page) => {
    params.value.page = page;
    fetchWorlds();
};

// 进入世界
const handleEnter = (world) => {
    router.push(`/world/${world.id}`);
};

// 编辑世界
const handleEdit = (world) => {
    router.push(`/world/edit/${world.id}`);
};

// 删除世界
const handleDelete = async (world) => {
    try {
        const response = await axios.delete(`/nls/web/world/delete/${world.id}`);
        if (response.data.success) {
            message.success('删除成功');
            fetchWorlds();
        } else {
            message.error(response.data.message);
        }
    } catch (error) {
        message.error('删除失败');
    }
};

// 创建新世界
const handleCreate = () => {
    router.push('/world/create');
};

const formatDate = (date) => {
    return dayjs(date).format('YYYY-MM-DD HH:mm:ss');
};

onMounted(() => {
    fetchWorlds();
});
</script>

<style scoped>
.world-list-container {
    padding: 24px;
}

.search-section {
    display: flex;
    justify-content: space-between;
    margin-bottom: 24px;
}

.world-list {
    margin-bottom: 24px;
}

.world-info {
    display: flex;
    flex-direction: column;
    gap: 8px;
}

.world-stats {
    display: flex;
    gap: 16px;
    color: rgba(0, 0, 0, 0.45);
}

.pagination {
    display: flex;
    justify-content: center;
    margin-top: 24px;
}

:deep(.ant-card-cover img) {
    height: 200px;
    object-fit: cover;
}
</style>