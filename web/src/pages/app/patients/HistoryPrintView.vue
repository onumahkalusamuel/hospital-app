<script lang="ts" setup>
import ActionButton from '@/components/ActionButton.vue';
import apiRequest from '@/services/http/api-requests';
import { Patient, PatientHistory } from '@/interfaces';
import { DocumentArrowDownIcon } from '@heroicons/vue/24/solid';
import { XCircleIcon } from '@heroicons/vue/24/outline';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import dayjs from 'dayjs';
import * as html2pdf from 'html2pdf.js';
import { hospital } from '@/stores';

const route = useRoute();
const patient = ref({} as Patient);
const records = ref([] as PatientHistory[])
const printArea = ref(null);

const fetchPatientHistory = async () => {
  records.value = await apiRequest.post(`patients/${route.params.id}/patient-history/all`, route.query);
}


const fetchPatient = async () => {
  patient.value = await apiRequest.get(`patients/${route.params.id}`);
}

const downloadPdf = () => {
  const filename = `${patient.value.lastname}-${patient.value.firstname}-${route.query.date_from}-${route.query.date_to}.pdf`;

  var opt = {
    margin: 1,
    filename,
    image: { type: 'jpeg', quality: 0.80 },
    html2canvas: { scale: 2, useCORS: true },
    jsPDF: { unit: 'mm', format: 'a4', orientation: 'p' }
  };

  html2pdf().set(opt).from(printArea.value).save();

}

onMounted(async () => {
  // get hospital
  const req = await apiRequest.get("hospital-details");
  if (req) hospital.setAll(req);
  await fetchPatient();
  await fetchPatientHistory();
})
</script>;

<template>
  <div class="page-wrapper w-[840px] my-5 m-auto">
    <div class="flex justify-between">
      <div class="px-[15px] flex justify-center py-2 gap-2">
        <ActionButton dark v-on:click="downloadPdf" :icon-src="DocumentArrowDownIcon">Download as PDF</ActionButton>
      </div>
      <ActionButton v-on:click="() => $router.go(-1)" :icon-src="XCircleIcon">Exit</ActionButton>
    </div>
    <div class="border-[1px] p-2 border-stone-950 m-2 gap-5 flex flex-col" ref="printArea">

      <div class="flex border-b-[1px] p-2 border-stone-600">
        <div class="m-auto flex gap-3 items-center">
          <img :src="`${hospital.asset_base_url}/files/logo/${hospital.hospital_logo}`" class="max-w-[100px]" alt="logo">
          <div>
            <div class="text-2xl font-bold" style="text-transform:uppercase">{{ hospital.hospital_name }}</div>
            <div>{{ hospital.hospital_address }}</div>
            <div>{{ hospital.hospital_phone }}</div>
          </div>
        </div>
      </div>

      <div class="text-2xl text-center">Patient Details</div>
      <table class="w-full border-separate border text-left">
        <tr>
          <th class="border-[1px] border-black p-1 w-[100px]">Name</th>
          <td class="border-[1px] border-black p-1 w-[170px]">{{ `${patient.title} ${patient.lastname},
                      ${patient.firstname} ${patient.middlename}` }}</td>
          <th class="border-[1px] border-black p-1 w-[100px]">Card No.</th>
          <td class="border-[1px] border-black p-1 w-[120px]">{{ patient.card_no }}</td>
          <th class="border-[1px] border-black p-1 w-[120px]">Phone No.</th>
          <td class="border-[1px] border-black p-1 w-[170px]">{{ patient.phone }}</td>
        </tr>
        <tr>
          <th class="border-[1px] border-black p-1">Email</th>
          <td class="border-[1px] border-black p-1">{{ patient.email }}</td>
          <th class="border-[1px] border-black p-1">Sex</th>
          <td class="border-[1px] border-black p-1">{{ patient.sex }}</td>
          <th class="border-[1px] border-black p-1">Marital Status</th>
          <td class="border-[1px] border-black p-1">{{ patient.marital_status }}</td>
        </tr>
        <tr>
          <th class="border-[1px] border-black p-1">Age.</th>
          <td class="border-[1px] border-black p-1">{{ patient.age ?? 0 }} yrs</td>
          <th class="border-[1px] border-black p-1">Appointment</th>
          <td class="border-[1px] border-black p-1">{{ patient.current_appointment }}</td>
          <th class="border-[1px] border-black p-1">Status</th>
          <td class="border-[1px] border-black p-1">{{ patient.current_status }}</td>
        </tr>
      </table>

      <!-- Record Proper -->
      <div class="text-2xl text-center">History Details</div>
      <table class="w-full border-separate border text-left">
        <tr>
          <th class="border-[1px] border-black p-1 w-[45px] text-center">S/N</th>
          <th class="border-[1px] border-black p-1 w-[150px]">Date</th>
          <th class="border-[1px] border-black p-1 w-[100px]">Type</th>
          <th class="border-[1px] border-black p-1">Details</th>
        </tr>
        <tr v-if="records.length" v-for="history, i in records" :key="patient.id">
          <td class="border-[1px] border-black p-1 text-center">{{ i + 1 }}</td>
          <td class="border-[1px] border-black p-1">{{ dayjs(history.date_time).format('DD-MM-YYYY hh:mm A') }}</td>
          <td class="border-[1px] border-black p-1"> {{ history.type }}</td>
          <td class="border-[1px] border-black p-1">
            <div>
              <div>{{ history.details.note }}</div>
              <div v-if="history.type == 'Admission'"> Ward: {{ history.details.ward_number }}; Room: {{
                history.details.room_number }}</div>
              <div v-if="history.type == 'Appointment'"> Appt: {{ history.details.appointment_type }}</div>
            </div>
          </td>
        </tr>
        <tr v-else>
          <td colspan="4">
            <div class="text-center mt-4 mb-4 pt-4">No records found</div>
          </td>
        </tr>
      </table>
    </div>
  </div>
</template>

<style scoped></style>