<script>
export default {	
	data(){
		return{
			commentValue:"",
		}
	},
	props:['modal_id','comments_list','photo_owner','photo_id'],

	methods: {
		async addComment(){
			try{
				// Comment post: /users/:id/photos/:photo_id/comments
				let response = await this.$axios.post("/users/"+ this.photo_owner +"/photos/"+this.photo_id+"/comments",{
					user_id: localStorage.getItem('token'),
					comment: this.commentValue
				},{
					headers:{
						'Content-Type': 'application/json'
					}
				})

				this.comments_list.push({
					IdComment: response.data.comment_id, 
					IdPhoto: this.photo_id, 
					IdUser: localStorage.getItem('token'), 
					Comment: this.commentValue}
				)

				this.commentValue = ""
			}catch(e){
				console.log(e.toString())
			}
		},

		eliminateCommentToParent(value){
			this.$emit('eliminateComment',value)
		}
	},
}
</script>

<template>
    <div class="modal fade my-modal-disp-none" :id="this.modal_id" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-dialog modal-dialog-scrollable ">
            <div class="modal-content">

                <div class="modal-header">
                    <h1 class="modal-title fs-5" :id="this.modal_id">Comments</h1>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>

                <div class="modal-body">
                    <PhotoComment v-for="(comm,index) in comments_list" 
					:key="index" 
					:author="comm.IdUser" 
					:comment_id="comm.IdComment"
					:photo_id="comm.IdPhoto"
					:content="comm.Comment"
					:photo_owner="this.photo_owner"
					:comments_list="this.comments_list"

					@eliminateComment="this.eliminateCommentToParent"
					/>

                </div>
                <div class="modal-footer d-flex justify-content-center w-100">
                    <div class="row w-100 ">
                        <div class="col-10">
                            <div class="mb-3 me-auto">
                                
                                <textarea class="form-control" id="exampleFormControlTextarea1" 
								placeholder="Add a comment..." rows="1" maxLength="30" v-model="this.commentValue"></textarea>
                            </div>
                        </div>

                        <div class="col-2 d-flex align-items-center">
                            <button type="button" class="btn btn-primary" @click.prevent="this.addComment">Send</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

</template>

<style> 
.my-modal-disp-none{
	display: none;
}
</style>
