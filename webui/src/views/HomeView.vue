<script>
export default {
	data: function() {
		return {
			errormsg: null,
			photos: [],
		}
	},
	methods: {
		async loadStream() {
			try {
				this.errormsg = null
				let response = await this.$axios.get("/users/" + localStorage.getItem('token') + "/home")
				
				if (response.data != null) {
					this.photos = response.data
				}
			} catch (e) {
				this.errorMsg = e.toString()
			}
			
		}
	},
	async mounted() {
		await this.loadStream()
	}
}
</script>

<template>
	<div class="container-fluid">
	
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		
		<div class="row">
			<Photo
				v-for="(photo, index) in photos"
				:key="index"
				:phid="photo.phid"
				:uid="photo.user"
				:username="photo.username"
				:path="photo.path"
				:comments="photo.comments"
				:likes="photo.likes"
				:date="photo.date"
				/>
		</div>
		
		
		<div v-if="photos.length === 0" class="row">
			<h1 class="d-flex justify-content-center no-content" style="color: #3F749C;">No content, try to follow somebody!</h1>
		</div>
	</div>

</template>

