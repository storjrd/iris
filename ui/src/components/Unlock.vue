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
	<div>
		<h2 class="text-center mb-5 title">Unlock a bucket</h2>

		<div class="container col-6 white-background">
			<div class="mb-3">
				<label for="name" class="form-label">Bucket Name</label>
				<input
					type="text"
					class="form-control"
					disabled
					v-model="bucket"
				/>
			</div>

			<div class="mb-3">
				<label for="passphrase" class="form-label">Passphrase</label>
				<input type="text" class="form-control" v-model="passphrase" />
			</div>

			<button type="submit" class="btn btn-primary" v-on:click="unlock">
				Unlock
			</button>
		</div>
	</div>
</template>

<script>
export default {
	data: () => ({
		passphrase: ""
	}),
	computed: {
		bucket() {
			return this.$route.params.bucket;
		}
	},
	methods: {
		unlock() {
			this.$store.commit("buckets/unlock", {
				name: this.bucket,
				passphrase: this.passphrase
			});

			this.$router.push({
				name: "browse",
				params: {
					bucket: this.bucket
				}
			});
		}
	}
};
</script>
