<script setup lang="ts">
import dayjs from 'dayjs';
import PopUp from '@/components/PopUp.vue';
import { Patient, PatientHistory } from '@/interfaces';
import { UserIcon } from '@heroicons/vue/24/solid';
import { hospital } from '@/stores';
defineProps<{ popupId: string, history: PatientHistory, patient: Patient }>()
</script>

<template>
    <pop-up title="History Details" :id="popupId">
        <div class="mb-5">
            <div class="flex">
                <div class="w-[80px] h-[80px] mr-2">
                    <component :is="UserIcon" class="w-[80px] h-[80px] text-[#0078d4]" />
                </div>
                <div class="flex flex-col justify-center">
                    <div class=" text-2xl text-[#333]"> {{ patient.title }} {{ patient.lastname }} {{ patient.firstname }}
                        {{ patient.middlename }}</div>
                    <div class="page-header-subtitle">{{ patient.sex }} - {{ patient.card_no }}</div>
                </div>
            </div>
            <div class="grid grid-cols-2 gap-x-2 mt-4">
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Date:</div>
                    <div class="">{{ dayjs(history.date_time).format('DD-MM-YYYY hh:mm A') }}</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Type:</div>
                    <div class="">{{ history.type || "N/A" }}</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded" v-if="history.type == 'Admission'">
                    <div class="title">Ward:</div>
                    <div class="">{{ history.details.ward_number || "N/A" }}</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded" v-if="history.type == 'Admission'">
                    <div class="title">Room:</div>
                    <div class="">{{ history.details.room_number || "N/A" }}</div>
                </div>
            </div>
            <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                <div class="title">Note:</div>
                <div class="">{{ history.details.note || "N/A" }}</div>
            </div>
            <div class="mb-2 border-[1px] bg-blue-100 p-2 rounded pointer-cusrsor" v-if="history.document.length">
                <div class="title">Document:</div>
                <a :href="`${hospital.asset_base_url}/files/images/${history.document}`" target="_blank">
                    <img class="max-h-[250px]" :src="`${hospital.asset_base_url}/files/images/${history.document}`"
                        alt="Supporting Document" />
                </a>
            </div>
        </div>
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