<script setup lang="ts">
import { FunctionalComponent, HTMLAttributes, VNodeProps } from 'vue';

defineProps<{
  iconSrc?: FunctionalComponent<HTMLAttributes & VNodeProps>;
  disabled?: boolean;
  outline?: boolean;
  classes?: string;
  iconClass?: string;
  dark?: boolean;
}>()

</script>
<template>
  <div style="display:inline-block">  
    <div class='rounded cursor-pointer flex gap-2 items-center justify-center h-[36px] px-2' :class=" dark ? 'hover:bg-blue-800 bg-[#0078d4] text-white': 'hover:bg-blue-100 text-[#333]'">
      <div class="text-[#0078d4] flex items-center" v-if="iconSrc">
        <component :is="iconSrc" :class=" dark ? 'text-white': ''" class="h-4 w-4" />
      </div>
      <div class="whitespace-nowrap">
        <slot class="whitespace-nowrap"></slot>
      </div>
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
</style>
