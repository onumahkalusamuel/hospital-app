import { RouteRecordRaw } from 'vue-router';
import MainLayout from '../layouts/MainLayout.vue';
import AppLayout from '../layouts/AppLayout.vue';
import Activate from '../pages/Activate.vue';
import Login from '../pages/Login.vue';
import CreateAdmin from '../pages/CreateAdmin.vue';
import HospitalSetup from '../pages/HospitalSetup.vue';
import Dashboard from '../pages/app/Dashboard.vue';

export const routes: RouteRecordRaw[] = [
  {
    path: '/', component: MainLayout, children: [
      { path: '', redirect: { name: 'login' } },
      { path: 'activate/:code?', component: Activate, name: 'activate' },
      { path: 'create-admin', component: CreateAdmin, name: 'create-admin' },
      { path: 'hospital-setup', component: HospitalSetup, name: 'hospital-setup' },
    ]
  },
  { path: '/login', component: Login, name: 'login' },
  {
    path: '/app/', component: AppLayout, children: [
      { path: 'dashboard', component: Dashboard, name: 'dashboard' },
      { path: ':pathMatch(.*)', redirect: { name: 'dashboard' } },
    ]
  },
  { path: '/:pathMatch(.*)*', redirect: { name: 'login' } },
];