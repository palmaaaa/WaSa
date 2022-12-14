<script>
import axios from "axios";
export default {
	data: function() {
		return {
			errormsg: null,
			identifier: ""
		}
	},
	methods: {
		async login() {
			// this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/session",{
					identifier: this.identifier
				});
				
				localStorage.setItem('token',response.data.identifier);
				
				this.$router.replace("/")
			} catch (e) {
				this.errormsg = e.toString();
			}
			// this.loading = false;
		},
	},
	
}
</script>

<template>
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	<div v-else>
		<div class="d-flex justify-content-center align-items-center border-bottom mt-5">
			<h1 class="h2">Login Page</h1>

			<form @submit.prevent="login">
				<input v-model="identifier" placeholder="identifier" />
				
				<button class="btn btn-primary" > Login </button> 
			</form>
		</div>

		
	</div>
</template>

<style>
</style>
