<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '../../../components/Breadcrumbs.vue';
import ActionButton from '../../../components/ActionButton.vue';
import PageHeader from '../../../components/PageHeader.vue';
import apiRequest from '../../../services/http/api-requests';
import { Patient } from '../../../interfaces'
import TextField from '../../../components/form/TextField.vue';
import { UsersIcon, MagnifyingGlassIcon } from '@heroicons/vue/24/solid';
import { onMounted, ref } from 'vue';
import { toasts } from '../../../stores/toasts';

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Patients", current: true },
] as BreadcrumbItem[]);

const patients= ref([] as Patient[]);

const deletItem = async (id: string)  => {
  const del = await apiRequest.deleteRecord(`patients/${id}`);
  if(del && del.message) toasts.addToast({message: del.message, type: 'success'});
  fetchPatients();
}

const fetchPatients =async () => {
  patients.value = await apiRequest.get('patients');  
}

onMounted(async() => {fetchPatients()})
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="Patients" subtitle="Manage patients" :icon-src="UsersIcon">
    </PageHeader>
    <div style="padding: 0 15px; display: flex; justify-content: space-between; border-top:1px solid #333">
      <div>
        <ActionButton v-on:click="() => $router.push({name: 'add-patient'})" :icon-src="UsersIcon">Add patient</ActionButton>
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
      <div v-if="!patients.length">
        <div class="text-center mt-4 mb-4 pt-4">No records found</div>
        <hr/>
      </div>
      <div v-else class="table">
        <div class="table-row table-header">
          <div class="cell-data cell-sn">S/N</div>
          <div class="cell-data cell-size-1">Name</div>
          <div class="cell-data">Card No</div>
          <div class="cell-data">Sex</div>
          <div class="cell-data">Phone</div>
          <div class="cell-data cell-size-1">Current Appnt.</div>
          <div class="cell-data">Actions</div>
        </div>
        <div class="table-body">
          <div class="table-row" v-for="patient, i in patients" :key="patient.id">
            <div class="cell-data cell-sn">{{ i + 1 }}</div>
            <div class="cell-data cell-size-1">
              <router-link class="text-blue-600 hover:underline" :to="{name: 'view-patient', params: { id: patient.id }}">
                {{ `${patient.firstname} ${patient.lastname}` }}
              </router-link>
            </div>
            <div class="cell-data">{{ patient.card_no }}</div>
            <div class="cell-data">{{ patient.sex }}</div>
            <div class="cell-data">{{ patient.phone }}</div>
            <div class="cell-data cell-size-1">{{ patient.current_appointment }}</div>
            <div class="cell-data">
              <ActionButton v-on:click="deletItem(patient.id)" :icon-src="UsersIcon">Delete</ActionButton>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>