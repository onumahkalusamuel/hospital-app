<script setup lang="ts">
// invoice (invoice id, amount, status),
import { ref, onMounted } from 'vue';
import HistoryForm from "./HistoryForm.vue";
import TextField from '../../../../components/form/TextField.vue'
import SelectField from '../../../../components/form/SelectField.vue'
import PrimaryButton from '../../../../components/form/PrimaryButton.vue';
const counter = ref(1);
const dosages = ref([]);
const frequencies = ref([]);

onMounted(async()=>{

});
</script>
<template>
  <history-form>
    <!-- <div class="min-w-[250px]"><SelectField label="Invoice" name="invoice_id" :options="invoiceOptions" required/></div> -->
    <!-- <div class="min-w-[250px]"><TextField label="Amount" type="number" name="amount" required></TextField></div>       -->

    <div class="w-full flex gap-3 flex-wrap" v-for="x in counter">
      <input type="hidden" :value="x"/>
      <div class="min-w-[250px]"><TextField label="Drug" placeholder="Paracetamol" :name="`items[${x}][drug]`" required></TextField></div>
      <div class="min-w-[250px]"><SelectField label="Dosage" :name="`items[${x}][dosage]`" :options="dosages" required/></div>
      <div class="min-w-[250px]"><SelectField label="Frequency" :name="`items[${x}][freqency]`" :options="frequencies" required/></div>
      <div class="min-w-[250px]"><TextField label="Timeline" placeholder="3 weeks" :name="`items[${x}][timeline]`" required></TextField></div>      
    </div>
    <div class="w-full flex">
      <div>
        <PrimaryButton type="button" @click="() => counter++">Add Invoice Item</PrimaryButton>
      </div>
    </div>

  </history-form>
</template>