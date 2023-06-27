<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '../../../components/Breadcrumbs.vue';
import ActionButton from '../../../components/ActionButton.vue';
import PageHeader from '../../../components/PageHeader.vue';
import apiRequest from '../../../services/http/api-requests';
import { Staff, Roles } from '../../../interfaces'
import TextField from '../../../components/form/TextField.vue';
import { UsersIcon, MagnifyingGlassIcon, TrashIcon } from '@heroicons/vue/24/solid';
import { onMounted, ref } from 'vue';
import dayjs from 'dayjs';
import { toasts } from '../../../stores/toasts';

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Staff", current: true },
] as BreadcrumbItem[]);

const staff = ref([] as Staff[]);

const deletItem = async (id: string)  => {
  const del = await apiRequest.deleteRecord(`staff/${id}`);
  if(del && del.message) toasts.addToast({message: del.message, type: 'success'});
  fetchStaff();
}

const fetchStaff =async () => {
  staff.value = await apiRequest.get('staff');  
}

onMounted(async() => {fetchStaff()})
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="Staff" subtitle="Manage staff" :icon-src="UsersIcon">
    </PageHeader>
    <div style="padding: 0 15px; display: flex; justify-content: space-between; border-top:1px solid #333">
      <div>
        <ActionButton v-on:click="() => $router.push({name: 'add-staff'})" :icon-src="TrashIcon">Add staff</ActionButton>
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
      <div v-if="!staff.length">
        <div class="text-center mt-4 mb-4 pt-4">No records found</div>
        <hr/>
      </div>
      <div v-else class="table">
        <div class="table-row table-header">
          <div class="cell-data cell-sn">S/N</div>
          <div class="cell-data cell-size-1">Name</div>
          <div class="cell-data">Role</div>
          <div class="cell-data">Username</div>
          <div class="cell-data cell-size-1">Created At</div>
          <div class="cell-data">Actions</div>
        </div>
        <div class="table-body">
          <div class="table-row" v-for="st, i in staff" :key="st.id">
            <div class="cell-data cell-sn">{{ i + 1 }}</div>
            <div class="cell-data cell-size-1">
              <router-link class="text-blue-600 hover:underline" :to="{name: 'view-staff', params: { id: st.id }}">
                {{ `${st.firstname} ${st.lastname}` }}
              </router-link>
            </div>
            <div class="cell-data">{{ Roles[st.role] }}</div>
            <div class="cell-data">{{ st.username }}</div>
            <div class="cell-data cell-size-1">{{ dayjs(st.created_at).format('DD-MM-YYYY hh:mm A') }}</div>
            <div class="cell-data">
              <ActionButton v-if="st.role > 1" v-on:click="deletItem(st.id)" :icon-src="TrashIcon">Delete</ActionButton>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>