import { createApp } from 'vue';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import axios from 'axios';
import Chat from 'vue3-beautiful-chat';
import App from './App.vue';
import router from './router';

const app = createApp(App);
app.config.globalProperties.$http = axios;
axios.defaults.baseURL = 'qanda-bauhinia.app.secoder.net/';
app.use(ElementPlus);
app.use(router);
app.use(Chat);
app.mount('#app');
