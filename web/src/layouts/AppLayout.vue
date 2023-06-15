<script lang="ts">
import { RouterView, RouterLink } from 'vue-router';
import apiRequest from '../services/http/api-requests';
import { user } from '../stores/user';
import { hospital } from '../stores/hospital';
import TextField from '../components/form/TextField.vue';
import {
  MagnifyingGlassIcon,
  Bars4Icon,
  BellAlertIcon,
  Cog8ToothIcon,
} from '@heroicons/vue/24/outline';


export default {
    name: "AppLayout",
    methods: {},
    setup() {
        return {
            user,
            hospital
        };
    },
    async mounted() {
        try {
            // get hospital
            const req = await apiRequest.get("hospital-details");
            if (req)
                hospital.setAll(req);
            // get user
            const profile = await apiRequest.get("profile");
            if (profile)
                user.setAll(profile);
        }
        catch (e) {
            console.log(e);
        }
    },
    components: { TextField, MagnifyingGlassIcon, Bars4Icon, BellAlertIcon, Cog8ToothIcon, RouterLink }
}; </script>;

<template>
  <div class="flex flex-col h-screen">
    <header class="w-full flex toolbar bg-[#0078d4] h-[40px]">
      <a class="header-icon-link w-[48px] hover:bg-[#1664a7]">
        <bars4-icon class="text-white h-6 w-6" />
      </a>
      <router-link class="header-icon-link font-bold text-[14px] px-3 hover:bg-[#1664a7]" :to="{name: 'dashboard'}">
        {{ hospital.get('hospital_name') }}
      </router-link>   
      <div class="search-container flex flex-col h-full items-center">
        <text-field name="search" class="max-w-[800px] border-none mb-none" placeholder="Search">
          <template #prepend>
            <magnifying-glass-icon class="h-5 w-5 text-gray-500"/>
          </template>
        </text-field>
      </div>
      <div style="display:flex; flex: 0 0 auto">
        <a class="header-icon-link w-[48px]">
          <bell-alert-icon class="text-white h-5 w-5"/>
        </a>
        <a class="header-icon-link w-[48px]">
          <cog8-tooth-icon class="text-white h-5 w-5"/>
        </a>
      </div>
      <a class="header-icon-link hover:bg-[#1664a7]">
        <div class="avatarmenu-container">
          <div class="avatarmenu-username">{{ user.get('username') }}</div>
          <div class="avatarmenu-userid">{{  user.get('id') }}</div>
        </div>
        <div class="avatarmenu-image-container">
          <img src="/avatar.png" alt="avatar" class="avatar"/>
        </div>
      </a>
    </header>
    <main class="flex-1 flex-scroll">
      <router-view></router-view>
    </main>
  </div>
</template>

<style scoped>
.toolbar {
  flex: 0 0 auto;
  box-sizing: border-box;
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
}
.header-icon-link:hover {
  background-color: #1664a7;
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
</style>