import {
    createApp,
    reactive
} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import Navbar from './components/Navbar.vue'
import Photo from './components/Photo.vue'
import UserMiniCard from './components/UserMiniCard.vue'
import PageNotFound from './components/PageNotFound.vue'
import LikeModal from './components/LikeModal.vue'
import CommentModal from './components/CommentModal.vue'
import PhotoComment from './components/PhotoComment.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;

app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("Navbar", Navbar);
app.component("Photo", Photo);
app.component("UserMiniCard", UserMiniCard);
app.component("PageNotFound", PageNotFound);
app.component("LikeModal", LikeModal);
app.component("CommentModal", CommentModal);
app.component("PhotoComment", PhotoComment);
app.use(router)
app.mount('#app')