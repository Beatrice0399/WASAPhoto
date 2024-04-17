import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import SessionView from '../views/SessionView.vue'
import SearchView from '../views/SearchView.vue'
import ProfileView from '../views/ProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			redirect: '/session'
		},
		{
			path: '/session', 
			component: SessionView
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
			path: '/users/:uid',
			component: ProfileView
		},
	]
})

export default router
