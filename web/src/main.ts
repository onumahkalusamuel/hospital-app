import { createApp } from 'vue';
import './style.css';
import App from './App.vue';
import fluentDesigns from './fluent-designs';
import { router } from './router';

// bring in fluent design components
fluentDesigns();

// create app and attach router
createApp(App).use(router).mount('#app');
