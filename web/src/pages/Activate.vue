<script lang="ts" >
import apiRequest from '../services/http/api-requests';

export default {
  name: 'Activate',
  data: () => ({
    installation_code: ""
  }),

  async mounted() {
    // check installation first
    try {
      const req = await apiRequest.get('hospital-details');
      if(req) { this.$router.push({name: 'login'})}

    } catch(e) {
      console.log(e)
    }
  
    // the process the rest
    this.installation_code = this.$route.params.code as string;
    if(!this.installation_code) {
      const req = await apiRequest.get('installation-code');
      if (req.installation_code) {
        this.installation_code = req.installation_code;
      }
    }
  },
  methods: {
    async copyInstallationCode() {
      const a = this.$refs.installationCode as HTMLFormElement;
        await navigator.clipboard.writeText(a.value as string);
      alert('copied');
    },
    async activate() {
      const formData = new FormData(this.$refs.activate as HTMLFormElement)
      const activate = await apiRequest.post('activate', Object.fromEntries(formData.entries()));
      if(activate.message) {
        alert(activate.message)
        this.$router.push({name: 'create-admin'})
      }
    }
  }
}; </script>;

<template>
    <p style="font-size: 1.2rem;">Activate HCMS to continue</p>
    <form v-on:submit.prevent="activate" ref="activate">
      <div>
        <div class="w-100 mb-3" style="display: flex; align-items: center;">
          <fluent-text-area ref="installationCode" readonly :value="installation_code" class="w-100"></fluent-text-area>
          <fluent-button v-on:click.stop.prevent="copyInstallationCode">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" class="bi bi-clipboard-fill" viewBox="0 0 16 16">
            <path fill-rule="evenodd" d="M10 1.5a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5v-1Zm-5 0A1.5 1.5 0 0 1 6.5 0h3A1.5 1.5 0 0 1 11 1.5v1A1.5 1.5 0 0 1 9.5 4h-3A1.5 1.5 0 0 1 5 2.5v-1Zm-2 0h1v1A2.5 2.5 0 0 0 6.5 5h3A2.5 2.5 0 0 0 12 2.5v-1h1a2 2 0 0 1 2 2V14a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V3.5a2 2 0 0 1 2-2Z"/>
          </svg>
          </fluent-button>
        </div>
        <fluent-text-field name="activation_code" placeholder="Activation Code" class="w-100 mb-2">
        </fluent-text-field>
        <fluent-button class="w-100 mb-2" type="submit" appearance="accent">Activate</fluent-button>
      </div>
    </form>
    <small>Copy the installation code above and forward to <strong>activations@hcms.org</strong> for activation.</small>
</template>
