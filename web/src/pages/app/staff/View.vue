<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '../../../components/Breadcrumbs.vue';
import ActionButton from '../../../components/ActionButton.vue';
import PageHeader from '../../../components/PageHeader.vue';
import apiRequest from '../../../services/http/api-requests';
import { Staff } from '../../../interfaces'
import SelectField from '../../../components/form/SelectField.vue';
import TextField from '../../../components/form/TextField.vue';
import PrimaryButton from '../../../components/form/PrimaryButton.vue';
import SecondaryButton from '../../../components/form/SecondaryButton.vue';
import { useRoute, useRouter } from 'vue-router';
import { onMounted, ref } from 'vue';
import {UserIcon} from '@heroicons/vue/24/solid'
import { toasts } from '../../../stores/toasts';

const route = useRoute();
const router = useRouter();

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Staff", link: { name: "staff" } },
] as BreadcrumbItem[]);

const staff = ref({} as Staff);
const form = ref(null);

const fetchStaff = async () => {
  staff.value = await apiRequest.get(`staff/${route.params.id}`);
  
  breadcrumbs.value.push({ title: `${staff.value.lastname} ${staff.value.firstname}`, current: true });
}
const update = async () => {
  const formData = Object.fromEntries(new FormData(form as never as HTMLFormElement).entries())
  const create = await apiRequest.post('staff', formData);
  if(create.id) {
    toasts.addToast({message: "account created successfully.", type: 'success'});
    router.push({ name: 'view-staff', params: { id: create.id } })
  }
}

onMounted(async() => { await fetchStaff()});

</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader :title="`${staff.lastname} ${staff.firstname}`" :icon-src="UserIcon"></PageHeader>
    <div style="padding: 0 15px; display: flex; justify-content: space-between; border-bottom:1px solid #333">
      <div>
        <ActionButton v-on:click="() => $router.push({name: 'deliveries'})" :icon-src="UserIcon">Deliveries</ActionButton>
        <ActionButton :icon-src="UserIcon">Invoices</ActionButton>
      </div>
    </div>
    <div class="page-scroll-area">
      <form method="POST" v-on:submit.prevent="update" class="form flex flex-wrap gap-3 pl-4" ref="form">
          <div class="min-w-[250px]">
            <TextField label="First Name" placeholder="John" :value="staff.firstname"></TextField>
          </div>
          <div class="min-w-[250px]">
            <TextField label="Middle Name" placeholder="M." :value="staff.middlename"></TextField>
          </div>
          <div class="min-w-[250px]">
            <TextField label="Last Name" placeholder="Doe" :value="staff.lastname"></TextField>
          </div>
          <div class="min-w-[250px]">
            <SelectField label="Gender" name="sex" :options="[['Male'], ['Female']]" required :value="staff.sex" />
          </div>
          <div class="min-w-[250px]">
            <TextField label="Phone Number" name="phone" type="tel" :value="staff.phone"></TextField>
          </div>
          <div class="min-w-[250px]">
            <TextField label="Email Address" name="email" :value="staff.email"></TextField>
          </div>
        <div class="w-full"></div>
        <div class="min-w-[250px] flex space-x-3">
          <PrimaryButton type="submit" class="w-full">Update</PrimaryButton>
          <SecondaryButton type="button" class="w-full" v-on:click.prevent="()=>$router.push({name: 'staff'})">Cancel</SecondaryButton>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
</style>