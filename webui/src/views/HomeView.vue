<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
		}
	},
	
}
</script>

<template>
	<div>
		
		<h1>Home</h1>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
