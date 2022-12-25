<script>
export default {
	data(){
		return{
			photoURL: "",
			liked: false,
			comment: "",
		}
	},

	props: ['owner','likes','comments',"upload_date","photo_id","isOwner"], 

	methods:{
		async loadPhoto(){
			/*
			//console.log("aoooo, ",this.photo_id)
			let response = await this.$axios.get("/users/"+this.owner+"/photos/"+this.photo_id)
			this.photoURL = response.config.baseURL+response.config.url // +"/"+
			//console.log("aeeeeee, ",this.photoURL)
			*/
			//console.log(this.ret.__API_URL__)
			this.photoURL = "http://localhost:3000"+"/users/"+this.owner+"/photos/"+this.photo_id // +"/"+this.owner+"/photos/"+this.photo_id 
		},

		async deletePhoto(){
			try{
				// Delete photo: /users/:id/photos/:photo_id
				await this.$axios.delete("/users/"+this.owner+"/photos/"+this.photo_id)
				location.reload()
			}catch(e){
				//
			}
		},

		likeClick: function(){
			console.log("clicked on likes")
		},

		commentClick: function(){
			console.log("clicked on comments")
		},

		photoOwnerClick: function(){
			console.log("clicked on photo owner")
			this.$router.replace("/users/"+this.owner)
		},

		refreshComments(){
			this.$forceUpdate()
		},

		toggleLike() {
			if(this.isOwner){ 
				return
			}

			const bearer = localStorage.getItem('token')

			try{
				if (!this.liked){
					// Put like: /users/:id/photos/:photo_id/likes/:like_id"
					this.$axios.put("/users/"+ this.owner +"/photos/"+this.photo_id+"/likes/"+ bearer)
					this.likes.push('temp')
				}else{
					// Delete like: /users/:id/photos/:photo_id/likes/:like_id"
					this.$axios.delete("/users/"+ this.owner  +"/photos/"+this.photo_id+"/likes/"+ bearer)
					this.likes.pop()
				}

				this.liked = !this.liked;
			}catch(e){
				//
			}
      		
    	},
	},
	

	async mounted(){
		this.loadPhoto()
		if (this.likes != null){
			this.liked = this.likes.some(obj => obj.IdUser === localStorage.getItem('token'))
		}
	},

}
</script>

<template>
	<div class="container-fluid mt-3 mb-5 ">

        <LikeModal :modal_id="'like_modal'+this.photo_id" />

        <CommentModal :modal_id="'comment_modal'+this.photo_id" 
		:comments_list="this.comments" 
		:photo_owner="this.owner" 
		:photo_id="this.photo_id"
		/>

        <div class="d-flex flex-row justify-content-center">

            <div class="card my-card">
                <div class="d-flex justify-content-end">

                    <button v-if="this.isOwner" class="btn btn-primary" @click="this.deletePhoto">Delete</button>

                </div>
                <div class="d-flex justify-content-center photo-background-color">
                    <img :src="this.photoURL" class="card-img-top img-fluid">
                </div>

                <div class="card-body">

                    <div class="container">

                        <div class="d-flex flex-row justify-content-end align-items-center mb-2">

							<button class="btn my-trnsp-btn m-0 p-1 me-auto" @click="this.photoOwnerClick">
                            	<i> From {{owner}}</i>
							</button>

                            <button class="btn my-trnsp-btn m-0 p-1 d-flex justify-content-center align-items-center">
                                <i @click="toggleLike" :class="'me-1 my-heart-color w-100 h-100 fa '+(this.liked ? 'fa-heart' : 'fa-heart-o') "></i>
                                <i data-bs-toggle="modal" :data-bs-target="'#like_modal'+this.photo_id" class="my-comment-color ">
                                    {{likes != null ? likes.length : 0}}
                                </i>
                            </button>

                            <button class="btn my-trnsp-btn m-0 p-1  d-flex justify-content-center align-items-center" 
							data-bs-toggle="modal" :data-bs-target="'#comment_modal'+this.photo_id">

                                <i class="my-comment-color fa-regular fa-comment me-1" @click="this.commentClick"></i>
                                <i class="my-comment-color-2"> {{comments != null ? comments.length : 0}}</i>

                            </button>
                        </div>

                        <div class="d-flex flex-row justify-content-start align-items-center ">
                            <p> Uploaded on {{upload_date}}</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
.photo-background-color{
	background-color: grey;
}

.my-card{
	width: 27rem;
	border-color: black;
	border-width: thin;
}

.my-trnsp-btn{
	border: none;
}
.my-trnsp-btn:hover{
	border: none;
}

.my-heart-color{
	color: grey;
}
.my-heart-color:hover{
	color: red;
}

.my-comment-color {
	color: grey;
}
.my-comment-color:hover{
	color: black;
}

.my-comment-color-2{
	color:grey
}
</style>
