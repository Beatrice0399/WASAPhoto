<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	data() {
		return{
			logged: false,
			searchValue: "",
		}
	},
	methods: {
		logout(value){
			this.logged = value
			this.$router.replace("/session")
		},
		updateLogged(newLogged) {
			this.logged = newLogged
		},
		updateView(newRoute) {
			this.$router.replace(newRoute)
		},
		search(queryParam) {
			this.searchValue = queryParam
			this.$router.replace("/search")
		},
	},
	mounted() {
		if (!localStorage.getItem('token')){
			this.$router.replace("/session")
		} else {
			this.logged = true
		}
	},
}
</script>

<template>
	
	<div class="container-fluid app">
		<div class="row">
			<div class="col p-0">
				<main >
				
					<NavBar v-if="logged" 
					@logoutNavBar="logout" 
					@requestUpdateView="updateView"
					@searchNavBar="search"/>
					
					<RouterView 
					@updatedLoggedChild="updateLogged" 
					@requestUpdateView="updateView"
					
					:searchValue="searchValue"/>
				</main>
			</div>
		</div>
	</div>
	


</template>

<style>
.app {
background: linear-gradient(to top, #abc9dc, #ffffff);
  min-height: 100vh;
}

</style>
