<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '../../../components/Breadcrumbs.vue';
import ActionButton from '../../../components/ActionButton.vue';
import PageHeader from '../../../components/PageHeader.vue';
import apiRequest from '../../../services/http/api-requests';
import { Patient, PatientHistory } from '../../../interfaces'
import { UsersIcon, TrashIcon } from '@heroicons/vue/24/solid';
import { onMounted, ref } from 'vue';
import { toasts, user } from '../../../stores';
import { useRoute } from 'vue-router';


const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Patients", link: { name: "patients" } },
] as BreadcrumbItem[]);

const route = useRoute();
const pHistory = ref([] as PatientHistory[]);
const patient = ref({} as Patient);

const deletItem = async (id: string)  => {
  const del = await apiRequest.deleteRecord(`patients/${patient.value.id}/patient-history/${id}`);
  if(del && del.message) toasts.addToast({message: del.message, type: 'success'});
  fetchPatient();
}

const fetchPatient = async () => {
  patient.value = await apiRequest.get(`patients/${route.params.id}/patient-history`);
  pHistory.value = patient.value.patient_history as PatientHistory[];
}

onMounted(async() => {
  await fetchPatient()
  if(patient.value) {
    breadcrumbs.value.push({
      title: `${patient.value.lastname} ${patient.value.firstname} - ${patient.value.card_no}`,
      link: {name: 'view-patient', params: {id: patient.value.id}}
      }, { title: 'History', current: true }
    );
  }
})
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader
      title="Patient History"
      :subtitle="`${patient.lastname} ${patient.firstname} - ${patient.card_no}`" :icon-src="UsersIcon">
    </PageHeader>
    <div style="padding: 0 15px; display: flex; justify-content: space-between; border-top:1px solid #333">
      <div>
        <ActionButton v-on:click="() => $router.push({name: 'patient-history-add'})" :icon-src="TrashIcon">Add history</ActionButton>
      </div>
    </div>

    <hr class="ml-4 mr-4" />
    <div class="page-scroll-area">
      <div v-if="!pHistory.length">
        <div class="text-center mt-4 mb-4 pt-4">No records found</div>
        <hr/>
      </div>
      <div v-else class="table">
        <div class="table-row table-header">
          <div class="cell-data cell-sn">S/N</div>
          <div class="cell-data cell-size-1">Date</div>
          <div class="cell-data cell-size-1">Type</div>
          <div class="cell-data cell-size-1">Note</div>
          
          <div class="cell-data">Actions</div>
        </div>
        <div class="table-body">
          <div class="table-row" v-for="history, i in pHistory" :key="patient.id">
            <div class="cell-data cell-sn">{{ i + 1 }}</div>
            <div class="cell-data cell-size-1">{{ history.date_time }}</div>
            <div class="cell-data cell-size-1">{{ history.type }}</div>
            <div class="cell-data cell-size-1">{{ history.details.note }}</div>
            <div class="cell-data">
              <ActionButton v-if="user.role == '1'" v-on:click="deletItem(patient.id)" :icon-src="TrashIcon">Delete</ActionButton>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>