<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {}
</script>

<template>
	<div v-cloak class="container-fluid">
		<div class="row">
			<main class="col">
				<RouterView />
			</main>
		</div>
	</div>
</template>

<style>
</style>
