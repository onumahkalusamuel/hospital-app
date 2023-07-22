<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import dayjs from 'dayjs'
import TextField from '@/components/form/TextField.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import SelectField from '@/components/form/SelectField.vue';
import { ref } from 'vue';
import { router } from '@/router';
import { PlusCircleIcon } from '@heroicons/vue/24/outline';
import { Patient } from '@/interfaces';
import { toasts, popupStore } from '@/stores';
import SelectPatientPopup from './popups/SelectPatientPopup.vue';

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Deliveries",  link: { name: "deliveries" } },
  { title: "Add Delivery", current: true },
] as BreadcrumbItem[]);

const createForm = ref(null);
const selectPatientPopupId = ref('selectPatient');
const patient =ref({} as Patient);

const showSelectPatientPopup = () => { popupStore.id = selectPatientPopupId.value; popupStore.show = true;}

const create = async() => {
  const formData = Object.fromEntries(new FormData(createForm.value as never as HTMLFormElement).entries())
  if(formData.delivery_date_time) {
    formData.delivery_date_time = dayjs(formData.delivery_date_time as string).format('YYYY-MM-DDTHH:mm:ss[Z]');
  }

  const create = await apiRequest.post('deliveries', formData);
  if(create?.id) {
    toasts.addToast({message: "record created successfully.", type: 'success'});
    router.push({ name: 'view-delivery', params: { id: create.id } })
  }
}

</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="New Delivery" :icon-src="PlusCircleIcon"></PageHeader>
    <hr class="ml-4 mr-4" />
    <div class="page-scroll-area">
      <div class="flex mx-[15px] items-center h-[60px]">
        <div class="flex items-center w-[525px]">
          <div class="font-bold">Patient:&nbsp</div>
          <div v-if="patient.id"> {{ `${patient.lastname} ${patient.firstname} - ${patient.card_no} ` }}</div>
        </div>
        <div class="w-[50%] flex">
          <div><PrimaryButton @click="showSelectPatientPopup">Select Patient</PrimaryButton></div>
        </div>
      </div>
      <form method="POST" v-on:submit.prevent="create" class="form flex flex-wrap gap-3 pl-4" ref="createForm">
        <input type="hidden" name="patient_id" v-model="patient.id" />
        <div class="min-w-[250px]"><SelectField label="Delivery Mode" name="delivery_mode" :options="[['Vaginal'],['C-section']]" required /></div>
        <div class="min-w-[250px]"><SelectField label="Baby Sex" name="baby_sex" :options="[['Male'],['Female']]" required /></div>
        <div class="min-w-[250px]"><TextField label="Baby Weight (kg)" placeholder="5.43" name="baby_weight" required type="number" step="any"></TextField></div>
        <div class="min-w-[250px]"><SelectField label="Condition" name="condition" :options="[['Baby cried'],['Baby did not cry']]" /></div>
        <div class="min-w-[250px]"><TextField label="Delivery Date/Time" name="delivery_date_time" type="datetime-local"></TextField></div>
        <div class="min-w-[250px]"><TextField label="Note" placeholder="Safe delivery" name="note"></TextField></div>
        <div class="w-full"></div>
        <div class="min-w-[250px] flex space-x-3">
          <PrimaryButton type="submit" class="w-full">Submit</PrimaryButton>
          <SecondaryButton type="button" class="w-full" v-on:click.prevent="()=>$router.go(-1)">Cancel</SecondaryButton>
        </div>
      </form>
    </div>

    <SelectPatientPopup :popup-id="selectPatientPopupId" v-model="patient" />
  </div>
</template>