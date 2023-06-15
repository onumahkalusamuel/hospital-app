import { RouteRecordRaw } from 'vue-router';
import { appGuard } from '../guards/app-guard';

import MainLayout from '../layouts/MainLayout.vue';
import AppLayout from '../layouts/AppLayout.vue';

import Activate from '../pages/Activate.vue';
import Login from '../pages/Login.vue';
import CreateAdmin from '../pages/CreateAdmin.vue';

import HospitalSetup from '../pages/HospitalSetup.vue';
import Dashboard from '../pages/app/Dashboard.vue';

import Patients from '../pages/app/patients/Index.vue';
import AddPatient from '../pages/app/patients/Add.vue';
import ViewPatient from '../pages/app/patients/View.vue';
import NextOfKin from '../pages/app/patients/NextOfKin.vue';

import Deliveries from '../pages/app/deliveries/Index.vue';
import AddDelivery from '../pages/app/deliveries/Add.vue';
import ViewDelivery from '../pages/app/deliveries/View.vue';

import Staff from '../pages/app/staff/Index.vue';
import AddStaff from '../pages/app/staff/Add.vue';
import ViewStaff from '../pages/app/staff/View.vue';

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
    path: '/app/', component: AppLayout, beforeEnter: appGuard, children: [
      { path: 'dashboard', component: Dashboard, name: 'dashboard' },

      { path: 'patients', component: Patients, name: 'patients' },
      { path: 'patients/add', component: AddPatient, name: 'add-patient' },
      { path: 'patients/:id', component: ViewPatient, name: 'view-patient' },
      { path: 'patients/:id/next-of-kin', component: NextOfKin, name: 'next-of-kin' },

      { path: 'deliveries', component: Deliveries, name: 'deliveries' },
      { path: 'deliveries/add', component: AddDelivery, name: 'add-delivery' },
      { path: 'deliveries/:id', component: ViewDelivery, name: 'view-delivery' },

      { path: 'staff', component: Staff, name: 'staff' },
      { path: 'staff/add', component: AddStaff, name: 'add-staff' },
      { path: 'staff/:id', component: ViewStaff, name: 'view-staff' },

      { path: ':pathMatch(.*)', redirect: { name: 'dashboard' } },
    ]
  },
  { path: '/:pathMatch(.*)*', redirect: { name: 'login' } },
];