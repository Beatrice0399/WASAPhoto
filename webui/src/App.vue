<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	data() {
		return{
			logged: false,
		}
	},
	methods: {
		log(value){
			this.logged = value
			this.$router.replace("/session")
		},
	},

	mounted() {
		if (!localStorage.getItem("token")){
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
					<Navbar v-if="logged" 
					@logoutNavbar="logout" 
					@requestUpdateView="updateView"
					@searchNavbar="search"/>

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

</style>
