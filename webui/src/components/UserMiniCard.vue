<script>
export default {
	props: ['identifier'],
    methods:{
        clickedUser(){
            console.log(this.identifier)
            this.$emit('clickedUser',this.identifier)
        }
    }
}
</script>

<template>
    
    

   <div class="row mb-2 mt-2">
        <div class="col d-flex justify-content-center">
            <div class="user-mini-card card bg-transparent border-start">
                <div class="card-body">
                    <h5 @click="clickedUser" class="user-mini-card-title d-flex justify-content-center">{{ identifier }}</h5>
                    <!--
                        user-mini-card-title
                        <a href="#" class="btn btn-primary">Button</a>-->
                </div>
            </div>
        </div>
    </div>
    
</template>

<style>
.user{
    color: var(--color-text);
}

.user-mini-card{
    color: var(--color-text);
    width: 25%;
    border: none;
    border-color: black;
}

.user-mini-card-title:hover{
    cursor: pointer;
    font-weight: bold;
}
</style>
