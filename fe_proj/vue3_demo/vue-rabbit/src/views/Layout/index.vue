<script setup>
import LayoutFixed from "@/views/Layout/components/LayoutFixed.vue";
import LayoutNav from "@/views/Layout/components/LayoutNav.vue";
import LayoutHeader from "@/views/Layout/components/LayoutHeader.vue";
import LayoutFooter from "@/views/Layout/components/LayoutFooter.vue";
import {useCategoryStore} from "@/stores/category.js";
import {onMounted} from "vue";

const categoryStore = useCategoryStore()
onMounted(() => {
  categoryStore.getCategory()
})
</script>

<template>
  <LayoutFixed/>
  <LayoutNav/>
  <LayoutHeader/>
  <RouterView/>
  <LayoutFooter/>
</template>
