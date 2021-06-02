<style scoped>
.white-background {
	background-color: white;
}

.title {
	font-weight: 600;
}
</style>

<template>
	<div>
		<h2 class="text-center mb-3 title">Unlock a bucket</h2>

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

			<div class="d-flex">
				<router-link to="/app/buckets">
					<button class="btn btn-success d-flex align-items-center">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="16"
							height="16"
							fill="currentColor"
							class="bi bi-arrow-left-short"
							viewBox="0 0 16 16"
						>
							<path
								fill-rule="evenodd"
								d="M12 8a.5.5 0 0 1-.5.5H5.707l2.147 2.146a.5.5 0 0 1-.708.708l-3-3a.5.5 0 0 1 0-.708l3-3a.5.5 0 1 1 .708.708L5.707 7.5H11.5a.5.5 0 0 1 .5.5z"
							/>
						</svg>
						Back
					</button>
				</router-link>

				<div class="mx-1"></div>

				<button
					type="submit"
					class="btn btn-primary"
					v-on:click="unlock"
				>
					Unlock
				</button>
			</div>
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
