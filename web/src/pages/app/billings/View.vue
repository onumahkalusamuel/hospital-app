<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import dayjs from 'dayjs'
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import { onMounted, ref, watch } from 'vue';
import { Invoice } from '@/interfaces';
import { useRoute } from 'vue-router';
import { PrinterIcon, StarIcon, PlusCircleIcon, TrashIcon } from '@heroicons/vue/24/solid';
import ActionButton from '@/components/ActionButton.vue';
import AddPaymentPopup from './popups/AddPaymentPopup.vue';
import { hospital, popupStore, toasts, user } from '@/stores';

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Billings",  link: { name: "billings" } },
  { title: "View Invoice", current: true },
] as BreadcrumbItem[]);

const route = useRoute();
const invoice = ref({ } as Invoice);
const printInvoiceArea = ref();
const printReceiptArea = ref();
const addPaymentPopupId = ref('addpayment');

const printInvoice = () => {
  toasts.clearAll();
  (printInvoiceArea.value as HTMLElement).classList.remove("hidden");
  window.print();
  (printInvoiceArea.value as HTMLElement).classList.add("hidden");
}

const printReceipt = () => {
  toasts.clearAll();
  (printReceiptArea.value as HTMLElement).classList.remove("hidden");
  window.print();
  (printReceiptArea.value as HTMLElement).classList.add("hidden");
}

const fetchInvoice = async () => {
  invoice.value = await apiRequest.get('invoices/' + route.params.id) as Invoice;
}

const showAddPaymentPopup = () => { popupStore.id = addPaymentPopupId.value; popupStore.show = true;}

const deletePayment = async (id: string)  => {
  const del = await apiRequest.deleteRecord(`invoices/${route.params.id}/payments/${id}`);
  if(del && del.message) toasts.addToast({message: del.message, type: 'success'});
  await fetchInvoice();
}

// life-cycles
onMounted(async() => await fetchInvoice())
watch(() => popupStore.show, async (newValue) => {
  if (newValue === false) await fetchInvoice();
})
</script>;

