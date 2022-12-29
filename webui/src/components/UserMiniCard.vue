<script>
export default {
	props: ['identifier','nickname'],

    methods:{

        clickedUser(){
            this.$emit('clickedUser',this.identifier)
        }
    },
}
</script>

<template>
    
    <div class="row mb-2 mt-2">
        <div class="col d-flex justify-content-center">
            <div class="user-mini-card card bg-transparent border-start">
                <div class="card-body">
                    <h5 @click="clickedUser" class="user-mini-card-title d-flex justify-content-center ">{{nickname }} @{{identifier }}</h5>
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
    color: black;
    width: 25%;
    border: none;
    border-color: black;
}

.user-mini-card-title:hover{
    cursor: pointer;
    font-weight: bold;
}
</style>
