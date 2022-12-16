<script>
export default {
	data: function() {
		return {
			date: "01/01/1800"
		}
	},
	
}
// To add likes and comments put modal scrollable from bootstrap
</script>

<template>
	<div>
		<Navbar ></Navbar>
		
		<Photo :upload_date="this.date" /> 
		<Photo/> 
		
		
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
