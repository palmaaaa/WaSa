<script>
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
				
				this.$router.replace("/home")
			} catch (e) {
				this.errormsg = e.toString();
			}
			// this.loading = false;
		},
	},
	
}
</script>

<template>
	<div class="container-fluid h-100 m-0 p-0 login">

		<div class="row ">
			<div class="col">
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
			</div>
		</div>

		<div class="row h-100 w-100 border-top border-bottom m-0">
			
			<form @submit.prevent="login" class="d-flex flex-column align-items-center justify-content-center p-0">

				<div class="row mt-2 mb-3 border-bottom">
					<div class="col">
						<h2 class="login-title">WASAPhoto Login</h2>
					</div>
				</div>

				<div class="row mt-2 mb-3">
					<div class="col">
						<input type="text" class="form-control" v-model="identifier" placeholder="Your identifier" />
					</div>
				</div>

				<div class="row mt-2 mb-5 ">
					<div class="col ">
						<button class="btn btn-dark"> Register/Login </button>
					</div>
				</div>
			</form>
		</div>
	</div>
</template>

<style>
.login {
    background-image: url("../assets/images/people.png");
    background-color: gainsboro;
    background-repeat: no-repeat;
    height: 100vh;
}

.login-title {
    color: black;
}

</style>
