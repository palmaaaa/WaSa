<script>
export default {
    data(){
        return {
            user: "",
        }
    },
	props: ['content','author','photo_owner'],

    mounted(){
        this.user = localStorage.getItem('token')
    }

}
</script>

<template>
	<div class="container-fluid">

        <hr>
        <div class="row">
            <div class="col-10">
                {{author}}
            </div>

            <div class="col-2">
                <button v-if="this.user === this.author || this.user === photo_owner" class="btn my-btn-comm">
                    <i class="fa-regular fa-trash-can my-trash-icon"></i>
                </button>
            </div>

        </div>

        <div class="row">
            <div class="col-12">
                {{content}}
            </div>

        </div>
        <hr>
    </div>
</template>

<style>
.my-btn-comm{
    border: none;
}
.my-btn-comm:hover{
    border: none;
    transform: scale(1.1)
}

</style>
