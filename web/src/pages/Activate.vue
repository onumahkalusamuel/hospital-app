<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import apiRequest from '../services/http/api-requests';
import { useRoute, useRouter } from 'vue-router';
import TextField from '../components/form/TextField.vue';
import { DocumentDuplicateIcon } from '@heroicons/vue/24/solid'
import PrimaryButton from '../components/form/PrimaryButton.vue';
import { toasts } from '../stores/toasts';

const installation_code = ref('');
const router = useRouter();
const route = useRoute();
const form = ref(null);
const installationCodeField = ref(null);

onMounted(async() => {
  try {
    const req = await apiRequest.get('hospital-details');
    if(req) { router.push({name: 'login'})}

  } catch(e) {
    console.log(e)
  }

  // the process the rest
  installation_code.value = route.params.code as string;
  if(!installation_code.value) {
    const req = await apiRequest.get('installation-code');
    if (req.installation_code) {
      installation_code.value = req.installation_code;
    }
  }
})

const copyInstallationCode = async () => {
  await navigator.clipboard.writeText(installation_code.value as string);
  toasts.addToast({message: 'Installation code copied to clipboard.', type: 'success'});
}

const activate = async () => {
  const formData = new FormData(form.value as never as HTMLFormElement)
  const activate = await apiRequest.post('activate', Object.fromEntries(formData.entries()));
  if(activate.message) {
    toasts.addToast({message: activate.message, type: 'success'});
    router.push({name: 'create-admin'})
  }
}
</script>;

<template>
  <p class="text-xl">Activate HCMS to continue</p>
  <form v-on:submit.prevent="activate" ref="form">
    <div>
      <div class="w-full mb-3 flex items-center">
        <TextField ref="installationCodeField" readonly :value="installation_code" class="w-full">
          <template #append>
            <button class="w-5 h-5" type="button" v-on:click.stop.prevent="copyInstallationCode">
              <DocumentDuplicateIcon class="w-5 h-5" />
            </button>
          </template>
        </TextField>
      </div>
      <TextField name="activation_code" placeholder="Activation Code" class="w-full mb-2" />
      <PrimaryButton class="w-full mb-2" type="submit">Activate</PrimaryButton>
    </div>
  </form>
  <small>Copy the installation code above and forward to <strong>activations@hcms.org</strong> for activation.</small>
</template>
