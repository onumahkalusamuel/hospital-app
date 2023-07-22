<script setup lang="ts">
import { Pagination } from '@/interfaces';
import { watch } from 'vue';

const props = defineProps<{modelValue: Pagination}>()

const emit = defineEmits(['update:modelValue'])

const updateValue = () => {
    emit('update:modelValue', props.modelValue)
}

watch(() => props.modelValue, updateValue)

</script>
<template>
  <div class="px-[15px] flex justify-between border-t-[1px] border-gray-400">
    <div class="">
      <div class="py-2 text-center">Pagination</div>
      <div class="flex">
        <button :disabled="modelValue.page == 1" class="border-blue-600 p-2 py-1 border-[1px] rounded mx-1 hover:bg-blue-600 hover:text-white" @click="() => (modelValue.page = 1)">&lt;&lt;</button>
        <button :disabled="modelValue.page == 1" class="border-blue-600 p-2 py-1 border-[1px] rounded mx-1 hover:bg-blue-600 hover:text-white" @click="() => (modelValue.page = modelValue.page > 1 ? modelValue.page - 1: 1)">&lt;</button>
        <div class="border-blue-600 p-2 py-1 border-[1px] rounded mx-1 hover:bg-blue-600 hover:text-white bg-blue-600 text-white">{{ modelValue.page }}</div>
        <button :disabled="modelValue.page == modelValue.total_pages" class="border-blue-600 p-2 py-1 border-[1px] rounded mx-1 hover:bg-blue-600 hover:text-white" @click="() => (modelValue.page = modelValue.total_pages > modelValue.page ? modelValue.page + 1: modelValue.total_pages)">&gt;</button>
        <button :disabled="modelValue.page == modelValue.total_pages" class="border-blue-600 p-2 py-1 border-[1px] rounded mx-1 hover:bg-blue-600 hover:text-white" @click="() => (modelValue.page = modelValue.total_pages)">&gt;&gt;</button>
      </div>
    </div>
    <div class="text-center">
      <div class="py-2">Records Per Page</div>
      <div class="flex">
        <div
          v-for="num of [10,20,50,100]"
          class="border-blue-600 p-2 py-1 border-[1px] rounded mx-1 hover:bg-blue-600 hover:text-white"
          :class="num==modelValue.limit? 'bg-blue-600 text-white':'cursor-pointer'"
          @click="() => (modelValue.limit = num)"
        >{{ num }}</div>
      </div>
      
    </div>
  </div>
</template>