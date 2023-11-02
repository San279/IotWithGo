import { createApp } from 'vue'
import App from './App.vue'
import { createRouter, createWebHistory } from 'vue-router'
import createStore from '../store/user.js'
import HomePage from './components/HomePage'
import LoginPage from './components/LoginPage'
import HaierOne from './components/HaierOne'
//import {createStore} from 'vuex'


const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: '/home', component: HomePage },
        {path: '/login', component: LoginPage },
        {path: '/ac1', component: HaierOne}
    ]
});

createApp(App).use(createStore).use(router).mount('#app')
//createApp(App).use(store).mount('#app')