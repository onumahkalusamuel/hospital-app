<script setup lang="ts">
import HistoryForm from './HistoryForm.vue';
import TextField from '../../../../components/form/TextField.vue'

</script>
<template>
  <history-form>
    <div class="min-w-[250px]"><TextField label="Subject" placeholder="subject" name="subject" required></TextField></div>      
  </history-form>
</template>