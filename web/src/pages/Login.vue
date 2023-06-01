<script lang="ts" >
import apiRequest from '../services/http/api-requests';
import { auth } from '../stores/auth';

export default {
  name: 'Login',
  data: () => ({
    hospital_name: "",
    hospital_address: "",
    hospital_logo: ""
  }),
  async mounted() {
    auth.setJwt('');
    try {
      const req = await apiRequest.get('hospital-details');
      if(req) {
        this.hospital_name = req.hospital_name
        this.hospital_address = req.hospital_address  
        this.hospital_logo = req.hospital_logo
      }
    } catch(e) {
      console.log(e)
    }
  },
  methods: {
    async login() {
      const formData = new FormData(this.$refs.login as HTMLFormElement)
      const login = await apiRequest.post('login', Object.fromEntries(formData.entries()));
      if(login.message) {
        auth.setJwt(login.jwt);
        this.$router.push({name: 'dashboard'})
      }
    }
  }
}; </script>;

<template>
  <div class="login-container">
    <div class="image-container"></div>
    <div class="form-container">
      <div class="identity">
        <div class="name-addresss w-100">
          <div class="name">{{ hospital_name }}</div>
          <div v-if="hospital_address" class="address">{{ hospital_address }}</div>
        </div>
        <div class="logo">
          <img :src="hospital_logo" alt="logo" style="max-width: 70px;">
        </div>
      </div>
      <div class="form-proper">
        <form v-on:submit.prevent="login" ref="login">
          <p style="font-size: 20px;">Login to continue</p>
          <div >
            <fluent-text-field name="username" placeholder="Username" class="w-100 mb-2"></fluent-text-field>
            <fluent-text-field name="password" placeholder="Password" class="w-100 mb-2" type="password"></fluent-text-field>
            <fluent-button class="w-100 mb-2" type="submit" appearance="accent">Login</fluent-button>
          </div>
          <div class="text-center mt-3"><small>HCMS v1.0.0</small></div>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>
.form-container {
  padding: 30px;
  display: flex;
  flex-direction: column;
  max-width: 350px;
  min-height: 75vh;
  justify-content: space-between;
  margin: auto;
}
@media screen and (max-width: 520px) {
  .login-container {
    flex-direction: column;
  }
}
.login-container {
  display:flex;
  height:100vh;
}
.image-container {
  flex-grow: 1;
  background-image: url('/hospital-image.png');
  background-repeat: no-repeat;
  background-size: cover;
}
.identity {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.name {
  font-size: 22px;
}
.address {
  margin-top:5px;
  font-size: 14px;
}
</style>