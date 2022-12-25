<script>
export default {
	data: function() {
		return {
			errormsg: null,
		}
	},
	
}
// To add likes and comments put modal scrollable from bootstrap
</script>

<template>
	<div>
		<!--<Photo :upload_date="this.date" /> 
		<Photo/> -->
		
		
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