<template>
  <div class="absolute top-0 left-0 bg-white h-screen w-screen mx-auto hidden" ref="printInvoiceArea">
    <div class="p-[15px]">
      <div class="flex gap-2">
        <div class="flex gap-3 items-center w-[50%] border-[1px] p-2 border-stone-600">
          <img :src="hospital.hospital_logo" class="max-w-[50px]" alt="logo">
          <div>
            <div class="text-[12px] font-bold"  style="text-transform:uppercase">{{ hospital.hospital_name }}</div>
            <div>{{ hospital.hospital_address }}</div>
            <div>{{ hospital.hospital_phone }}</div>
          </div>
        </div>
        <div class="border-[1px] p-2 border-stone-600 w-[50%]">
          <div class="mb-2 font-bold">Bill To:</div>
          <div>{{ invoice.name }}</div>
          <div>{{invoice.billing_address}}</div>
        </div>
      </div>
      <div class="font-bold flex justify-between pt-4">
        <div>Invoice No: {{ invoice.invoice_number }}</div>
        <div>Invoice Date: {{ dayjs(invoice.invoice_date).format('DD-MM-YYYY') }}</div>
        <div>Due Date: {{ dayjs(invoice.due_date).format('DD-MM-YYYY') }}</div>
      </div>
      <table class="w-full my-5">
        <thead class="border-t-[1px] border-stone-400 border-b-[1px] h-[40px]">
          <tr>
            <th class="text-center w-[50px]">S/N</th>
            <th class="text-left">Description</th>
            <th class="text-center w-[100px]">Quantity</th>
            <th class="text-center w-[100px]">Unit</th>
            <th class="text-right w-[100px]">Price</th>
            <th class="text-right w-[100px]">Amount</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="d, key in invoice.details" :key="key" class="border-t-[1px] border-stone-300 h-[30px]">
            <td class="text-center">{{ key + 1 }}.</td>
            <td>
              <div class="w-full">{{ d.description }}</div>
            </td>
            <td><div class="w-[120px] text-center">{{ d.qty.toLocaleString("en-US") }}</div></td>
            <td><div class="w-[120px] text-center">{{ d.unit }}</div></td>
            <td><div class="text-right w-[120px]">{{ d.price.toLocaleString("en-US") }}</div></td>
            <td class="text-right">{{ d.amount?.toLocaleString("en-US") }}</td>
          </tr>
        </tbody>
        <tfoot class="border-t-[1px] border-stone-400 border-b-[1px]">
          <tr>
            <th class="text-right" colspan="5">SUB TOTAL</th>
            <th colspan="2" class="font-bold text-right">{{ invoice.sub_total?.toLocaleString("en-US") }}</th>
          </tr>
          <tr>
            <th class="text-right" colspan="5">DISCOUNT</th>
            <th colspan="2" class="font-bold text-right">{{ (invoice.discount || 0).toLocaleString("en-US") }}</th>
          </tr>
          <tr>
            <th class="text-right" colspan="5">GRAND TOTAL</th>
            <th colspan="2" class="font-bold text-right">{{ invoice.amount?.toLocaleString("en-US") }}</th>
          </tr>
        </tfoot>
      </table>
      <br/>
      <div class="text-[8px]">Printed by: {{ user.username }};<br/>Date: {{ new Date().toLocaleDateString('en-us', { weekday:"short", year:"numeric", month:"short", day:"numeric", hour: "2-digit", minute: "2-digit", second: "2-digit"})  }}</div>
    </div>  
  </div>

  <div class="absolute top-0 left-0 bg-white h-screen w-screen mx-auto hidden" ref="printReceiptArea">
    <div class="w-[220px] bg-white text-stone-950 text-[9px] p-[10px] pt-1">
        <div class="flex gap-2 items-center justify-center border-[1px] border-stone-800 p-1">
          <img :src="hospital.hospital_logo" class="max-w-[40px]" alt="logo">
          <div>
            <div class="text-[12px]">{{ hospital.hospital_name }}</div>
            <div>{{ hospital.hospital_address }}</div>
            <div>{{ hospital.hospital_phone }}</div>
          </div>
        </div>
        <div class="text-[12px] text-center py-1">INVOICE {{ invoice.invoice_number }}</div>
        <table>
          <thead class="border-y-[1px] border-black">
            <tr>
              <th class="w-[30px]">QTY</th>
              <th class="w-[100px]">DESCRIPTION</th>
              <th class="w-[40px]">PRICE</th>
              <th class="w-[40px]">AMOUNT</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item, key in invoice.details" :key="key">
              <td class="text-center">{{ item.qty.toLocaleString("en-US") }}</td>
              <td class="w-[1px]">{{ item.description }}</td>
              <td class="w-[1px] text-right">{{ item.price.toLocaleString("en-US") }}</td>
              <td class="w-[1px] text-right">{{ item.amount?.toLocaleString("en-US") }}</td>
            </tr>
          </tbody>
          <tfoot class="border-y-[1px] border-black">
            <tr>
              <td class=""></td>
              <td class="w-[1px]" colspan="2">SUBTOTAL</td>
              <td class="text-right font-bold">{{ (invoice.sub_total || 0).toLocaleString("en-US") }}</td>
            </tr>
            <tr>
              <td class=""></td>
              <td class="w-[1px]" colspan="2">DISCOUNT</td>
              <td class="text-right font-bold">-{{ (invoice.discount || 0).toLocaleString("en-US") }}</td>
            </tr>
            <tr>
              <td class=""></td>
              <td class="w-[1px]" colspan="2">GRAND TOTAL</td>
              <td class="text-right font-bold">{{ (invoice.amount || 0).toLocaleString("en-US") }}</td>
            </tr>
            <tr>
              <td class=""></td>
              <td class="w-[1px]" colspan="2">AMOUNT PAID</td>
              <td class="text-right font-bold">{{ ((invoice.amount || 0) - (invoice.balance || 0)).toLocaleString("en-US") }}</td>
            </tr>
            <tr>
              <td class=""></td>
              <td class="w-[1px]" colspan="2">BALANCE</td>
              <td class="text-right font-bold">{{ (invoice.balance || 0).toLocaleString("en-US") }}</td>
            </tr>
          </tfoot>
        </table>
        <div class="mt-[10px]"> Thank you for your patronage</div>
        <div> Trnx total is VAT inclusive</div>
        <div> NO REFUNDS. NO RETURNS. NO EXCHANGES.</div>
        <br/>
        <div class="text-[8px]">Printed by: {{ user.username }};<br/>Date: {{ new Date().toLocaleDateString('en-us', { weekday:"short", year:"numeric", month:"short", day:"numeric", hour: "2-digit", minute: "2-digit", second: "2-digit"})  }}</div>

    </div>
  </div>

  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="View Invoice" :icon-src="PlusCircleIcon"></PageHeader>
    <div class="px-[15px] flex gap-2 border-t-[1px] border-[#333] py-2">
      <ActionButton dark @click="printInvoice" :icon-src="PrinterIcon">Print Invoice</ActionButton>
      <ActionButton dark @click="printReceipt" :icon-src="StarIcon">Print Receipt</ActionButton>
      <ActionButton v-if="invoice.balance" @click="showAddPaymentPopup" :icon-src="PlusCircleIcon">Add payment</ActionButton>
    </div>

    <hr class="ml-4 mr-4" />
    <div class="p-[15px]">
      <div class="flex justify-between">
        <div class="flex">
          <div class="font-bold pr-3">
            <div class="mb-2">Bill To:</div>
          </div>
          <div class="space-y-2">
            <div class="w-[300px]"> {{ invoice.name }} </div>
            <div>{{invoice.billing_address}}</div>
          </div>
        </div>
        <div class="">
          <div class="space-y-2 flex flex-col">
            <div class="flex items-center">
              <div class="font-bold w-[120px]">Amount:</div>
              <div class="w-[250px] ml-3">{{ invoice.amount?.toLocaleString("en-US") }}</div>
            </div>
            <div class="flex items-center justify-between">
              <div class="font-bold w-[120px]">Paid:</div>
              <div class="w-[250px] ml-3">{{ (((invoice.amount || 0) - (invoice.balance || 0))).toLocaleString("en-US") }}</div>
            </div>
            <div class="flex items-center justify-between">
              <div class="font-bold w-[120px]">Balance:</div>
              <div class="w-[250px] ml-3">{{ invoice.balance?.toLocaleString("en-US") }}</div>
            </div>
          </div>
        </div>
        <div class="">
          <div class="space-y-2 flex flex-col">
            <div class="flex items-center">
              <div class="font-bold w-[120px]">Invoice Number:</div>
              <div class="w-[250px] ml-3">{{ invoice.invoice_number }}</div>
            </div>
            <div class="flex items-center justify-between">
              <div class="font-bold w-[120px]">Invoice Date:</div>
              <div class="w-[250px] ml-3">{{ dayjs(invoice.invoice_date).format('DD-MM-YYYY hh:mm A') }}</div>
            </div>
            <div class="flex items-center justify-between">
              <div class="font-bold w-[120px]">Due Date:</div>
              <div class="w-[250px] ml-3">{{ dayjs(invoice.due_date).format('DD-MM-YYYY hh:mm A') }}</div>
            </div>
          </div>
        </div>
      </div>
      <table class="w-full my-5">
        <thead class="border-t-[1px] border-stone-400 border-b-[1px] h-[40px]">
          <tr>
            <th class="text-center w-[50px]">S/N</th>
            <th class="text-left">Description</th>
            <th class="text-center w-[120px]">Quantity</th>
            <th class="text-center w-[120px]">Unit</th>
            <th class="text-right w-[120px]">Price</th>
            <th class="text-right w-[120px]">Amount</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="d, key in invoice.details" :key="key" class="border-t-[1px] border-stone-300 h-[30px]">
            <td class="text-center">{{ key + 1 }}.</td>
            <td>
              <div class="w-full">{{ d.description }}</div>
            </td>
            <td><div class="w-[120px] text-center">{{ d.qty.toLocaleString("en-US") }}</div></td>
            <td><div class="w-[120px] text-center">{{ d.unit }}</div></td>
            <td><div class="text-right w-[120px]">{{ d.price.toLocaleString("en-US") }}</div></td>
            <td class="text-right">{{ d.amount?.toLocaleString("en-US") }}</td>
          </tr>
        </tbody>
        <tfoot class="border-t-[1px] border-stone-400 border-b-[1px]">
          <tr>
            <th class="text-right" colspan="5">SUB TOTAL</th>
            <th colspan="2" class="font-bold text-right">{{ invoice.sub_total?.toLocaleString("en-US") }}</th>
          </tr>
          <tr>
            <th class="text-right" colspan="5">DISCOUNT</th>
            <th colspan="2" class="font-bold text-right">{{ (invoice.discount || 0).toLocaleString("en-US") }}</th>
          </tr>
          <tr>
            <th class="text-right" colspan="5">GRAND TOTAL</th>
            <th colspan="2" class="font-bold text-right">{{ invoice.amount?.toLocaleString("en-US") }}</th>
          </tr>
        </tfoot>
      </table>
      <div class="text-xl">Payments</div>
      <table class="w-full my-5">
        <thead class="border-t-[1px] border-stone-400 border-b-[1px] h-[40px]">
          <tr>
            <th class="text-center w-[50px]">S/N</th>
            <th class="text-left w-[170px]">Date/Time</th>
            <th class="text-left w-[170px]">Payment Reference</th>
            <th class="text-left">Note</th>
            <th class="text-right w-[140px]">Amount Paid</th>
            <th class="text-right w-[140px]">Balance</th>
            <th class="text-center w-[100px]">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="payment, key in invoice.payments" :key="key" class="border-t-[1px] border-stone-300 h-[30px]">
            <td class="text-center">{{ key + 1 }}.</td>
            <td>{{ dayjs(payment.created_at).format('DD-MM-YYYY hh:mm A') }}</td>
            <td>{{ payment.payment_reference }}</td>
            <td><div class="text-left px-1">{{ payment.note }}</div></td>
            <td><div class="text-right">{{ payment.amount_paid.toLocaleString("en-US") }}</div></td>
            <td><div class="text-right">{{ payment.balance.toLocaleString("en-US") }}</div></td>
            <td class="text-center">
              <ActionButton v-if="user.role == '1'" v-on:click="deletePayment(payment.id)" :icon-src="TrashIcon">Delete</ActionButton>
            </td>
          </tr>
        </tbody>
        <tfoot class="border-t-[1px] border-stone-400 border-b-[1px]">
        </tfoot>
      </table>
      <div class="min-w-[250px] flex space-x-3">
        <div>
          <SecondaryButton type="button" class="w-full" v-on:click.prevent="()=>$router.go(-1)">Go back</SecondaryButton>
        </div>
      </div>
    </div>
    <AddPaymentPopup :popup-id="addPaymentPopupId" />
  </div>
</template>