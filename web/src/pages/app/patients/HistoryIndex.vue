<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import ActionButton from '@/components/ActionButton.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import { Pagination, Patient, PatientHistory } from '@/interfaces'
import { UsersIcon, TrashIcon } from '@heroicons/vue/24/solid';
import { onMounted, ref, watch } from 'vue';
import { popupStore, toasts, user } from '@/stores';
import { useRoute } from 'vue-router';
import dayjs from 'dayjs';
import Paging from '@/components/Paging.vue';
import { useDebounce } from '@/utils/debounce';
import DownloadPatientHistoryPopup from './popups/DownloadPatientHistoryPopup.vue';


const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Patients", link: { name: "patients" } },
] as BreadcrumbItem[]);

const downloadHistoryPopupId = ref('downloadhistory');

const pagination = ref({
  limit: 10,
  page: 1,
  rows: [] as PatientHistory[],
  sort_key: '',
  sort_order: 'desc',
  total_pages: 0,
  total_rows: 0,
  query: ''
} as Pagination);


const route = useRoute();
const debounce = useDebounce();
const patient = ref({} as Patient);
const showDownloadHistoryPopup = () => { popupStore.id = downloadHistoryPopupId.value; popupStore.show = true;}

const deletItem = async (id: string)  => {
  const del = await apiRequest.deleteRecord(`patients/${patient.value.id}/patient-history/${id}`);
  if(del && del.message) toasts.addToast({message: del.message, type: 'success'});
  await fetchPatientHistory();
}

const fetchPatientHistory = async () => {
  debounce(async () => {
  const filter = {...pagination.value, rows: null };
  pagination.value = await apiRequest.get(`patients/${route.params.id}/patient-history?${new URLSearchParams(filter as never as Record<string, string>)}`);
  })
}

const fetchPatient = async () => {
  // pHistory.value = await apiRequest.get(`patients/${route.params.id}/patient-history`);
  patient.value = await apiRequest.get(`patients/${route.params.id}`);
  await fetchPatientHistory();
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

watch(() => pagination.value.limit, fetchPatientHistory);
watch(() => pagination.value.page, fetchPatientHistory);
watch(() => pagination.value.query, fetchPatientHistory);
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader
      title="Patient History"
      :subtitle="`${patient.lastname} ${patient.firstname} - ${patient.card_no}`" :icon-src="UsersIcon">
    </PageHeader>
    <div class="px-[15px] flex justify-between border-t-[1px] border-[#333] py-2">
      <div class="flex gap-2">
        <ActionButton dark v-on:click="() => $router.push({name: 'patient-history-add'})" :icon-src="TrashIcon">Add history</ActionButton>
        <ActionButton v-on:click="showDownloadHistoryPopup" :icon-src="TrashIcon">Download Patient History</ActionButton>
      </div>
    </div>

    <hr class="ml-4 mr-4" />
    <div class="page-scroll-area">
      <div v-if="!pagination.rows.length">
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
          <div class="table-row" v-for="history, i in pagination.rows" :key="patient.id">
            <div class="cell-data cell-sn">{{ i + 1 }}</div>
            <div class="cell-data cell-size-1">{{ dayjs(history.date_time).format('DD-MM-YYYY hh:mm A') }}</div>
            <div class="cell-data cell-size-1">{{ history.type }}</div>
            <div class="cell-data cell-size-1">
              <div>
                <div>{{ history.details.note }}</div>
                <div v-if="history.type=='Admission'"> Ward: {{ history.details.ward_number }}; Room: {{ history.details.room_number }}</div>
                <div v-if="history.type=='Appointment'"> Appt: {{ history.details.appointment_type }}</div>
              </div>
            </div>
            <div class="cell-data">
              <ActionButton v-if="user.role == '1'" v-on:click="deletItem(patient.id)" :icon-src="TrashIcon">Delete</ActionButton>
            </div>
          </div>
        </div>
      </div>
    </div>
    <Paging v-model="pagination" />
  </div>
  <DownloadPatientHistoryPopup :popup-id="downloadHistoryPopupId" />
</template>

<style scoped>
</style>