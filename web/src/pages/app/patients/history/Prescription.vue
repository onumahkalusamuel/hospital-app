<script setup lang="ts">
import HistoryForm from './HistoryForm.vue';
import TextField from '../../../../components/form/TextField.vue'
import SelectField from '../../../../components/form/SelectField.vue'
import { ref } from 'vue';

const dosages = ref([]);
const frequencies = ref([]);
const counter = ref(1);

</script>
<template>
  <history-form>
    <div class="w-full flex gap-3 flex-wrap" v-for="x in counter">
      <input type="hidden" :value="x"/>
      <div class="min-w-[250px]"><TextField label="Drug" placeholder="Paracetamol" name="drug" required></TextField></div>
      <div class="min-w-[250px]"><SelectField label="Dosage" name="dosage" :options="dosages" required/></div>
      <div class="min-w-[250px]"><SelectField label="Frequency" name="dosage" :options="frequencies" required/></div>
      <div class="min-w-[250px]"><TextField label="Timeline" placeholder="3 weeks" name="timeline" required></TextField></div>      
    </div>
  </history-form>
</template>