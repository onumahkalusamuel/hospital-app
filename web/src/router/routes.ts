import { RouteRecordRaw } from 'vue-router';
import { appGuard } from '@/guards/app-guard';

import MainLayout from '@/layouts/MainLayout.vue';
import AppLayout from '@/layouts/AppLayout.vue';

import Activate from '@/pages/Activate.vue';
import Login from '@/pages/Login.vue';
import CreateAdmin from '@/pages/CreateAdmin.vue';

import HospitalSetup from '@/pages/HospitalSetup.vue';
import Dashboard from '@/pages/app/Dashboard.vue';
import Settings from '@/pages/app/Settings.vue';

import Patients from '@/pages/app/patients/Index.vue';
import AddPatient from '@/pages/app/patients/Add.vue';
import ViewPatient from '@/pages/app/patients/View.vue';
import NextOfKin from '@/pages/app/patients/NextOfKin.vue';
import PatientHistory from '@/pages/app/patients/HistoryIndex.vue';
import PatientHistoryAdd from '@/pages/app/patients/HistoryAdd.vue';
import PatientHistoryPrintView from '@/pages/app/patients/HistoryPrintView.vue';
import PatientHistoryView from '@/pages/app/patients/HistoryView.vue';

import Deliveries from '@/pages/app/deliveries/Index.vue';
import AddDelivery from '@/pages/app/deliveries/Add.vue';
import ViewDelivery from '@/pages/app/deliveries/View.vue';

import Staff from '@/pages/app/staff/Index.vue';
import AddStaff from '@/pages/app/staff/Add.vue';
import ViewStaff from '@/pages/app/staff/View.vue';

import Invoices from '@/pages/app/invoices/Index.vue';
import AddInvoice from '@/pages/app/invoices/Add.vue';
import ViewInvoice from '@/pages/app/invoices/View.vue';
import InvoicePayment from '@/pages/app/invoices/PaymentIndex.vue';
import InvoicePaymentAdd from '@/pages/app/invoices/PaymentAdd.vue';
import InvoicePaymentView from '@/pages/app/invoices/PaymentView.vue';


import Billings from '@/pages/app/staff/Index.vue';
import Reports from '@/pages/app/staff/Index.vue';


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
  { path: '/app/patients/:id/patient-history/print', component: PatientHistoryPrintView, name: 'patient-history-print-view' },
  {
    path: '/app/', component: AppLayout, beforeEnter: appGuard, children: [
      { path: 'dashboard', component: Dashboard, name: 'dashboard' },
      { path: 'settings', component: Settings, name: 'settings' },
      { path: 'patients', component: Patients, name: 'patients' },
      { path: 'patients/add', component: AddPatient, name: 'add-patient' },
      { path: 'patients/:id', component: ViewPatient, name: 'view-patient' },
      { path: 'patients/:id/next-of-kin', component: NextOfKin, name: 'next-of-kin' },
      { path: 'patients/:id/patient-history', component: PatientHistory, name: 'patient-history' },
      { path: 'patients/:id/patient-history/add', component: PatientHistoryAdd, name: 'patient-history-add' },
      { path: 'patients/:id/patient-history/:hid', component: PatientHistoryView, name: 'patient-history-view' },

      { path: 'deliveries', component: Deliveries, name: 'deliveries' },
      { path: 'deliveries/add', component: AddDelivery, name: 'add-delivery' },
      { path: 'deliveries/:id', component: ViewDelivery, name: 'view-delivery' },

      { path: 'staff', component: Staff, name: 'staff' },
      { path: 'staff/add', component: AddStaff, name: 'add-staff' },
      { path: 'staff/:id', component: ViewStaff, name: 'view-staff' },

      { path: 'invoices', component: Invoices, name: 'invoices' },
      { path: 'invoices/add', component: AddInvoice, name: 'add-invoice' },
      { path: 'invoices/:id', component: ViewInvoice, name: 'view-invoice' },
      { path: 'invoices/:id/payments', component: InvoicePayment, name: 'payments' },
      { path: 'invoices/:id/payments/add', component: InvoicePaymentAdd, name: 'payments-add' },
      { path: 'invoices/:id/payments/:hid', component: InvoicePaymentView, name: 'payments-view' },
      
      { path: 'billings', component: Billings, name: 'billings' },

      { path: 'reports', component: Reports, name: 'reports' },

      { path: ':pathMatch(.*)', redirect: { name: 'dashboard' } },
    ]
  },
  { path: '/:pathMatch(.*)*', redirect: { name: 'login' } },
];