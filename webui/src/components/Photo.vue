<script>
export default {
	data(){
		return{

		}
	},
	methods:{
		likeClick: function(){
			console.log("clicked on likes")
		},
		commentClick: function(){
			console.log("clicked on comments")
		},
		photoOwnerClick: function(){
			console.log("clicked on photo owner")
		},
	},
	props: ['owner','num_likes','num_comments','upload_date','path'], 

}
</script>

<template>

<div class="container-fluid mt-5 mb-5 ">
        <div class="d-flex flex-row justify-content-center">

            <div class="card my-card ">
				<div class="d-flex justify-content-center photo-background-color">
                	<img :src="this.path" class="card-img-top  img-size ">
				</div>
                <div class="card-body bruh">
                    
                    <div class="container">

                    	<div class="d-flex flex-row justify-content-start align-items-center ">
                            
                        </div>

						<div class="d-flex flex-row justify-content-end align-items-center mb-2">
							<a class="me-auto" @click="this.photoOwnerClick"> From {{owner}}</a>
							<a class="ms-2 me-2" @click="this.likeClick"> <i class="fa-regular fa-heart"></i>{{num_likes}}12933</a>
							<a class="ms-2 me-2" @click="this.commentClick"> <i class="fa-regular fa-comment"></i>{{num_comments}}210</a>
							
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
.bruh{
	color: blue;
	background-color:transparent;
}

.photo-background-color{
	background-color: grey;
}

.my-card{
	width: 34rem;
	border-color: black;
	border-width: thin;
}

.img-size{
	width: 150px;
}
</style>
