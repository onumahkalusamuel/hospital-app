<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import apiRequest from '@/services/http/api-requests';
import PopUp from '@/components/PopUp.vue';
import { toasts, popupStore } from '@/stores';
import TextField from '@/components/form/TextField.vue';
import TextArea from '@/components/form/TextArea.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import { Patient, PatientHistoryTypes } from '@/interfaces';

defineProps<{ popupId: string, patient: Patient }>()
const addHistoryRef = ref(null);
const route = useRoute();
const emit = defineEmits(["update:closed"]);
const historyTypes = ref(['General', 'Diagnosis', 'Examination', 'TestResult', 'Treatment'] as PatientHistoryTypes[]);

const addHistory = async () => {
    const formData = new FormData(addHistoryRef.value as never as HTMLFormElement)
    const admit = await apiRequest.postMulti(`patients/${route.params.id}/patient-history`, formData);
    if (admit.id) {
        toasts.addToast({ message: "history admitted successfully.", type: 'success' });
        emit('update:closed');
        popupStore.show = false;
    }
}

const selectedType = ref('General' as PatientHistoryTypes);
const setHistoryType = (type: PatientHistoryTypes) => {
    selectedType.value = type;
}

</script>
<template>
    <pop-up :width="600" :title="`${patient.lastname} ${patient.firstname} - Add History`" :id="popupId">
        <form method="POST" v-on:submit.prevent="addHistory" ref="addHistoryRef" enctype="multipart/form-data">
            <div>
                <div class="font-bold py-2">Select History Type <span class="text-red-500">*</span></div>
                <div class="w-full flex flex-wrap mb-3 gap-3">
                    <div v-for="(type, i) in historyTypes" :key="i" class="cursor-pointer p-2 px-2 border-[1px] rounded"
                        :class="selectedType == type ? 'bg-blue-600 border-blue-600 text-white' : 'border-gray-400'"
                        @click="() => setHistoryType(type)"> {{ type }} </div>
                </div>
                <input type="hidden" v-model="selectedType" name="type" />
            </div>

            <div class="flex">
                <div class="w-full"><TextArea label="Note" placeholder="Note" name="note" rows="5"></TextArea></div>
            </div>
            <div class="w-full my-5">
                <TextField label="Document (.png, .jpg, .jpeg)" type="file" name="document" accept=".png, .jpg, .jpeg">
                </TextField>
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