<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '../../../components/Breadcrumbs.vue';
import PageHeader from '../../../components/PageHeader.vue';
import apiRequest from '../../../services/http/api-requests';
import dayjs from 'dayjs'
import TextField from '../../../components/form/TextField.vue';
import PrimaryButton from '../../../components/form/PrimaryButton.vue';
import SecondaryButton from '../../../components/form/SecondaryButton.vue';
import SelectField from '../../../components/form/SelectField.vue';
import { ref } from 'vue';
import { router } from '../../../router';
import { UserPlusIcon } from '@heroicons/vue/24/outline';

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Patients",  link: { name: "patients" } },
  { title: "Add Patient", current: true },
] as BreadcrumbItem[]);

const createForm = ref(null);


const create = async () => {
  const formData = Object.fromEntries(new FormData(createForm.value as never as HTMLFormElement).entries())
  if(formData.dob) {
    formData.dob = dayjs(formData.dob as string).format('YYYY-MM-DDTHH:mm:ss[Z]');
  }
  const create = await apiRequest.post('patients', formData);
  if(create.id) {
    alert("account created successfully.")
    router.push({ name: 'view-patient', params: { id: create.id } })
  }
}
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="New Patient" :icon-src="UserPlusIcon"></PageHeader>
    <div style="padding: 0 15px; display: flex; justify-content: space-between; border-top:1px solid #333">
    </div>
    <hr class="ml-4 mr-4" />
    <div class="page-scroll-area">
      <form method="POST" v-on:submit.prevent="create" class="form flex flex-wrap gap-3 pl-4" ref="createForm">
          <div class="min-w-[250px]"><SelectField label="Title" name="title" :options="[['Mr.'],['Mrs.'],['Prof.'],['Dr.']]" /></div>
          <div class="min-w-[250px]"><TextField label="First Name" placeholder="John" name="firstname" required></TextField></div>
          <div class="min-w-[250px]"><TextField label="Middle Name" placeholder="M." name="middlename"></TextField></div>
          <div class="min-w-[250px]"><TextField label="Last Name" placeholder="Doe" name="lastname" required></TextField></div>
          <div class="min-w-[250px]"><SelectField label="Gender" name="sex" :options="[['Male'], ['Female']]" required/></div>
          <div class="min-w-[250px]"><TextField label="Card No" placeholder="HC0090" name="card_no"></TextField></div>
          <div class="min-w-[250px]"><TextField label="Phone Number" placeholder="08123456789" name="phone" type="tel"></TextField></div>
          <div class="min-w-[250px]"><TextField label="Email Address" placeholder="john.doe@example.com" name="email"></TextField></div>
          <div class="min-w-[250px]"><TextField label="Date of birth" type="date" name="dob" required></TextField></div>
          <div class="min-w-[250px]"><TextField label="Home Address" placeholder="20 John Doe Avenue..." name="address" required></TextField></div>
          <div class="min-w-[250px]"><TextField label="Occupation" placeholder="Engineer" name="occupation"></TextField></div>
          <div class="min-w-[250px]"><TextField label="Appointment Type" placeholder="emergency" name="current_appointment" required></TextField></div>
          <div class="min-w-[250px]"><TextField label="Patient Status" placeholder="outpatient" name="current_status" required></TextField></div>
          <div class="w-full"></div>
          <div class="min-w-[250px] flex space-x-3">
            <PrimaryButton type="submit" class="w-full">Submit</PrimaryButton>
            <SecondaryButton type="button" class="w-full" v-on:click.prevent="()=>$router.push({name: 'patients'})">Cancel</SecondaryButton>
          </div>
      </form>
    </div>
  </div>
</template>