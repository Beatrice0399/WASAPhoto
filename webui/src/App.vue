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
	
	
	<div class="Wasaphoto" style="width: 100%; height: 100%; position: relative">
    <div class="Wasaphoto" style="left: 700px; top: 0px; position: absolute; text-align: center; color: #1E1E1E; font-size: 40px; font-family: Inter; font-style: italic; font-weight: 800; word-wrap: break-word">WASAPhoto</div>
</div>

</template>

<style>

</style>
