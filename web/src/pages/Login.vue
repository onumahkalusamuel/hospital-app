<script lang="ts" setup>
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import TextField from '@/components/form/TextField.vue';
import apiRequest from '@/services/http/api-requests';
import { auth } from '@/stores/auth';
import { toasts } from '@/stores/toasts';
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';

const hospital_name = ref('');
const hospital_address = ref('');
const hospital_logo = ref('');
const loginForm = ref(null);
const router = useRouter();

onMounted(async () => {
  auth.setJwt("");
  try {
      const req = await apiRequest.get("hospital-details");
      if (req) {
          hospital_name.value = req.hospital_name;
          hospital_address.value = req.hospital_address;
          hospital_logo.value = req.hospital_logo;
      }
  }
  catch (e) {
      console.log(e);
  }
})

const login = async () => {
  const formData = new FormData(loginForm.value as never as HTMLFormElement);
  const login = await apiRequest.post("login", Object.fromEntries(formData.entries()));
  if (login.message) {
    auth.setJwt(login.jwt);
    toasts.addToast({ message: 'Login successful', title: 'Success', type: 'success' })
    router.push({ name: "dashboard" });
  }
}
</script>;

<template>
  <div class="h-screen flex flex-col sm:flex-row">
    <div class="bg-[url('/hospital-image.png')] bg-cover bg-no-repeat flex-1"></div>
    <div class="p-5 flex flex-col w-[400px] min-h-[85vh] justify-between mx-auto sm:py-10">
      <div class="flex justify-between items-center">
        <div class="w-full">
          <div class="text-2xl">{{ hospital_name }}</div>
          <div v-if="hospital_address" class="mt-1">{{ hospital_address }}</div>
        </div>
        <div class="logo">
          <img :src="hospital_logo" class="max-w-[70px]" alt="logo">
        </div>
      </div>
      <div class="form-proper">
        <p class="text-xl mt-10 mb-4">Login to continue</p>
        <form v-on:submit.prevent="login" ref="loginForm" autocomplete="off">
          <TextField class="mb-2" name="username" placeholder="Username" required/>
          <TextField class="mb-2" name="password" placeholder="Password" type="password" required />
          <PrimaryButton type="submit">Login</PrimaryButton>
        </form>
        <div class="text-center mt-5"><small>HCMS v1.0.0</small></div>
      </div>
    </div>
  </div>
</template>