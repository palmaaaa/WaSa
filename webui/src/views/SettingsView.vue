<script>
export default {
	data: function () {
		return {
			errormsg: null,
			nickname: "",
		}
	},

	methods:{
		async modifyNickname(){
			try{
				// Nickname put: /users/:id
				let resp = await this.$axios.put("/users/"+this.$route.params.id,{
					nickname: this.nickname,
				})

				this.nickname=""
			}catch (e){
				this.errormsg = e.toString();
			}
		},
	},

}
</script>

<template>
	<div class="container-fluid">
		<div class="row">
			<div class="col d-flex justify-content-center mb-2">
				<h1>{{ this.$route.params.id }}'s Settings</h1>
			</div>
		</div>

		<div class="row ">
			<div class="col-12 d-flex justify-content-center">
				<p class="me-1" style="color: var(--color-red-danger);">[Disclaimer] </p> <p> A user has this structure: </p> <p class="ms-1 me-1" style="color: green;">nickname</p> <p> @identifier. </p>
			</div>
			<div class="col-12 d-flex justify-content-center">
				<p>It's only possible to modify the part before the @</p> (the <p class="ms-1 me-1" style="color: green;">nickname</p>) <p>and not the one after (the identifier of the user) </p>
			</div>
			<div class="col-12 d-flex justify-content-center">
				<p>Username has been intepreted as a nickname (they're the same thing).  </p>
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
						v-model="nickname"
					/>
					<div class="input-group-append">
						<button class="btn btn-outline-secondary" 
						@click="modifyNickname"
						:disabled="nickname === null || nickname.length >16 || nickname.length <3 || nickname.trim().length===0">
						Modify</button>
					</div>
				</div>
			</div>
		</div>

		<div class="row" >
			<div v-if="nickname.trim().length>0" class="col d-flex justify-content-center">
				Preview: {{nickname}} @{{ this.$route.params.id }}
			</div>
		</div>

		<div class="row">
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		</div>
	</div>
	
</template>

<style>
</style>
