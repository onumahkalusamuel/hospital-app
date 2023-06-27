<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '../../../components/Breadcrumbs.vue';
import PageHeader from '../../../components/PageHeader.vue';
import apiRequest from '../../../services/http/api-requests';
import dayjs from 'dayjs'
import TextField from '../../../components/form/TextField.vue'
import PrimaryButton from '../../../components/form/PrimaryButton.vue';
import SecondaryButton from '../../../components/form/SecondaryButton.vue';
import { onMounted, ref } from 'vue';
import { router } from '../../../router';
import { UserPlusIcon } from '@heroicons/vue/24/outline';
import { toasts } from '../../../stores/toasts';
import { Patient, PatientHistory } from '../../../interfaces';
import { useRoute } from 'vue-router';
import HistoryForm from './history/HistoryForm.vue';
// history parts
import Diagnosis from './history/Diagnosis.vue';
import Examination from './history/Examination.vue';
import General from './history/General.vue';
import TestResult from './history/TestResult.vue';
import Treatment from './history/Treatment.vue';

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Patients",  link: { name: "patients" } },
] as BreadcrumbItem[]);

const route = useRoute();
const createForm = ref(null);
const patient = ref({} as Patient);
const historyComponents = [Examination, Diagnosis, TestResult, Treatment, General];
const historyType = ref(historyComponents[0]);

const fetchPatient = async () => {
  patient.value = await apiRequest.get(`patients/${route.params.id}`);
}

const setHistoryType = (type: typeof HistoryForm) => {
  historyType.value = type;
}

onMounted(async() => {
  await fetchPatient()
  if(patient.value) {
    breadcrumbs.value.push(
      {
      title: `${patient.value.lastname} ${patient.value.firstname} - ${patient.value.card_no}`,
      link: {name: 'view-patient', params: {id: patient.value.id}}
      },
      {
        title: 'History',
        link: {name: 'patient-history', params: {id: patient.value.id }}
      },
      { title: "Add History", current: true },
    );
  }
})

const createHistory = async () => {
  const details = Object.fromEntries(new FormData(createForm.value as never as HTMLFormElement).entries())

  const formData = {
    details,
    date_time: dayjs().format('YYYY-MM-DDTHH:mm:ss[Z]'),
    type: historyType.value.__name,
  } as PatientHistory;
  
  const create = await apiRequest.post(`patients/${route.params.id}/patient-history`, formData);
  if(create.id) {
    toasts.addToast({message: "history added successfully.", type: 'success'});
    router.push({ name: 'patient-history' })
  }
}
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="Add Patient History" :icon-src="UserPlusIcon"></PageHeader>
    <div style="padding: 0 15px; display: flex; justify-content: space-between; border-top:1px solid #333">
    </div>
    <hr class="ml-4 mr-4" />
    <div class="page-scroll-area">
      <form method="POST" v-on:submit.prevent="createHistory" class="form flex flex-wrap gap-3 pl-4" ref="createForm">
        <span>History Type</span>
          <div class="w-full flex flex-wrap mb-3 gap-3">
            <div 
              v-for="(type, i) in historyComponents"
              :key="i"
              class="cursor-pointer p-2 px-3 border-[1px] border-gray-400 rounded"
              :class="historyType.__name==type.__name?'bg-blue-600 text-white':''"
              @click="() => setHistoryType(type)"
            > {{ type.__name }} </div>
          </div>
          <component :is="historyType" />  
          <div class="min-w-[510px]">
            <TextField label="Note" placeholder="note" name="note"></TextField>
          </div>
          <div class="w-full"></div>
          <div class="min-w-[250px] flex space-x-3">
            <PrimaryButton type="submit" class="w-full">Submit</PrimaryButton>
            <SecondaryButton type="button" class="w-full" v-on:click.prevent="()=>$router.push({name: 'patient-history'})">Cancel</SecondaryButton>
          </div>
      </form>
    </div>
  </div>
</template>