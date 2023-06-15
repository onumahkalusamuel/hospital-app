<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '../../../components/Breadcrumbs.vue';
import ActionButton from '../../../components/ActionButton.vue';
import PageHeader from '../../../components/PageHeader.vue';
import apiRequest from '../../../services/http/api-requests';
import TextField from '../../../components/form/TextField.vue';
import { Delivery } from '../../../interfaces';
import { MagnifyingGlassIcon } from '@heroicons/vue/24/outline';
import { UserGroupIcon } from '@heroicons/vue/24/solid';
import { onMounted, ref } from 'vue';

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Deliveries", current: true },
] as BreadcrumbItem[]);

const deliveries = ref([] as Delivery[]);

const deletItem = async (id: string) => {
  const del = await apiRequest.deleteRecord(`deliveries/${id}`);
  if(del && del.message) alert(del.message);
  fetch();
}
const fetch = async() => {
  deliveries.value = await apiRequest.get('deliveries');  
}

onMounted(async() =>{fetch()});
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="Deliveries" subtitle="Manage deliveries" :icon-src="UserGroupIcon" />
    <div style="padding: 0 15px; display: flex; justify-content: space-between; border-top:1px solid #333">
      <div>
        <ActionButton v-on:click="() => $router.push({name: 'add-delivery'})" icon-src="/icons/plus.svg">Add delivery</ActionButton>
      </div>
      <div>
        <TextField name="" placeholder="Search">
          <template #prepend>
            <MagnifyingGlassIcon class="h-5 w-5" />
          </template>
        </TextField>
      </div>
    </div>
    <hr class="ml-4 mr-4" />
    <div class="page-scroll-area">
      <div v-if="!deliveries.length">
        <div class="text-center mt-4 mb-4 pt-4">No records found</div>
        <hr/>
      </div>
      <div v-else class="table">
        <div class="table-row table-header">
          <div class="cell-data cell-sn">S/N</div>
          <div class="cell-data cell-size-1">Patient Name</div>
          <div class="cell-data">Delivery Mode</div>
          <div class="cell-data">Sex</div>
          <div class="cell-data">Baby Weight</div>
          <div class="cell-data cell-size-1">Condition.</div>
          <div class="cell-data cell-size-1">Delivery Time</div>
          <div class="cell-data">Actions</div>
        </div>
        <div class="table-body">
          <div class="table-row" v-for="delivery, i in deliveries" :key="delivery.id">
            <div class="cell-data cell-sn">{{ i + 1 }}</div>
            <div class="cell-data cell-size-1">
              <router-link class="text-blue-600 hover:underline" :to="{name: 'view-delivery', params: { id: delivery.id }}">
                {{ `${delivery.patient?.firstname} ${delivery.patient?.lastname}` }}
              </router-link>
            </div>
            <div class="cell-data">{{ delivery.delivery_mode }}</div>
            <div class="cell-data">{{ delivery.baby_sex }}</div>
            <div class="cell-data">{{ delivery.baby_weight /1000 }} {{ delivery.baby_weight_unit }}</div>
            <div class="cell-data cell-size-1">{{ delivery.condition }}</div>
            <div class="cell-data cell-size-1">{{ delivery.delivery_date_time }}</div>
            <div class="cell-data">
              <ActionButton v-on:click="deletItem(delivery.id)" icon-src="/icons/trash.svg">Delete</ActionButton>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>