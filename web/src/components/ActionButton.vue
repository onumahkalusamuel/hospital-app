<script setup lang="ts">
defineProps<{
  iconSrc?: string,
  disabled?: boolean,
  outline?: boolean
}>()
</script>
<template>
  <div style="display:inline-block">
    <div :class="(disabled?' disabled ':' ') + (outline ? ' outline ':' ')" class='action-button-container'>
      <div class="action-button-icon" v-if="iconSrc">
        <img :src="iconSrc" alt="" class="action-button-icon-img"/>
      </div>
      <div class="action-button-label"><slot></slot></div>
    </div>
  </div>
</template>
<style scoped>
.disabled {
  cursor: default !important;
  opacity: 0.5;
}
.outline {
  border: 1px solid #555;
  box-sizing: border-box;
}
.action-button-container {
  cursor:pointer;
  display: flex;
  align-items: center;
  height: 36px;
  padding: 0 4px;
}
.action-button-container:hover {
  background-color: azure;
}

.action-button-icon {
  margin: 0 4px;
  color: #0078d4;
  display: flex;
  align-items: center;
  filter: invert(38%) sepia(81%) saturate(7137%) hue-rotate(193deg) brightness(101%) contrast(102%);
}

.action-button-icon-img {
  width: 16px;
  height: 16px;
}

.action-button-label {
  color: #333;
  margin: 0 4px;
}
</style>
