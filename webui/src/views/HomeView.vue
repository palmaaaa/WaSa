<script>
export default {
	data: function () {
		return {
			// errormsg: null,
			photos: [],
			// id: ""

		}
	},

	methods: {
		async loadStream() {
			try {
				// Home get: "/users/:id/home"
				let response = await this.$axios.get("/users/" + localStorage.getItem('token') + "/home")

				//console.log(response)
				this.photos = response.data
			} catch (e) {

			}
		}
	},

	mounted() {
		// this.id = 
		this.loadStream()
	}

}
// To add likes and comments put modal scrollable from bootstrap
</script>

<template>
	<div>
		<Photo
			v-for="(photo,index) in photos"
			:key="index"
			:owner="photo.Owner"
			:photo_id="photo.PhotoId"
			:comments="photo.Comments"
			:likes="photo.Likes"
			:upload_date="photo.Date"
		/>

		<!--<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>-->
	</div>
</template>

<style>
</style>
