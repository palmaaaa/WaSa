<script>
export default {
	data(){
		return{
			owner: null,

		}
	},
	methods:{
		prove: function(){
			console.log("bruh")
		}

	},
	props: ['owner','num_likes','num_comments',"upload_date"]
}
</script>

<template>

<div class="container-fluid mt-5 mb-5 ">
        <div class="d-flex flex-row justify-content-center">

            <div class="card my-card">
                <img src="../assets/images/nebula.jpg" class="card-img-top">

                <div class="card-body bruh">
                    
                    <div class="container">

                    	<div class="d-flex flex-row justify-content-start align-items-center ">
                            
                        </div>

						<div class="d-flex flex-row justify-content-end align-items-center mb-2">
							<a class="me-auto" @click="this.prove"> From {{owner}}</a>
							<a class="ms-2 me-2" @click="this.prove"> <i class="fa-regular fa-heart"></i>{{num_likes}}12933</a>
							<a class="ms-2 me-2" @click="this.prove"> <i class="fa-regular fa-comment"></i>{{num_comments}}210</a>
							
							
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


.my-card{
	width: 34rem;
	border-color: black;
	border-width: medium;
}
</style>
