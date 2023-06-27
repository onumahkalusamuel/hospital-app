<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '../../../components/Breadcrumbs.vue';
import PageHeader from '../../../components/PageHeader.vue';
import apiRequest from '../../../services/http/api-requests';
import dayjs from 'dayjs'
import TextField from '../../../components/form/TextField.vue';
import PrimaryButton from '../../../components/form/PrimaryButton.vue';
import SecondaryButton from '../../../components/form/SecondaryButton.vue';
import SelectField from '../../../components/form/SelectField.vue';
import { onMounted, ref } from 'vue';
import { router } from '../../../router';
import { PlusCircleIcon } from '@heroicons/vue/24/outline';
import { Patient } from '../../../interfaces';
import { toasts } from '../../../stores/toasts';

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Deliveries",  link: { name: "deliveries" } },
  { title: "Add Delivery", current: true },
] as BreadcrumbItem[]);

const createForm = ref(null);
const patients = ref([] as [string, string?][]);

const create = async() => {
  const formData = Object.fromEntries(new FormData(createForm.value as never as HTMLFormElement).entries())
  if(formData.delivery_date_time) {
    formData.delivery_date_time = dayjs(formData.delivery_date_time as string).format('YYYY-MM-DDTHH:mm:ss[Z]');
  }
  
  (formData.baby_weight as any) = Math.ceil(formData.baby_weight as any * 1000)

  const create = await apiRequest.post('deliveries', formData);
  if(create?.id) {
    toasts.addToast({message: "record created successfully.", type: 'success'});
    router.push({ name: 'view-delivery', params: { id: create.id } })
  }
}

onMounted(async() => {
  const pats = (await apiRequest.get('patients')) as Patient[];
  if(pats?.length) {
    pats.forEach(pat => {patients.value.push([`${pat.lastname} - ${pat.card_no}`, pat.id]);});
  }
})
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="New Delivery" :icon-src="PlusCircleIcon"></PageHeader>
    <hr class="ml-4 mr-4" />
    <div class="page-scroll-area">
      <form method="POST" v-on:submit.prevent="create" class="form flex flex-wrap gap-3 pl-4" ref="createForm">
          <div class="min-w-[250px]"><SelectField label="Patient" name="patient_id" :options="patients" required /></div>
          <div class="min-w-[250px]"><SelectField label="Delivery Mode" name="delivery_mode" :options="[['Vaginal'],['C-section']]" required /></div>
          <div class="min-w-[250px]"><SelectField label="Baby Sex" name="baby_sex" :options="[['Male'],['Female']]" required /></div>
          <div class="min-w-[250px]"><TextField label="Baby Weight" placeholder="5.43" name="baby_weight" required type="number" step="any"></TextField></div>
          <div class="min-w-[250px]"><SelectField label="Baby Weight Unit" name="baby_weight_unit" :options="[['kg'],['lbs']]" /></div>
          <div class="min-w-[250px]"><SelectField label="Condition" name="baby_sex" :options="[['Baby cried'],['Baby did not cry']]" /></div>
          <div class="min-w-[250px]"><TextField label="Delivery Date/Time" name="delivery_date_time" type="datetime-local"></TextField></div>
          <div class="min-w-[250px]"><TextField label="Note" placeholder="Safe delivery" name="note"></TextField></div>
          <div class="w-full"></div>
          <div class="min-w-[250px] flex space-x-3">
            <PrimaryButton type="submit" class="w-full">Submit</PrimaryButton>
            <SecondaryButton type="button" class="w-full" v-on:click.prevent="()=>$router.push({name: 'deliveries'})">Cancel</SecondaryButton>
          </div>
      </form>
    </div>
  </div>
</template>