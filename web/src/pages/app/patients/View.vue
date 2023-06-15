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
import { useRoute, useRouter } from 'vue-router';
import { onMounted, ref } from 'vue';
import {UserIcon} from '@heroicons/vue/24/solid'

const route = useRoute();
const router = useRouter();

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Patients", link: { name: "patients" } },
] as BreadcrumbItem[]);

const patient = ref({} as Patient);
const form = ref(null);

const fetchPatient = async () => {
  patient.value = await apiRequest.get(`patients/${route.params.id}`);
  
  breadcrumbs.value.push({ title: `${patient.value.lastname} ${patient.value.firstname}`, current: true });
}
const update = async () => {
  const formData = Object.fromEntries(new FormData(form as never as HTMLFormElement).entries())
  if(formData.dob) {
    formData.dob = dayjs(formData.dob as string).format('YYYY-MM-DDTHH:mm:ss[Z]');
  }
  const create = await apiRequest.post('patients', formData);
  if(create.id) {
    alert("account created successfully.")
    router.push({ name: 'view-patient', params: { id: create.id } })
  }
}

onMounted(async() => { await fetchPatient()});

</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader
      :title="`${patient.lastname} ${patient.firstname}`"
      :subtitle="`${patient.sex} - ${patient.card_no}`"
      :icon-src="UserIcon"></PageHeader>

    <div style="padding: 0 15px; display: flex; justify-content: space-between; border-bottom:1px solid #333">
      <div>
        <ActionButton v-on:click="() => $router.push({name: 'next-of-kin'})" icon-src="/icons/person-plus-fill.svg">Next of kin</ActionButton>
        <ActionButton v-on:click="() => $router.push({name: 'add-patient'})" icon-src="/icons/basket.svg">Deliveries</ActionButton>
        <ActionButton icon-src="/icons/receipt.svg">Invoices</ActionButton>
        <ActionButton icon-src="/icons/receipt.svg">Reports</ActionButton>
        <ActionButton icon-src="/icons/receipt.svg">History</ActionButton>
      </div>
    </div>
    <div class="page-scroll-area">
      <form method="POST" v-on:submit.prevent="update" class="form flex flex-wrap gap-3 pl-4" ref="form">
        <div class="min-w-[250px]"><SelectField label="Title" name="title" :options="[['Mr.'],['Mrs.'],['Prof.'],['Dr.']]"/></div>
          <div class="min-w-[250px]">
            <TextField label="First Name" placeholder="John" :value="patient.firstname"></TextField>
          </div>
          <div class="min-w-[250px]">
            <TextField label="Middle Name" placeholder="M." :value="patient.middlename"></TextField>
          </div>
          <div class="min-w-[250px]">
            <TextField label="Last Name" placeholder="Doe" :value="patient.lastname"></TextField>
          </div>
          <div class="min-w-[250px]">
            <SelectField label="Gender" name="sex" :options="[['Male'], ['Female']]" required/>
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
            <TextField label="Home Address" placeholder="20 John Doe Avenue..." name="address" :value="patient.address" required></TextField>
          </div>
          <div class="min-w-[250px]">
            <TextField label="Occupation" placeholder="Engineer" name="occupation" :value="patient.occupation"></TextField>
          </div>
          <div class="min-w-[250px]">
            <TextField label="Appointment Type" placeholder="emergency" name="current_appointment" :value="patient.current_appointment" required></TextField>
          </div>
          <div class="min-w-[250px]">
            <TextField label="Patient Status" placeholder="outpatient" name="current_status" :value="patient.current_status" required></TextField>
          </div>
        <div class="w-full"></div>
          <div class="min-w-[250px] flex space-x-3">
            <PrimaryButton type="submit" class="w-full">Update</PrimaryButton>
            <SecondaryButton type="button" class="w-full" v-on:click.prevent="()=>$router.push({name: 'patients'})">Cancel</SecondaryButton>
          </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
</style>