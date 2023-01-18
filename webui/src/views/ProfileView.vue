<script>
export default {
	data: function() {
		return {
			errormsg: null,

			userExists: false,
			banStatus: false,

            nickname: "",


			followStatus: false,
			currentIsBanned: false,

			followerCnt: 0,
			followingCnt:0,
			postCnt:0,

			photos: [],
            following: [],
            followers: [],
		}
	},

    watch:{
        currentPath(newid,oldid){
            if (newid !== oldid){
                this.loadInfo()
            }
        },
    },

	computed:{

        currentPath(){
            return this.$route.params.id
        },
        

		sameUser(){
			return this.$route.params.id === localStorage.getItem('token')
		},
	},

	methods: {

        async uploadFile(){
            let fileInput = document.getElementById('fileUploader')

            const file = fileInput.files[0];
            const reader = new FileReader();

            reader.readAsArrayBuffer(file);

            reader.onload = async () => {
                // Post photo: /users/:id/photos
                let response = await this.$axios.post("/users/"+this.$route.params.id+"/photos", reader.result, {
                    headers: {
                    'Content-Type': file.type
                    },
                })
                //console.log(response)
                /*
                this.photos.unshift({
                    owner: response.data.owner,
                    date: response.data.date,
                    photo_id: response.data.photo_id,
                    likes: response.data.likes,
                    comments: response.data.comments,
                })
                */
                this.photos.unshift(response.data)
                this.postCnt += 1
            };
        },

		async followClick(){
            try{
                if (this.followStatus){ 
                    // Delete like: /users/:id/followers/:follower_id
                    await this.$axios.delete("/users/"+this.$route.params.id+"/followers/"+ localStorage.getItem('token'));
                    this.followerCnt -=1
                }else{
                    // Put like: /users/:id/followers/:follower_id
                    await this.$axios.put("/users/"+this.$route.params.id+"/followers/"+ localStorage.getItem('token'));
                    this.followerCnt +=1
                }
                this.followStatus = !this.followStatus
            }catch (e){
                this.errormsg = e.toString();
            }
            
		},

		async banClick(){
            try{
                if (this.banStatus){
                    // Delete ban: /users/:id/banned_users/:banned_id
                    await this.$axios.delete("/users/"+localStorage.getItem('token')+"/banned_users/"+ this.$route.params.id);
                    this.loadInfo()
                }else{
                    // Put ban: /users/:id/banned_users/:banned_id
                    await this.$axios.put("/users/"+localStorage.getItem('token')+"/banned_users/"+ this.$route.params.id);
                    this.followStatus = false
                }
                this.banStatus = !this.banStatus
            }catch(e){
                this.errormsg = e.toString();
            }
		},

		async loadInfo(){
            if (this.$route.params.id === undefined){
                return
            }

			try{
                // Get user profile: /users/:id
				let response = await this.$axios.get("/users/"+this.$route.params.id);

                this.banStatus = false
                this.userExists = true
                this.currentIsBanned = false

                if (response.status === 206){
                    this.banStatus = true
                    return
                }

				if (response.status === 204){
					this.userExists = false
				}
				
                this.nickname = response.data.nickname
				this.followerCnt = response.data.followers != null ? response.data.followers.length : 0
				this.followingCnt = response.data.following != null? response.data.following.length : 0
				this.postCnt = response.data.posts != null ? response.data.posts.length : 0
				this.followStatus = response.data.followers != null ? response.data.followers.find(obj => obj.user_id === localStorage.getItem('token')) : false
                this.photos = response.data.posts != null ? response.data.posts : []
                this.followers = response.data.followers != null ? response.data.followers : []
                this.following = response.data.following != null ? response.data.following : []

			}catch(e){
				this.currentIsBanned = true
			}
		},

        goToSettings(){
            this.$router.push(this.$route.params.id+'/settings')
        },

        removePhotoFromList(photo_id){
			this.photos = this.photos.filter(item => item.photo_id !== photo_id)
		},
	},

	async mounted(){
		await this.loadInfo()
	},

}
</script>

<template>

    <div class="container-fluid" v-if="!currentIsBanned && userExists">

        <div class="row">
            <div class="col-12 d-flex justify-content-center">
                <div class="card w-50 container-fluid">

                    <div class="row">
                        <div class="col">
                            <div class="card-body d-flex justify-content-between align-items-center">
                                <h5 class="card-title p-0 me-auto mt-auto">{{nickname}} @{{this.$route.params.id}}</h5>

                                <button v-if="!sameUser && !banStatus" @click="followClick" class="btn btn-success ms-2">
                                    {{followStatus ? "Unfollow" : "Follow"}}
                                </button>

                                <button v-if="!sameUser" @click="banClick" class="btn btn-danger ms-2">
                                    {{banStatus ? "Unban" : "Ban"}}
                                </button>

                                <button v-else class="my-trnsp-btn ms-2" @click="goToSettings">
                                    <!--Settings  <font-awesome-icon icon="fa-solid fa-gear" />-->
                                    <i class="my-nav-icon-gear fa-solid fa-gear"></i>
                                </button>
                            </div>
                        </div>
                    </div>

                    <div v-if="!banStatus" class="row mt-1 mb-1">
                        <div class="col-4 d-flex justify-content-start">
                            <h6 class="ms-3 p-0 ">Posts: {{postCnt}}</h6>
                        </div>
                    
                        <div class="col-4 d-flex justify-content-center">
                            <h6 class=" p-0 ">Followers: {{followerCnt}}</h6>
                        </div>
                    
                        <div class="col-4 d-flex justify-content-end">
                            <h6 class=" p-0 me-3">Following: {{followingCnt}}</h6>
                        </div>
                    </div>
                </div>
            </div>
        </div>


        <div class="row">

            <div class="container-fluid mt-3">

                <div class="row ">
                    <div class="col-12 d-flex justify-content-center">
                        <h2>Posts</h2>
                        <input id="fileUploader" type="file" class="profile-file-upload" @change="uploadFile" accept=".jpg, .png">
                        <label v-if="sameUser" class="btn my-btn-add-photo ms-2 d-flex align-items-center" for="fileUploader"> Add </label>
                    </div>
                </div>

                <div class="row ">
                    <div class="col-3"></div>
                    <div class="col-6">
                        <hr class="border border-dark">
                    </div>
                    <div class="col-3"></div>
                </div>
            </div>
        </div>

        <div class="row">
            <div class="col">

                <div v-if="!banStatus && postCnt>0">
                    <Photo v-for="(photo,index) in photos" 
                    :key="index" 
                    :owner="this.$route.params.id" 
                    :photo_id="photo.photo_id" 
                    :comments="photo.comments" 
                    :likes="photo.likes" 
                    :upload_date="photo.date" 
                    :isOwner="sameUser" 
                    
                    @removePhoto="removePhotoFromList"
                    />

                </div>
                
                <div v-else class="mt-5 ">
                    <h2 class="d-flex justify-content-center" style="color: white;">No posts yet</h2>
                </div>

            </div>
        </div>

    
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    </div>
    <div v-else class="h-25 ">
        <PageNotFound />
    </div>
    

</template>

<style>
.profile-file-upload{
    display: none;
}

.my-nav-icon-gear{
    color: grey;
}
.my-nav-icon-gear:hover{
    transform: scale(1.3);
}

.my-btn-add-photo{
    background-color: green;
    border-color: grey;
}
.my-btn-add-photo:hover{
    color: white;
    background-color: green;
    border-color: grey;
}

</style>
