<script lang="ts" >
import { RouterView } from 'vue-router';
import apiRequest from '../services/http/api-requests';
export default {
  name: 'AppLayout',
  data: () => ({
    hospital_name: "",
    hospital_address: "",
    hospital_logo: ""
  }),
    async mounted() {
    // check installation first
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
  }
}; </script>;

<template>
  <div class="app">
    <header class="w-100 toolbar">
      <a class="header-icon-link w-48">
        <img class="white-svg" src="/icons/waffle.svg" alt="menu"/>
      </a>
      <a class="header-icon-link site-title">
        {{ hospital_name }}
      </a>   
      <div class="search-container">
        <fluent-text-field class="w-100" placeholder="Search" style="max-width:800px; opacity: 0.8;" dense>
          <i slot="start" class="flex">
            <img src="/icons/search.svg" alt="menu"/>
          </i>
        </fluent-text-field>
      </div>
      <div style="display:flex; flex: 0 0 auto">
        <a class="header-icon-link w-48">
          <img class="white-svg" src="/icons/notification.svg" alt="notification"/>
        </a>
        <a class="header-icon-link w-48">
          <img class="white-svg" src="/icons/gear.svg" alt="settings"/>
        </a>
      </div>
      <a class="header-icon-link">
        <div class="avatarmenu-container">
          <div class="avatarmenu-username">Onumah Kalu Samuel</div>
          <div class="avatarmenu-userid">45f4-54d3-43ac43453cde-4435</div>
        </div>
        <div class="avatarmenu-image-container">
          <img src="/avatar.png" alt="avatar" class="avatar"/>
        </div>
      </a>
    </header>
    <main style="">
      <router-view></router-view>
    </main>
    <footer>
      <a class="header-icon-link w-48">
        <img class="white-svg" src="/icons/waffle.svg" alt="menu"/>
      </a>
      <a class="header-icon-link site-title">
        {{ hospital_name }}
      </a>   
      <div class="search-container">
        <fluent-text-field class="w-100" placeholder="Search" style="max-width:800px; opacity: 0.8;" dense>
          <i slot="start" class="flex">
            <img src="/icons/search.svg" alt="menu"/>
          </i>
        </fluent-text-field>
      </div>
      <div style="display:flex; flex: 0 0 auto">
        <a class="header-icon-link w-48">
          <img class="white-svg" src="/icons/notification.svg" alt="notification"/>
        </a>
        <a class="header-icon-link w-48">
          <img class="white-svg" src="/icons/gear.svg" alt="settings"/>
        </a>
      </div>
      <a class="header-icon-link">
        <div class="avatarmenu-container">
          <div class="avatarmenu-username">Onumah Kalu Samuel</div>
          <div class="avatarmenu-userid">45f4-54d3-43ac43453cde-4435</div>
        </div>
        <div class="avatarmenu-image-container">
          <img src="/avatar.png" alt="avatar" class="avatar"/>
        </div>
      </a>
    </footer>
  </div>
</template>

<style scoped>
.app {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
main {
  flex-grow: 1;
  overflow:scroll;
}
.toolbar {
  display: flex;
  flex: 0 0 auto;
  height: 40px;
  box-sizing: border-box;
  background-color: #0078d4;
}
.header-icon-link {
  height: 40px;
  display:flex;
  flex: 0 0 auto;
  align-items: center;
  justify-content: center;
  color:white;
  text-decoration: none;
  cursor:pointer;
  /* filter: invert(100%) sepia(100%) saturate(0%) hue-rotate(288deg) brightness(102%) contrast(102%); */
}
.header-icon-link:hover {
  background-color: #1664a7;
}
.w-48 {
  width: 48px;
}
.site-title {
  padding: 0 20px;
  font-weight: 600;
  font-size: 14px;
}
.search-container {
  margin-right: 15px;
  min-width: 200px;
  margin-left: 7px;
  display: flex;
  flex-grow: 1;
  align-items: center;
  justify-content: center;
}
.flex {
  display: flex;
}
.avatar {
  height: 28px;
  width: 28px;
  border-radius: 28px;
  border:0;
}
.avatarmenu-container {
padding-left: 20px ;text-align:right; flex:1 1 auto
}
.avatarmenu-username {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  font-size: 14px;
  line-height: normal;
  max-width: 160px;
}
.avatarmenu-userid {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  font-weight: 600;
  font-size: 10px;
  line-height: normal;
  max-width: 160px;
}
.avatarmenu-image-container {
  height: 28px;
  width: 28px;
  flex: 0 0 auto;
  padding-left: 6px;
  padding-right: 6px;
  box-sizing: content-box;
}
.white-svg {
  filter: invert(100%) sepia(100%) saturate(0%) hue-rotate(288deg) brightness(102%) contrast(102%);
}

footer {
  display: flex;
  flex: 0 0 auto;
  height: 40px;
  box-sizing: border-box;
  background-color: #0078d4;
}
</style>