<script setup lang="ts">
import { FunctionalComponent, HTMLAttributes, VNodeProps } from 'vue';

defineProps<{
  iconSrc?: FunctionalComponent<HTMLAttributes & VNodeProps>,
  title: string,
  subtitle?: string
}>()
</script>
<template>
  <div class="page-header-container">
    <div class="w-[80px] h-[80px] mr-2" v-if="iconSrc">
      <component :is="iconSrc" class="w-[80px] h-[80px] text-[#0078d4]" />
    </div>
    <div class="flex flex-col justify-center">
      <div class=" text-2xl text-[#333]">{{ title }}</div>
      <div v-if="subtitle" class="page-header-subtitle">{{ subtitle }}</div>
    </div>
  </div>
</template>
<style scoped>
.page-header-container {
  display:flex;
  padding: 10px 15px;
}

.page-header-subtitle {
  font-size: 1rem;
  color: #555;
  margin-top: 8px;
}
</style>