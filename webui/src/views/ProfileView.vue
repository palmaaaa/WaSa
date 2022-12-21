<script>
export default {
	data: function() {
		return {
			errormsg: null,
			userExists: false,
			followStatus: false,
			banStatus: false,
			currentIsBanned: false,
			followerCnt: 0,
			followingCnt:0,
			postCnt:0,

			photURL: "",
		}
	},

	computed:{

		sameUser(){
			return this.$route.params.id === localStorage.getItem('token')
		},
	},

	methods: {

		async currentCheckBan(){
			try {
				let response = await this.$axios.get("/users/"+this.$route.params.id+"/banned_users",{
						params: {
						id: localStorage.getItem('token'),
					},
				});
				return response.data.IsBanned

			} catch (e) {
				this.errormsg = e.toString();
			}
		},

		async userExistsCheck(){
			try {
				let response = await this.$axios.get("/users/"+this.$route.params.id);
				return response.data.UserExists

			} catch (e) {
				this.errormsg = e.toString();
			}
		},

		async followClick(){
			if (this.followStatus){ 
				await this.$axios.delete("/users/"+this.$route.params.id+"/followers/"+ localStorage.getItem('token'));
			}else{
				await this.$axios.put("/users/"+this.$route.params.id+"/followers/"+ localStorage.getItem('token'));
			}
			this.followStatus = !this.followStatus
		},

		async banClick(){
			if (this.banStatus){
				await this.$axios.delete("/users/"+localStorage.getItem('token')+"/banned_users/"+ this.$route.params.id);
			}else{
				await this.$axios.put("/users/"+localStorage.getItem('token')+"/banned_users/"+ this.$route.params.id);
				this.followStatus = false
			}
			this.banStatus = !this.banStatus
		},

		async test(){
			let response = await this.$axios.get("/users/"+localStorage.getItem('token')+"/photos/brand.png");
			//console.log(response)
			this.photURL = response.request.responseURL // 'http://localhost:3000/users/...../photos/.....'


		},
	},

	async mounted(){
		this.currentIsBanned = await this.currentCheckBan()
		this.userExists = await this.userExistsCheck()
		await this.test()

		
		//console.log(this.userExists,!this.banStatus)
	},
	async updated(){
		//console.log("update profile")
		this.currentIsBanned = await this.currentCheckBan()
		this.userExists = await this.userExistsCheck()
		await this.test()
	}
}
// To add likes and comments put modal scrollable from bootstrap
</script>

<template>
	<div class="container-fluid" v-if="!this.currentIsBanned && this.userExists">

        <div class="row">
            <div class="col-12">

                <div class="d-flex justify-content-center">
                    <div class="card w-50 container-fluid">
                        <div class="row">
                            <div class="col">
                                <div class="card-body d-flex justify-content-between align-items-center">
                                    <h5 class="card-title p-0 me-auto mt-auto">{{this.$route.params.id}}</h5>

                                    <button v-if="!sameUser && !banStatus" @click="followClick" class="btn btn-primary ms-2"> <!--:disabled="sameUser" -->
                                        {{this.followStatus ? "Unfollow" : "Follow"}}
                                    </button>

                                    <button v-if="!sameUser" @click="banClick" class="btn btn-primary ms-2"> <!--:disabled="sameUser" -->
                                        {{this.banStatus ? "Unban" : "Ban"}}
                                    </button>

									<button v-else @click="this.test()" class="btn btn-primary ms-2">
                                        Settings
                                    </button>
                                </div>
                            </div>
                        </div>

                        <div v-if="!banStatus" class="row mt-1 mb-1">
                            <div class="col">
                                <div class="container-fluid d-flex justify-content-between align-items-center">
                                    <div class="row">
                                        <div class="col">
                                            <h6 class=" p-0 ">Posts: 1</h6>
                                        </div>
                                    </div>

                                    <div class="row">
                                        <div class="col">
                                            <h6 class=" p-0 ">Followers: 1</h6>
                                        </div>
                                    </div>

                                    <div class="row">
                                        <div class="col">
                                            <h6 class=" p-0 ">Following: 1</h6>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>

                    </div>
                </div>
            </div>
        </div>
		


        <div class="row">
        	<div v-if="!banStatus">
        		<Photo :path="this.photURL" />
        	</div>
			<div v-else class="mt-5 d-flex justify-content-center">
				No posts yet
			</div>
		</div>

        <!--<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>-->
    </div>
	<div v-else class="h-25 " >
		<PageNotFound/>
	</div>

</template>

<style>
</style>
