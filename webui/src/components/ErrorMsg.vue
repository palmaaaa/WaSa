<script>
export default {
	props: ['msg']
}
</script>

<template>
	<div class="d-flex justify-content-center alert alert-danger mt-3 error" role="alert">
		{{ msg }}
	</div>
</template>

<style>
.error {
    position: absolute;
    width: 50%;
    left: 25%;
    bottom: 0%;
}
</style>
