<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	methods:{
		mounted(){
			if (localStorage.getItem('token')){

				this.$router.replace("/login")
			}else{
				this.$router.replace("/home")
			}
			
		}
	}
}
</script>

<template>
	<div class="container-fluid">
		<div class="row">
			<div class="col p-0">
				<main >
					<RouterView />
				</main>
			</div>
		</div>
	</div>
</template>

<style>
</style>
