<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import apiRequest from '@/services/http/api-requests';
import PopUp from '@/components/PopUp.vue';
import { toasts, popupStore, hospital } from '@/stores';
import TextField from '@/components/form/TextField.vue';
import TextArea from '@/components/form/TextArea.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import { Patient, PatientHistory } from '@/interfaces';
import SelectField from '@/components/form/SelectField.vue';

const props = defineProps<{ popupId: string, patient: Patient, history: PatientHistory }>()
const updateHistoryRef = ref(null);
const route = useRoute();
const emit = defineEmits(["update:closed"]);

const updateHistory = async () => {
    const formData = new FormData(updateHistoryRef.value as never as HTMLFormElement)
    const update = await apiRequest.putMulti(`patients/${route.params.id}/patient-history/${props.history.id}`, formData);
    console.log(update)
    if (update.message) {
        toasts.addToast({ message: "history updated successfully", type: 'success' });
        emit('update:closed');
        popupStore.show = false;
    }
}

</script>
<template>
    <pop-up :width="600" :title="`${patient.lastname} ${patient.firstname} - Update History`" :id="popupId">
        <form method="POST" v-on:submit.prevent="updateHistory" ref="updateHistoryRef" enctype="multipart/form-data">
            <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                <div class="title">History Type:</div>
                <div class="">{{ history.type }}</div>
                <input type="hidden" :value="history.type" name="type" />
            </div>
            <div class="flex gap-2" v-if="history.type == 'Admission'">
                <div class="w-full">
                    <TextField label="Ward" placeholder="Ward 1" name="ward_number" :value="history.details.ward_number">
                    </TextField>
                </div>
                <div class="w-full">
                    <TextField label="Room" placeholder="R-045" name="room_number" :value="history.details.room_number">
                    </TextField>
                </div>
            </div>
            <div class="flex gap-2" v-if="history.type == 'Appointment'">
                <div class="w-full">
                    <SelectField label="Appointment Type" name="appointment_type"
                        :options="[[''], ['Emergency'], ['Regular']]" v-model="history.details.appointment_type" required>
                    </SelectField>
                </div>
            </div>

            <div class="w-full">
                <TextArea label="Note" placeholder="Note" name="note" rows="2" :value="history.details.note"></TextArea>
            </div>
            <div class="w-full my-5">
                <TextField label="Document (.png, .jpg, .jpeg)" type="file" name="document" accept=".png, .jpg, .jpeg">
                </TextField>
            </div>
            <div class="mb-2 border-[1px] bg-blue-100 p-2 rounded pointer-cusrsor" v-if="history.document.length">
                <div class="title">Previous Document:</div>
                <a :href="`${hospital.asset_base_url}/files/images/${history.document}`" target="_blank">
                    <img class="max-h-[150px]" :src="`${hospital.asset_base_url}/files/images/${history.document}`"
                        alt="Supporting Document" />
                </a>
            </div>
            <div class="flex gap-3 mt-5">
                <div>
                    <PrimaryButton type="submit">Submit</PrimaryButton>
                </div>
                <div>
                    <SecondaryButton type="button" @click.prevent="() => popupStore.show = false">Cancel</SecondaryButton>
                </div>
            </div>
        </form>
    </pop-up>
</template>

<style scoped>
.title {
    color: #0078d4;
    border-radius: 5px;
    font-weight: bold;
    margin-right: 5px;
}
</style>