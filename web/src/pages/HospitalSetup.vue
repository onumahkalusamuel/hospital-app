<script lang="ts" >
import apiRequest from '../services/http/api-requests';

export default {
  name: 'HospitalSetup',
  async mounted() {
    try {
      const req = await apiRequest.get('hospital-details');
      if(req) { this.$router.push({name: 'login'})}

    } catch(e) {
      console.log(e)
    }
  },
  methods: {
    async hospitalSetup() {
      console.log(this.$refs.logo);
      console.log((this.$refs.logo as HTMLInputElement).files);
      const formData = new FormData(this.$refs.hospitalsetup as HTMLFormElement)
      const hospitalSetup = await apiRequest.postMulti('hospital-setup', formData);
      if(hospitalSetup?.message) {
        alert(hospitalSetup.message)
        this.$router.push({name: 'login'})
      }
    }
  }
}; </script>;

<template>
    <p class="" style="font-size: 1.2rem;">Hospital Setup</p>
    <form v-on:submit.prevent="hospitalSetup" ref="hospitalsetup" enctype="multipart/form-data">
      <div>
        <fluent-text-field name="hospital_name" placeholder="Hospital Name" class="w-100 mb-2"></fluent-text-field>
        <fluent-text-field name="hospital_address" placeholder="Physical Address" class="w-100 mb-2"></fluent-text-field>
        <fluent-text-field name="hospital_email" placeholder="Email Address" class="w-100 mb-2"></fluent-text-field>
        <fluent-text-field name="hospital_phone" placeholder="Phone Number" class="w-100 mb-2"></fluent-text-field>
        <!-- <fluent-text-field name="hospital_logo" placeholder="Logo" ref="logo" class="w-100 mb-2" type="file"></fluent-text-field> -->
        <input name="hospital_logo" placeholder="Logo" ref="logo" class="w-100 mb-2" type="file"/>
        <fluent-button class="w-100 mb-2" type="submit" appearance="accent">Submit</fluent-button>
      </div>
    </form>
</template>
