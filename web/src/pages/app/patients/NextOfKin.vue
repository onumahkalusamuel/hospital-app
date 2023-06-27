<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '../../../components/Breadcrumbs.vue';
import PageHeader from '../../../components/PageHeader.vue';
import apiRequest from '../../../services/http/api-requests';
import TextField from '../../../components/form/TextField.vue';
import PrimaryButton from '../../../components/form/PrimaryButton.vue';
import SecondaryButton from '../../../components/form/SecondaryButton.vue';
import SelectField from '../../../components/form/SelectField.vue';
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { UserPlusIcon } from '@heroicons/vue/24/outline';
import { Patient } from '../../../interfaces';
import { toasts } from '../../../stores/toasts';

const route = useRoute();
const router = useRouter();

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Patients",  link: { name: "patients" } },
] as BreadcrumbItem[]);

const form = ref(null);
const patient = ref({} as Patient);
const relationships = ref([['Father'], ['Mother'], ['Sister'], ['Brother'], ['Uncle'], ['Niece'], ['Nephew'], ['Others']] as [string, string?][]);

const updateNextOfKin = async () => {
  const formData = Object.fromEntries(new FormData(form.value as never as HTMLFormElement).entries())
  const update = await apiRequest.post(`patients/${route.params.id}/next-of-kin`, formData);
  if(update?.message) {
    toasts.addToast({message: "next of kin record updated successfully.", type: 'success'});
    router.push({ name: 'view-patient', params: { id: route.params.id } })
  }
}

onMounted(async() => {
  patient.value = await apiRequest.get(`patients/${route.params.id}`);
  if(patient.value) {
   breadcrumbs.value.push({
    title: `${patient.value.lastname} ${patient.value.firstname} - ${patient.value.card_no}`,
    link: {name: 'view-patient', params: {id: patient.value.id}}
  });
  }
});
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="Next of Kin" :icon-src="UserPlusIcon"></PageHeader>
    <hr class="ml-4 mr-4" />
    <div class="page-scroll-area">
      <form method="POST" v-on:submit.prevent="updateNextOfKin" class="form flex flex-wrap gap-3 pl-4" ref="form">
          <div class="min-w-[250px]"><SelectField label="Title" name="title" :options="[['Mr.'],['Mrs.'],['Prof.'],['Dr.']]" /></div>
          <div class="min-w-[250px]"><TextField label="First Name" placeholder="John" name="firstname" required :value="patient.next_of_kin?.firstname"></TextField></div>
          <div class="min-w-[250px]"><TextField label="Middle Name" placeholder="M." name="middlename" :value="patient.next_of_kin?.middlename"></TextField></div>
          <div class="min-w-[250px]"><TextField label="Last Name" placeholder="Doe" name="lastname" required :value="patient.next_of_kin?.lastname"></TextField></div>
          <div class="min-w-[250px]"><SelectField label="Gender" name="sex" :options="[['Male'], ['Female']]" required/></div>
          <div class="min-w-[250px]"><TextField label="Phone Number" placeholder="08123456789" name="phone" type="tel" :value="patient.next_of_kin?.phone"></TextField></div>
          <div class="min-w-[250px]"><TextField label="Email Address" placeholder="john.doe@example.com" name="email" :value="patient.next_of_kin?.email"></TextField></div>
          <div class="min-w-[250px]"><TextField label="Home / Business Address" placeholder="20 John Doe Avenue..." name="address" required :value="patient.next_of_kin?.address"></TextField></div>
          <div class="min-w-[250px]"><TextField label="Occupation" placeholder="Engineer" name="occupation" :value="patient.next_of_kin?.occupation"></TextField></div>
          <div class="min-w-[250px]"><SelectField label="Relationship" name="relationship" :options="relationships" required :value="patient.next_of_kin?.relationship" /></div>
          <div class="w-full"></div>
          <div class="min-w-[250px] flex space-x-3">
            <PrimaryButton type="submit" class="w-full">Save</PrimaryButton>
            <SecondaryButton type="button" class="w-full" v-on:click.prevent="()=>$router.push({name: 'view-patient', params: {id: route.params.id}})">Cancel</SecondaryButton>
          </div>
      </form>
    </div>
  </div>
</template>