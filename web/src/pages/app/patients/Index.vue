<script lang="ts" >
import Breadcrumbs, { BreadcrumbItem } from '../../../components/Breadcrumbs.vue';
import ActionButton from '../../../components/ActionButton.vue';
import PageHeader from '../../../components/PageHeader.vue';
import apiRequest from '../../../services/http/api-requests';
import { Patient } from '../../../interfaces/patient'

export default {
    name: "Patients",
    data: () => ({
      breadcrumbs: [
        { title: "Dashboard", link: { name: "dashboard" } },
        { title: "Patients", current: true },
      ] as BreadcrumbItem[],
      patients: [] as Patient[],
    }),
    components: { Breadcrumbs, PageHeader, ActionButton },
    methods: {
      check() {
        console.log(444);
      }
    },
    async mounted () {
      this.patients = await apiRequest.get('patients');  
    }
}; </script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="Patients" subtitle="Manage patients" icon-src="/icons/people-fill.svg"></PageHeader>
    <div style="padding: 0 15px; display: flex; justify-content: space-between; border-top:1px solid #333">
      <div>
        <ActionButton icon-src="/icons/trash.svg" label="Delete permanently"></ActionButton>
        <ActionButton icon-src="/icons/recycle.svg" label="Refresh"></ActionButton>
        <ActionButton v-on:click="() => $router.push({name: 'dashboard'})" icon-src="/icons/plus.svg" label="Add patient"></ActionButton>
      </div>
      <div>
        <fluent-text-field appearance="outline" placeholder="Search">
          <i slot="start" style="display:flex"><img src="/icons/search.svg" alt="menu"/></i>
        </fluent-text-field>
      </div>
    </div>
    <fluent-divider class="ml-4 mr-4"></fluent-divider>
    <div class="page-scroll-area">
      <div v-if="!patients">
        <div class="text-center mt-4 mb-4 pt-4">No records found</div>
        <fluent-divider></fluent-divider>
      </div>
      <div v-else class="table">
        <div class="table-row table-header">
          <div class="cell-data cell-checkbox"><fluent-checkbox></fluent-checkbox></div>
          <div class="cell-data cell-size-1">Name</div>
          <div class="cell-data">Card No</div>
          <div class="cell-data">Sex</div>
          <div class="cell-data">Phone</div>
          <div class="cell-data">Current Appointment</div>
        </div>
        <div class="table-body">
          <div class="table-row" v-for="patient in patients" :key="patient.id">
            <div class="cell-data cell-checkbox"><fluent-checkbox></fluent-checkbox></div>
            <div class="cell-data cell-size-1">
              <router-link :to="{name: 'patient', params: { id: patient.id }}">
                {{ `${patient.firstname} ${patient.lastname}` }}
              </router-link>
            </div>
            <div class="cell-data">{{ patient.card_no }}</div>
            <div class="cell-data">{{ patient.sex }}</div>
            <div class="cell-data">{{ patient.phone }}</div>
            <div class="cell-data">{{ patient.current_appointment }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>