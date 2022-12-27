<script>
export default {
	data: function () {
		return {
			// errormsg: null,
			photos: [],
		}
	},

	methods: {
		async loadStream() {
			try {
				// Home get: "/users/:id/home"
				let response = await this.$axios.get("/users/" + localStorage.getItem('token') + "/home")

				//console.log("homevie",response)
				if (response.data != null){
					this.photos = response.data
				}
				
			} catch (e) {

			}
		}
	},

	async mounted() {
		// this.id = 
		await this.loadStream()

		//console.log("prova home", this.photos)
	}

}
</script>

<template>
	<div class="container-fluid">
		<div class="row">
			<Photo
				v-for="(photo,index) in photos"
				:key="index"
				:owner="photo.owner"
				:photo_id="photo.photo_id"
				:comments="photo.comments"
				:likes="photo.likes"
				:upload_date="photo.date"
			/>
		</div>

		<div v-if="this.photos.length === 0" class="row ">
			<h1 class="d-flex justify-content-center mt-5" style="color: white;">There's no content yet, follow somebody!</h1>
		</div>
		<!--<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>-->
	</div>
</template>

<style>
</style>
