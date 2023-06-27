<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '../../../components/Breadcrumbs.vue';
import ActionButton from '../../../components/ActionButton.vue';
import PageHeader from '../../../components/PageHeader.vue';
import apiRequest from '../../../services/http/api-requests';
import { Patient } from '../../../interfaces'
import dayjs from 'dayjs'
import SelectField from '../../../components/form/SelectField.vue';
import TextField from '../../../components/form/TextField.vue';
import PrimaryButton from '../../../components/form/PrimaryButton.vue';
import SecondaryButton from '../../../components/form/SecondaryButton.vue';
import { useRoute } from 'vue-router';
import { onMounted, ref } from 'vue';
import {UserIcon} from '@heroicons/vue/24/solid'
import { toasts, popupStore } from '../../../stores';
import AdmitPatientPopup from './popups/AdmitPatientPopup.vue';
import DischargePatientPopup from './popups/DischargePatientPopup.vue';
import AppointmentPopup from './popups/AppointmentPopup.vue';

const route = useRoute();

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Patients", link: { name: "patients" } },
] as BreadcrumbItem[]);

const patient = ref({} as Patient);
const form = ref(null);
const admitPatientPopupId = ref('admitpatient');
const dischargePatientPopupId = ref('dischargepatient');
const appointmentPopupId = ref('appointment');

const fetchPatient = async () => {
  patient.value = await apiRequest.get(`patients/${route.params.id}`);
  if(!breadcrumbs.value[breadcrumbs.value.length-1]?.current) {
    breadcrumbs.value.push({ title: `${patient.value.lastname} ${patient.value.firstname}`, current: true });
  }
}
const update = async () => {
  const formData = Object.fromEntries(new FormData(form.value as never as HTMLFormElement).entries())
  if(formData.dob) {
    formData.dob = dayjs(formData.dob as string).format('YYYY-MM-DDTHH:mm:ss[Z]');
  }
  const update = await apiRequest.put(`patients/${route.params.id}`, formData);
  if(update.message) {
    toasts.addToast({message: "account updated successfully.", type: 'success'});
    await fetchPatient()
  }
}

const showAdmitPatientPopup = () => { popupStore.id = admitPatientPopupId.value; popupStore.show = true;}
const showDischargePatientPopup = () => { popupStore.id = dischargePatientPopupId.value; popupStore.show = true;}
const showAppointmentPopup = () => { popupStore.id = appointmentPopupId.value; popupStore.show = true;}

onMounted(async() => { await fetchPatient()});

