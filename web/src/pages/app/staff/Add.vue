<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '../../../components/Breadcrumbs.vue';
import PageHeader from '../../../components/PageHeader.vue';
import apiRequest from '../../../services/http/api-requests';
import TextField from '../../../components/form/TextField.vue';
import PrimaryButton from '../../../components/form/PrimaryButton.vue';
import SecondaryButton from '../../../components/form/SecondaryButton.vue';
import SelectField from '../../../components/form/SelectField.vue';
import { ref } from 'vue';
import { router } from '../../../router';
import { UserPlusIcon } from '@heroicons/vue/24/outline';
import { Staff } from '../../../interfaces';
import { toasts } from '../../../stores/toasts';

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Staff",  link: { name: "staff" } },
  { title: "Add staff", current: true },
] as BreadcrumbItem[]);

const createForm = ref(null);


const create = async () => {
  const formData = Object.fromEntries(new FormData(createForm.value as never as HTMLFormElement).entries()) as never as Staff;
  formData.role *= 1;
  const create = await apiRequest.post('staff', formData);
  if(create.id) {
    toasts.addToast({message: "account created successfully.", type: 'success'});
    router.push({ name: 'view-staff', params: { id: create.id } })
  }
}
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="New Staff" :icon-src="UserPlusIcon"></PageHeader>
    <hr class="ml-4 mr-4" />
    <div class="page-scroll-area">
      <form method="POST" v-on:submit.prevent="create" class="form flex flex-wrap gap-3 pl-4" ref="createForm">
        <div class="min-w-[250px]"><SelectField label="Role" name="role" :options="[['Nurse', 3],['Doctor', 2]]" required/></div>
        <div class="min-w-[250px]"><TextField label="First Name" placeholder="John" name="firstname" required></TextField></div>
        <div class="min-w-[250px]"><TextField label="Middle Name" placeholder="M." name="middlename"></TextField></div>
        <div class="min-w-[250px]"><TextField label="Last Name" placeholder="Doe" name="lastname" required></TextField></div>
        <div class="min-w-[250px]"><SelectField label="Sex" name="sex" :options="[['Male'], ['Female']]" required/></div>
        <div class="min-w-[250px]"><TextField label="Phone Number" placeholder="08123456789" name="phone" type="tel"></TextField></div>
        <div class="min-w-[250px]"><TextField label="Email Address" placeholder="john.doe@example.com" name="email"></TextField></div>
        <div class="min-w-[250px]"><TextField label="Username" placeholder="Username" name="username" required></TextField></div>
        <div class="min-w-[250px]"><TextField label="Password" placeholder="Password" name="password" required></TextField></div>
        <div class="w-full"></div>
        <div class="min-w-[250px] flex space-x-3">
          <PrimaryButton type="submit" class="w-full">Submit</PrimaryButton>
          <SecondaryButton type="button" class="w-full" v-on:click.prevent="()=>$router.push({name: 'patients'})">Cancel</SecondaryButton>
        </div>
      </form>
    </div>
  </div>
</template>