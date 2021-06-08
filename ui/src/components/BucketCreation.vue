<style scoped>
.white-background {
	background-color: white;
}

.title {
	font-weight: 700;
}

.bucketname-input:focus {
	color: #fe5d5d;
	box-shadow: 0 0 0 0.2rem rgba(254, 93, 93, 0.5) !important;
	border-color: #fe5d5d !important;
	outline: none !important;
}

.bucketpassphrase-input:focus {
	color: #fe5d5d;
	box-shadow: 0 0 0 0.2rem rgba(254, 93, 93, 0.5) !important;
	border-color: #fe5d5d !important;
	outline: none !important;
}
</style>

<template>
	<div class="row justify-content-center">
		<div
			class="col-12 col-sm-10 col-md-8 col-lg-7 col-xl-5 white-background"
		>
			<h2 class="text-center mb-5 title">Create a New Bucket</h2>
			<div
				v-if="creatingBucketSpinner"
				class="mt-4 d-flex justify-content-center"
			>
				<div class="spinner-border text-center" role="status"></div>
			</div>

			<div class="mb-4">
				<label for="name" class="form-label">Bucket Name</label>
				<input
					type="text"
					class="form-control"
					id="name"
					v-model="bucketName"
					v-bind:class="bucketNameClass"
				/>
			</div>
			<div class="mb-4">
				<label for="passphrase" class="form-label">Passphrase</label>
				<input
					type="text"
					class="form-control"
					id="passphrase"
					v-model="bucketPassphrase"
					v-bind:class="bucketPassphraseClass"
				/>
			</div>
			<div class="d-flex">
				<router-link to="/app/buckets" class="btn btn-light btn-block">
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
				</router-link>

				<div class="mx-1"></div>

				<button
					class="btn btn-primary btn-block"
					v-bind:disabled="!createBucketEnabled"
					v-on:click="createBucket"
				>
					Create
				</button>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	name: "BucketCreation",
	data: () => ({
		bucketName: "",
		bucketPassphrase: "",
		creatingBucketSpinner: false
	}),
	computed: {
		bucketNameValid() {
			return (
				this.bucketName.length > 0 &&
				!this.$store.state.buckets.names.includes(this.bucketName) &&
				!this.bucketName.includes(" ") &&
				this.bucketName.toLowerCase() === this.bucketName
			);
		},

		bucketPassphraseValid() {
			return this.bucketPassphrase.length > 0;
		},

		createBucketEnabled() {
			return this.bucketNameValid && this.bucketPassphraseValid;
		},

		bucketNameClass() {
			return {
				"bucketname-input":
					this.bucketName.length > 0 && !this.bucketNameValid
			};
		},

		bucketPassphraseClass() {
			return {
				"bucketpassphrase-input":
					this.bucketPassphrase.length > 0 &&
					!this.bucketPassphraseValid
			};
		}
	},
	methods: {
		async createBucket() {
			if (this.createBucketEnabled) {
				this.creatingBucketSpinner = true;

				await this.$store.dispatch("buckets/createBucket", {
					name: this.bucketName
				});

				this.$store.commit("buckets/unlock", {
					name: this.bucketName,
					passphrase: this.bucketPassphrase
				});

				this.$router.push({
					name: "browse",
					params: {
						bucket: this.bucketName
					}
				});

				this.creatingBucketSpinner = false;
				this.bucketName = "";
				this.bucketPassphrase = "";
			}
		}
	},

	async created() {
		await this.$store.dispatch("buckets/list");
	}
};
</script>