</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader
      :title="`${patient.lastname} ${patient.firstname}`"
      :subtitle="`${patient.sex} - ${patient.card_no}`"
      :icon-src="UserIcon">
      <div class="flex flex-col gap-2 font-bold text-right" style="justify-content: end;">
        <div class="flex items-center">
          <div class="min-w-[100px]">Appointment:</div>
          <div
            class="border-[1px] border-gray-500 px-2 py-1 rounded ml-5 min-w-[150px]"
            :class="patient.current_appointment == 'Emergency' ?
            'bg-red-500 text-white border-red-500':
            patient.current_appointment == 'Regular' ?
            'bg-yellow-300':
            'bg-gray-200'
            "
          >
            {{ patient.current_appointment || 'No appointment' }}
          </div>
        </div>
        <div class="flex items-center jusitify-end">
          <div class="min-w-[100px]">Status:</div>
          <div
            class="border-[1px] border-gray-500 px-2 py-1 rounded ml-5 min-w-[150px]"
            :class="patient.current_status == 'Admitted' ?
            'bg-red-500 text-white border-red-500':
            patient.current_status == 'Discharged' ?
            'bg-green-500 text-white border-green-500':
            'bg-gray-200'
            "
          >
            {{ patient.current_status || 'Not admitted' }}
          </div>
        </div>
      </div>
    </PageHeader>
    <div class="px-[15px] flex justify-between">
      <div class="flex gap-2">
        <ActionButton icon-class="text-white" classes="rounded bg-red-700 text-white hover:bg-red-500" v-on:click="() => $router.push({name: 'next-of-kin'})" :icon-src="UserIcon">Next of kin</ActionButton>
        <ActionButton icon-class="text-white" classes="bg-red-700 text-white hover:bg-red-500" v-on:click="() => $router.push({name: 'add-patient'})" :icon-src="UserIcon">Deliveries</ActionButton>
        <ActionButton icon-class="text-white" classes="bg-red-700 text-white hover:bg-red-500" v-if="'Admitted' !== patient.current_status" :icon-src="UserIcon" @click="showAdmitPatientPopup">Admit patient</ActionButton>
        <ActionButton v-if="'Admitted' === patient.current_status" :icon-src="UserIcon" @click="showDischargePatientPopup">Discharge patient</ActionButton>
        <ActionButton :icon-src="UserIcon" @click="showAppointmentPopup">Initiate appointment</ActionButton>
        <ActionButton v-on:click="() => $router.push({name: 'patient-history'})" :icon-src="UserIcon">Patient History</ActionButton>
        <!-- <ActionButton :icon-src="UserIcon">Reports</ActionButton> -->
      </div>
    </div>
    <div class="page-scroll-area">
      <form method="POST" v-on:submit.prevent="update" class="form flex flex-wrap gap-3 pl-4" ref="form">
        <div class="min-w-[250px]"><SelectField label="Title" name="title" :options="[['Mr.'],['Mrs.'],['Prof.'],['Dr.']]" :value="patient.title"/></div>
        <div class="min-w-[250px]">
          <TextField label="Last Name" placeholder="Doe" :value="patient.lastname"></TextField>
        </div>
        <div class="min-w-[250px]">
          <TextField label="Middle Name" placeholder="M." :value="patient.middlename"></TextField>
        </div>
        <div class="min-w-[250px]">
          <TextField label="First Name" placeholder="John" :value="patient.firstname"></TextField>
        </div>
          <div class="min-w-[250px]">
            <SelectField label="Gender" name="sex" :options="[['Male'], ['Female']]" required :value="patient.sex"/>
          </div>
          <div class="min-w-[250px]">
            <TextField label="Card No" :value="(patient.card_no as string)" readonly></TextField>
          </div>
          <div class="min-w-[250px]">
            <TextField label="Phone Number" name="phone" type="tel" :value="patient.phone"></TextField>
          </div>
          <div class="min-w-[250px]">
            <TextField label="Email Address" name="email" :value="patient.email"></TextField>
          </div>
          <div class="min-w-[250px]">
            <TextField label="Date of birth" type="date" name="dob" required :value="dayjs(patient.dob).format('YYYY-MM-DD')"></TextField>
          </div>
          <div class="min-w-[250px]">
            <SelectField label="Marital Status" name="marital_status" :options="[[''],['Single'],['Married'],['Divorced'],['Separated']]" :value="patient.marital_status"/>
          </div>
          <div class="min-w-[250px]">
            <TextField label="Home Address" placeholder="20 John Doe Avenue..." name="address" :value="patient.address" required></TextField>
          </div>
          <div class="min-w-[250px]">
            <TextField label="Occupation" placeholder="Engineer" name="occupation" :value="patient.occupation"></TextField>
          </div>
        <div class="w-full"></div>
          <div class="min-w-[250px] flex space-x-3">
            <PrimaryButton type="submit" class="w-full">Update</PrimaryButton>
            <SecondaryButton type="button" class="w-full" v-on:click.prevent="()=>$router.push({name: 'patients'})">Cancel</SecondaryButton>
          </div>
      </form>
    </div>
    <AdmitPatientPopup :popup-id="admitPatientPopupId" />
    <DischargePatientPopup :popup-id="dischargePatientPopupId" />
    <AppointmentPopup :popup-id="appointmentPopupId" />

  </div>
</template>