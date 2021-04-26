<style scoped>
.title {
	color: #4b93fe;
	font-weight: 600;
}
</style>

<template>
	<div>
	<h2 class="text-center mb-5 title">Access test-bucket</h2>
	<form class="container col-6">
		<div class="mb-3">
			<label for="name" class="form-label">Bucket Name</label>
			<input
				type="text"
				class="form-control"
				id="name"
				value="test-bucket"
				disabled
			/>
		</div>
		<div class="mb-3">
			<label for="passphrase" class="form-label">Passphrase</label>
			<input type="text" class="form-control" id="passphrase" />
		</div>
		<button type="submit" class="btn btn-primary">Continue</button>
	</form>
	</div>
</template>
