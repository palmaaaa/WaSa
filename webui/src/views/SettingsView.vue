<script>
export default {
	data: function () {
		return {
			// errormsg: null,
			nickname: "",
		}
	},

	methods:{
		async modifyNickname(){
			console.log("aooo")
			try{
				// Nickname put: /users/:id
				let resp = await this.$axios.put("/users/"+this.$route.params.id,{
					nickname: this.nickname,
				})
				console.log(resp)
				this.nickname=""
			}catch (e){
				//console.log(e.toString())
			}
		},
	},

}
// To add likes and comments put modal scrollable from bootstrap
</script>

<template>
	<div class="container-fluid">
		<div class="row">
			<div class="col d-flex justify-content-center">
				<h1>Impostazioni di {{ this.$route.params.id }}</h1>
			</div>
		</div>

		<div class="row mt-2">
			<div class="col d-flex justify-content-center">
				<div class="input-group mb-3 w-25">
					<input
						type="text"
						class="form-control w-25"
						placeholder="Your new nickname..."
						maxlength="16"
						minlength="3"
						v-model="this.nickname"
					/>
					<div class="input-group-append">
						<button class="btn btn-outline-secondary" 
						@click="this.modifyNickname"
						:disabled="this.nickname == null || this.nickname.length >16 || this.nickname.length <3">
						Modify</button>
					</div>
				</div>
			</div>
		</div>

		<!--<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>-->
	</div>
</template>

<style>
</style>
