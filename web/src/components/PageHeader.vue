<script setup lang="ts">
defineProps<{
  iconSrc?: string,
  title: string,
  subtitle?: string
}>()
</script>
<template>
  <div class="page-header-container">
    <div class="page-header-icon" v-if="iconSrc">
      <img :src="iconSrc" class="page-header-icon-image" alt=""/>
    </div>
    <div class="page-header-title-container">
      <div class="page-header-title">{{ title }}</div>
      <div v-if="subtitle" class="page-header-subtitle">{{ subtitle }}</div>
    </div>
  </div>
</template>
<style scoped>
.page-header-container {
  display:flex;
  padding: 10px 15px;
}

.page-header-icon-image{
  height: 80px;
  width: 80px;
  filter: invert(38%) sepia(81%) saturate(7137%) hue-rotate(193deg) brightness(101%) contrast(102%);
}
.page-header-icon {
  height: 80px;
  width: 80px;
  margin-right: 12px;
}
.page-header-title-container {
  display:flex;
  flex-direction: column;
  justify-content: center;
}

.page-header-title {
  /* font-weight: 600; */
  font-size: 2rem;
  color: #333;
}

.page-header-subtitle {
  font-size: 1rem;
  color: #555;
  margin-top: 8px;
}

</style>