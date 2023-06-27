<script setup lang="ts">
import HistoryForm from './HistoryForm.vue';
import TextField from '../../../../components/form/TextField.vue'

</script>
<template>
  <history-form>
    <div class="min-w-[250px]"><TextField label="Diagnosis" name="image" class="w-full mb-2 font-xl" ref="logo" type="file"/></div>
  </history-form>
</template>