<script lang="ts" >
import apiRequest from '../services/http/api-requests';

export default {
  name: 'CreatAdmin',
  async mounted() {
    try {
      const req = await apiRequest.get('hospital-details');
      if(req) { this.$router.push({name: 'login'})}

    } catch(e) {
      console.log(e)
    }
  },
  methods: {
    async createAdmin() {
      const formData = new FormData(this.$refs.createadmin as HTMLFormElement)
      const createAdmin = await apiRequest.post('create-admin', Object.fromEntries(formData.entries()));
      if(createAdmin.message) {
        alert(createAdmin.message)
        this.$router.push({name: 'hospital-setup'})
      }
    }
  }
}; </script>;

<template>
    <p class="" style="font-size: 1.2rem;">Create Super Admin</p>
    <form method="post" v-on:submit.prevent="createAdmin" ref="createadmin">
      <div>
        <fluent-text-field name="firstname" placeholder="First Name" class="w-100 mb-2"></fluent-text-field>
        <fluent-text-field name="lastname" placeholder="Last Name" class="w-100 mb-2"></fluent-text-field>
        <fluent-text-field name="username" placeholder="Username" class="w-100 mb-2"></fluent-text-field>
        <fluent-text-field name="password" placeholder="Password" class="w-100 mb-2"></fluent-text-field>
        <fluent-button class="w-100 mb-2" type="submit" appearance="accent">Create Super Admin</fluent-button>
      </div>
    </form>
    <small>The super admin is the first user on the app. The account will have access to every section of the app.</small>
</template>
