<script lang="ts" >
import Breadcrumbs from '../../components/Breadcrumbs.vue';
import PageHeader from '../../components/PageHeader.vue';
import ActionButton from '../../components/ActionButton.vue';

export default {
    name: "Dashboard",
    data: () => ({
      links: [
        { label: 'Patients', iconSrc: '/icons/people-fill.svg', url: { name: 'patients' } },
        { label: 'Deliveries', iconSrc: '/icons/basket.svg', url: { name: 'patients' } },
        { label: 'Billings', iconSrc: '/icons/bank.svg',  url: { name: 'patients' } },
        { label: 'Reports', iconSrc: '/icons/receipt.svg', url: { name: 'patients' } },
      ]

    }),
    components: { Breadcrumbs, ActionButton, PageHeader }
}; </script>;

<template>
  <div class="dashboard-container">
    <div class="dashboard-title">Welcome, Dr. Hilda Baci</div>
    <fluent-divider></fluent-divider>
    <div class="dashboard-body">
      <router-link v-for="link, i in links" :key="i" class="dashboard-body-item" :to="link.url">
        <div class="dashboard-body-icon">
          <img :src="link.iconSrc" />
        </div>
        <div class="dashboard-body-label">{{ link.label }}</div>
      </router-link>
    </div>
   </div>
</template>
<style scoped>
.dashboard-container {
  padding: 15px;
}
.dashboard-title {
  font-size: 2rem;
  padding: 80px 0 20px 0;
}
.dashboard-body {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
}
.dashboard-body-item {
  width: 212px;
  height: 259px;
  margin-top: 47px;
  border: 2px solid #0078d4;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: #0078d4;
  text-decoration: none;
}
.dashboard-body-item:hover {
  background-color: azure;
}
.dashboard-body-icon {
  width: 156px;
  height: 156px;
  filter: invert(38%) sepia(81%) saturate(7137%) hue-rotate(193deg) brightness(101%) contrast(102%);
}
.dashboard-body-label {
  font-size: 1.5rem;
  margin-top: 10px;
}
</style>