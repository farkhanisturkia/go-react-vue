//import createApp from Vue
import { createApp } from 'vue';

//import component App
import App from './App.vue';

//import VueQueryPlugin
import { VueQueryPlugin } from '@tanstack/vue-query'

//import config router
import routes from './routes';

//create App Vue
const app = createApp(App);

//use VueQueryPlugin
app.use(VueQueryPlugin);

//use routes on vue
app.use(routes);

app.mount('#app');
