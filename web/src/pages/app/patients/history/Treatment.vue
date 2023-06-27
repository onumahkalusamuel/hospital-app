<script setup lang="ts">
import HistoryForm from './HistoryForm.vue';
import TextField from '../../../../components/form/TextField.vue'
import SelectField from '../../../../components/form/SelectField.vue'

</script>
<template>
  <history-form>
    <div class="min-w-[250px]"><TextField label="Ward Name" placeholder="00324" name="ward_name" required></TextField></div>
    <div class="min-w-[250px]"><TextField label="Room Number" placeholder="00324" name="room_number" required></TextField></div>
    
    <div class="min-w-[250px]"><TextField label="Middle Name" placeholder="M." name="middlename"></TextField></div>
    <div class="min-w-[250px]"><TextField label="Last Name" placeholder="Doe" name="lastname" required></TextField></div>
    <div class="min-w-[250px]"><SelectField label="Gender" name="sex" :options="[['Male'], ['Female']]" required/></div>
    <div class="min-w-[250px]"><SelectField label="Marital Status" name="marital_status" :options="[[''],['Single'],['Married'],['Divorced'],['Separated']]" /></div>
    <div class="min-w-[250px]"><TextField label="Card No" placeholder="HC0090" name="card_no"></TextField></div>
    <div class="min-w-[250px]"><TextField label="Phone Number" placeholder="08123456789" name="phone" type="tel"></TextField></div>
    <div class="min-w-[250px]"><TextField label="Email Address" placeholder="john.doe@example.com" name="email"></TextField></div>
    <div class="min-w-[250px]"><TextField label="Date of birth" type="date" name="dob" required></TextField></div>
    <div class="min-w-[250px]"><TextField label="Home Address" placeholder="20 John Doe Avenue..." name="address" required></TextField></div>
    <div class="min-w-[250px]"><TextField label="Occupation" placeholder="Engineer" name="occupation"></TextField></div>
    <div class="min-w-[250px]"><TextField label="Appointment Type" placeholder="emergency" name="current_appointment" required></TextField></div>
          
  </history-form>
</template>