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
	
	<div class="container-fluid">
		<div class="row">
			<div class="col p-0">
				<main >
				
					<NavBar v-if="logged" 
					@requestUpdateView="updateView"
					@logoutNavbar="logout" 
					@searchNavbar="search"/>
					
					<RouterView 
					@requestUpdateView="updateView"
					@updatedLoggedChild="updateLogged" 
					
					:searchValue="searchValue"/>
				</main>
			</div>
		</div>
	</div>
	


</template>

<style>

</style>
