<script lang="ts" setup>
import { RouterView, RouterLink } from 'vue-router';
import apiRequest from '@/services/http/api-requests';
import { user, auth, hospital } from '@/stores';
import { UserIcon, SparklesIcon, Cog8ToothIcon, LockClosedIcon, QueueListIcon, BanknotesIcon, UserGroupIcon, UsersIcon } from '@heroicons/vue/24/outline';
import { onMounted } from 'vue';

onMounted(async () => {
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
})

</script>;

<template>
  <div class="flex flex-col h-screen">
    <header class="w-full flex toolbar bg-[#0078d4] h-[40px]">
      <div class="flex items-center text-white">
        <router-link class="header-icon-link font-bold text-[14px] px-3 hover:bg-[#1664a7]" :to="{name: 'dashboard'}">
          <SparklesIcon class="text-white h-6 w-6 mr-2" />
          {{ hospital.get('hospital_name') }}
        </router-link>
      </div>
      <div class="search-container flex items-center flex-1 justify-center font-bold uppercase">
        
        <router-link :class="$route.name == 'staff'? 'bg-[#00000033]':''" :to="{name:'staff'}" class="header-icon-link px-3" v-if="user.role == '1'">
          <UserIcon class="text-white h-5 w-5 mr-2"/>
          <span>Staff</span>
        </router-link>

        <router-link :class="$route.name == 'patients'? 'bg-[#00000033]':''" :to="{name:'patients'}" class="header-icon-link px-3">
          <UsersIcon class="text-white h-5 w-5 mr-2"/>
          <span>Patients</span>
        </router-link>

        <router-link :class="$route.name == 'deliveries'? 'bg-[#00000033]':''" :to="{name:'deliveries'}" class="header-icon-link px-3">
          <UserGroupIcon class="text-white h-5 w-5 mr-2"/>
          <span>Deliveries</span>
        </router-link>

        <router-link :class="$route.name == 'invoices'? 'bg-[#00000033]':''" :to="{name:'invoices'}" class="header-icon-link px-3">
          <BanknotesIcon class="text-white h-5 w-5 mr-2"/>
          <span>Billings</span>
        </router-link>

        <router-link :class="$route.name == 'reports'? 'bg-[#00000033]':''" :to="{name:'reports'}" class="header-icon-link px-3 font-bold">
          <QueueListIcon class="text-white h-5 w-5 mr-2"/>
          <span>REPORTS</span>
        </router-link>

      </div>
      <div style="display:flex; flex: 0 0 auto">
        <router-link class="header-icon-link w-[48px]" :to="{ name: 'settings' }">
          <cog8-tooth-icon class="text-white h-5 w-5"/>
        </router-link>
      </div>
      <a class="header-icon-link hover:bg-[#1664a7]">
        <div class="flex flex-col text-right px-3">
          <div class="avatarmenu-username">[{{ user.get('username') }}]</div>
        </div>
      </a>
      <a class="header-icon-link px-3" @click="() => {auth.setJwt(''); user.reset(); $router.push({name:'login'});}">
        <lock-closed-icon class="text-white h-5 w-5 mr-2"/>
        <span>Logout</span>
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
}

.avatar {
  height: 28px;
  width: 28px;
  border-radius: 28px;
  border:0;
}

.avatarmenu-username {
  white-space: nowrap;
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