<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import dayjs from 'dayjs'
import SelectField from '@/components/form/SelectField.vue';
import TextField from '@/components/form/TextField.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import { onMounted, ref } from 'vue';
import { Delivery } from '@/interfaces';
import {UserIcon} from '@heroicons/vue/24/solid'
import { useRoute, useRouter } from 'vue-router';
import { toasts } from '@/stores/toasts';

const route = useRoute();
const router = useRouter();

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Deliveries", link: { name: "deliveries" } },
  { title: 'Delivery details', current: true }
] as BreadcrumbItem[]);

const delivery = ref({} as Delivery);
const form = ref(null);

const fetchDelivery = async () => {
  delivery.value = await apiRequest.get(`deliveries/${route.params.id}`);
}
const update = async() => {
  const formData = Object.fromEntries(new FormData(form.value as never as HTMLFormElement).entries())
  if(formData.delivery_date_time) {
    formData.delivery_date_time = dayjs(formData.delivery_date_time as string).format('YYYY-MM-DDTHH:mm:ss[Z]');
  }
  
  const update = await apiRequest.put(`deliveries/${delivery.value.id}`, formData);
  if(update.message) {
    toasts.addToast({message: update.message, type: 'success'});
    router.push({ name: 'view-delivery', params: { id: update.id } })
  }
}

onMounted(async() => {await fetchDelivery()});
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="Delivery details" :icon-src="UserIcon"></PageHeader>

    <div style="padding: 0 15px; display: flex; justify-content: space-between; border-bottom:1px solid #333">
    </div>
    <div class="page-scroll-area">
      <div class="pl-4 flex items-center">
        <div class="border-[1px] border-gray-500 text-gray-500 p-1 mr-2 my-2">Patient:</div>
        <div class="text-md text-blue-500 hover:text-blue-700">
          <router-link :to="{name: 'view-patient', params:{id: delivery.patient_id}}">
            {{ `${delivery.patient?.lastname} ${delivery.patient?.firstname} - ${delivery.patient?.card_no}` }}
          </router-link>
        </div>
      </div>
      <form method="POST" v-on:submit.prevent="update" class="form flex flex-wrap gap-3 pl-4" ref="form">
        <div class="min-w-[250px]"><SelectField label="Delivery Mode" name="delivery_mode" :options="[['Vaginal'],['C-section']]" required :value="delivery.delivery_mode" /></div>
        <div class="min-w-[250px]"><SelectField label="Baby Sex" name="baby_sex" :options="[['Male'],['Female']]" required :value="delivery.baby_sex" /></div>
        <div class="min-w-[250px]"><TextField label="Baby Weight (kg)" placeholder="5.43" name="baby_weight" required type="number" step="any" :value="`${delivery.baby_weight}`"></TextField></div>
        <div class="min-w-[250px]"><SelectField label="Condition" name="condition" :options="[['Baby cried'],['Baby did not cry']]" :value="delivery.condition"/></div>
        <div class="min-w-[250px]"><TextField label="Delivery Date/Time" type="datetime-local" :value="dayjs(delivery.delivery_date_time).format('YYYY-MM-DD HH:mm')" readonly></TextField></div>
        <div class="min-w-[250px]"><TextField label="Note" placeholder="Safe delivery" name="note" :value="delivery.note"></TextField></div>
        <div class="w-full"></div>
        <div>
          <div class="min-w-[250px] flex space-x-3">
            <PrimaryButton type="submit" class="w-full">Update</PrimaryButton>
            <SecondaryButton type="button" class="w-full" v-on:click.prevent="()=>$router.push({name: 'patients'})">Cancel</SecondaryButton>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>