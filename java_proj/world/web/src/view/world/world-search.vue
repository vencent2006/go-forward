<template>
    <div class="search-container">
        <div class="search-header">
            <a-input-search
                v-model:value="searchKey"
                placeholder="搜索世界"
                @search="handleSearch"
                size="large"
                class="search-input"
            />
        </div>

        <div class="search-result" v-if="hasSearched">
            <a-list
                :data-source="searchResults"
                :loading="loading"
                item-layout="horizontal"
            >
                <template #renderItem="{ item }">
                    <a-list-item>
                        <a-card hoverable class="world-card">
                            <template #cover>
                                <img :src="item.cover || 'default-cover.jpg'" alt="world cover"/>
                            </template>
                            <a-card-meta :title="item.name">
                                <template #description>
                                    <div class="world-info">
                                        <p class="description">{{ item.description }}</p>
                                        <div class="world-stats">
                                            <span><UserOutlined /> {{ item.memberCount }}</span>
                                            <span><ClockCircleOutlined /> {{ item.createTime }}</span>
                                        </div>
                                    </div>
                                </template>
                            </a-card-meta>
                            <template #actions>
                                <EnterOutlined key="enter" @click="handleEnter(item)" />
                            </template>
                        </a-card>
                    </a-list-item>
                </template>
            </a-list>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { message } from 'ant-design-vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

const router = useRouter();
const loading = ref(false);
const searchKey = ref('');
const searchResults = ref([]);
const hasSearched = ref(false);

const handleSearch = async () => {
    if (!searchKey.value.trim()) {
        message.warning('请输入搜索关键词');
        return;
    }

    loading.value = true;
    try {
        const response = await axios.get('/myworld/api/world/search', {
            params: { keyword: searchKey.value }
        });
        const data = response.data;
        if (data.success) {
            searchResults.value = data.data;
            hasSearched.value = true;
        } else {
            message.error(data.message);
        }
    } catch (error) {
        message.error('搜索失败');
    } finally {
        loading.value = false;
    }
};

const handleEnter = (world) => {
    router.push(`/world/${world.id}`);
};
</script>

<style scoped>
.search-container {
    padding: 24px;
}

.search-header {
    margin-bottom: 24px;
}

.search-input {
    max-width: 600px;
    width: 100%;
}

.world-card {
    width: 100%;
}

.description {
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
}

.world-stats {
    display: flex;
    gap: 16px;
    color: rgba(0, 0, 0, 0.45);
}

:deep(.ant-card-cover img) {
    height: 200px;
    object-fit: cover;
}

@media (max-width: 768px) {
    .search-container {
        padding: 16px;
    }

    :deep(.ant-card-cover img) {
        height: 160px;
    }
}
</style>