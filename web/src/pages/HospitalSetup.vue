<script lang="ts" setup>
import apiRequest from '../services/http/api-requests';
import TextField from '../components/form/TextField.vue';
import PrimaryButton from '../components/form/PrimaryButton.vue';
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import { toasts } from '../stores/toasts';

const router = useRouter();
const hospitalsetup = ref(null);
const logo = ref(null);

onMounted(async() => {
  try {
      const req = await apiRequest.get('hospital-details');
      if(req) { router.push({name: 'login'})}

    } catch(e) {
      console.log(e)
    }
})

const hospitalSetup = async () => {
  const formData = new FormData(hospitalsetup.value as never as HTMLFormElement)
  const hospitalSetup = await apiRequest.postMulti('hospital-setup', formData);
  if(hospitalSetup?.message) {
    toasts.addToast({message: hospitalSetup.message, type: 'success'});
    router.push({name: 'login'})
  }
}
</script>;

<template>
  <p class="text-xl">Hospital Setup</p>
  <form v-on:submit.prevent="hospitalSetup" ref="hospitalsetup" enctype="multipart/form-data">
    <div>
      <TextField name="hospital_name" placeholder="Hospital Name" class="w-full mb-2"/>
      <TextField name="hospital_address" placeholder="Physical Address" class="w-full mb-2"/>
      <TextField name="hospital_email" placeholder="Email Address" class="w-full mb-2"/>
      <TextField name="hospital_phone" placeholder="Phone Number" class="w-full mb-2"/>
      <TextField label="Hospital Logo" name="hospital_logo" class="w-full mb-2 font-xl" ref="logo" type="file"/>
      <PrimaryButton class="w-full mb-2" type="submit">Submit</PrimaryButton>
    </div>
  </form>
</template>
