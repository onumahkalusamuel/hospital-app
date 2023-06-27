<script setup lang="ts">
import { ref, onMounted } from 'vue';
import HistoryForm from "./HistoryForm.vue";
import TextField from '../../../../components/form/TextField.vue'
import SelectField from '../../../../components/form/SelectField.vue'
import { Invoice, Patient } from '../../../../interfaces';
import apiRequest from '../../../../services/http/api-requests';
import { useRoute } from 'vue-router';

const invoices = ref([] as Invoice[]);
const invoiceOptions = ref([] as [string, (string|number)?][]);
const route = useRoute();

onMounted(async()=> {
  invoices.value = (await apiRequest.get(`patient/${route.params.id}/patient-invoices`) as Patient).invoices as Invoice[];
  invoices.value.forEach((v) => {
    invoiceOptions.value.push([`${v.details}`, v.id])
  })
});
</script>
<template>
  <history-form>
    <div class="min-w-[250px]"><SelectField label="Invoice" name="invoice_id" :options="invoiceOptions" required/></div>
    <div class="min-w-[250px]"><TextField label="Amount" type="number" name="amount" required></TextField></div>      
  </history-form>
</template>