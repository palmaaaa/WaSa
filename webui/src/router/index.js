import {
    createRouter,
    createWebHashHistory
} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import SearchView from '../views/SearchView.vue'
import ProfileView from '../views/ProfileView.vue'
import PageNotFoundView from '../views/PageNotFoundView.vue'
import SettingsView from '../views/SettingsView.vue'

const router = createRouter({
    history: createWebHashHistory(
        import.meta.env.BASE_URL),
    routes: [{
            path: '/',
            redirect: '/login'
        },
        {
            path: '/login',
            component: LoginView
        },
        {
            path: '/home',
            component: HomeView
        },
        {
            path: '/search',
            component: SearchView
        },
        {
            path: '/users/:id',
            component: ProfileView

        },
        {
            path: '/users/:id/settings',
            component: SettingsView

        },
        {
            path: "/:catchAll(.*)",
            component: PageNotFoundView
        },
    ]
})

export default router