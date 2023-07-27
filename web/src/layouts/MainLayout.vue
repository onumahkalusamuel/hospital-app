<script lang="ts" setup>
import { RouterView } from 'vue-router';
import { XCircleIcon } from '@heroicons/vue/24/outline'
const exitApp = () => window.close();
</script>;

<template>
  <div @click="exitApp" class="fixed right-[20px] top-[10px] cursor-pointer">
    <XCircleIcon class="w-8 h-8 hover:bg-stone-200" />
  </div>
  <div class="container">
    <div class="inner-container">
      <router-view></router-view>
    </div>
  </div>
</template>

<style>
.container {
  display:flex;
  align-items: center;
  flex-direction: column;
  justify-content: center;
  height: 100vh;
}
.inner-container {
  max-width: 400px;
  min-width: 350px;
  text-align: center;
  border: 1px solid #888888;
  padding: 3rem 2rem;
  box-sizing: border-box;
}
</style>
