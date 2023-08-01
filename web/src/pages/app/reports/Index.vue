<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';

import PageHeader from '@/components/PageHeader.vue';
import { QueueListIcon } from '@heroicons/vue/24/solid';
import { onMounted, ref } from 'vue';

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Reports", current: true },
] as BreadcrumbItem[]);



onMounted(() => {})
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="Reports" subtitle="Generate reports" :icon-src="QueueListIcon">
    </PageHeader>
    <div class="px-[15px] flex justify-between border-t-[1px] border-[#333] py-2">
    </div>
    <div class="text-center p-5">Work in progress</div>
  </div>
</template>