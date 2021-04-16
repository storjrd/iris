<style scoped>
.white-background {
	background-color: white;
}

.title {
	color: #4b93fe;
	font-weight: 600;
}
</style>

<template>
	<h2 class="text-center mb-5 title">Create a New Bucket</h2>
	<form class="container col-6 white-background">
		<div class="mb-3">
			<label for="name" class="form-label">Bucket Name</label>
			<input type="text" class="form-control" id="name" />
		</div>
		<div class="mb-3">
			<label for="passphrase" class="form-label">Passphrase</label>
			<input type="text" class="form-control" id="passphrase" />
		</div>
		<button type="submit" class="btn btn-primary">Create</button>
	</form>
</template>
