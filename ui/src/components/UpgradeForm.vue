<style scoped>
label {
	font-weight: bold;
}

.form-check {
	display: flex;
	align-items: center;
}

.form-check-input {
	margin-top: 0px;
}
</style>

<template>
	<div class="row">
		<div class="col-1">
			<router-link to="/app/usage">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					width="16"
					height="16"
					fill="currentColor"
					class="bi bi-chevron-left mx-auto"
					viewBox="0 0 16 16"
				>
					<path
						fill-rule="evenodd"
						d="M11.354 1.646a.5.5 0 0 1 0 .708L5.707 8l5.647 5.646a.5.5 0 0 1-.708.708l-6-6a.5.5 0 0 1 0-.708l6-6a.5.5 0 0 1 .708 0z"
					/>
				</svg>
			</router-link>
		</div>

		<div class="col-10">
			<h2>Free Pro Upgrade - Usage Survey</h2>
			<p>
				Receive a free upgrade to a 1 TB Pro Plan, by answering a few
				usage questions.
			</p>
			<hr />

			<div class="col-7 px-0">
				<div class="mb-4">
					<label for="name" class="form-label">Name *</label>
					<input
						type="text"
						class="form-control"
						id="name"
						placeholder="Your Name"
						v-model="name"
					/>
				</div>

				<div class="mb-4">
					<label for="use-case" class="form-label"
						>What is your primary use case? *</label
					>
					<input
						type="text"
						class="form-control"
						id="use-case"
						placeholder="Choose"
						v-model="use_case"
					/>
				</div>

				<div class="mb-4">
					<label for="storage" class="form-label"
						>How much storage do you need for your use case?
						*</label
					>
					<input
						type="text"
						class="form-control"
						id="storage"
						placeholder="Choose"
						v-model="storage"
					/>
				</div>

				<div class="mb-4">
					<label for="bandwidth" class="form-label"
						>How much bandwidth do you need for your use case?
						*</label
					>
					<input
						type="text"
						class="form-control"
						id="bandwidth"
						placeholder="Choose"
						v-model="bandwidth"
					/>
				</div>

				<div class="mb-4">
					<label class="form-label"
						>What integrations would you be interested in? *</label
					>

					<div class="row mb-3">
						<div class="col-6">
							<div class="form-check mb-4">
								<input
									class="form-check-input"
									type="checkbox"
									value="rclone-integration"
									id="rclone-integration"
									v-model="rclone_integration"
								/>
								<label
									class="form-check-label"
									for="rclone-integration"
								>
									Rclone (File Transfer)
								</label>
							</div>

							<div class="form-check mb-4">
								<input
									class="form-check-input"
									type="checkbox"
									value="restic-integration"
									id="restic-integration"
									v-model="restic_integration"
								/>
								<label
									class="form-check-label"
									for="restic-integration"
								>
									Restic (Backup)
								</label>
							</div>

							<div class="form-check mb-4">
								<input
									class="form-check-input"
									type="checkbox"
									value="filezilla-integration"
									id="filezilla-integration"
									v-model="filezilla_integration"
								/>
								<label
									class="form-check-label"
									for="filezilla-integration"
								>
									Filezilla (File Transfer)
								</label>
							</div>

							<div class="form-check mb-4">
								<input
									class="form-check-input"
									type="checkbox"
									value="duplicati-integration"
									id="duplicati-integration"
									v-model="duplicati_integration"
								/>
								<label
									class="form-check-label"
									for="duplicati-integration"
								>
									Duplicati (Backup)
								</label>
							</div>

							<div class="form-check mb-4">
								<input
									class="form-check-input"
									type="checkbox"
									value="qnap-integration"
									id="qnap-integration"
									v-model="qnap_integration"
								/>
								<label
									class="form-check-label"
									for="qnap-integration"
								>
									QNAP (Backup)
								</label>
							</div>
						</div>

						<div class="col-6">
							<div class="form-check mb-4">
								<input
									class="form-check-input"
									type="checkbox"
									value="synology-integration"
									id="synology-integration"
									v-model="synology_integration"
								/>
								<label
									class="form-check-label"
									for="synology-integration"
								>
									Synology (Backup)
								</label>
							</div>

							<div class="form-check mb-4">
								<input
									class="form-check-input"
									type="checkbox"
									value="dropbox-integration"
									id="dropbox-integration"
									v-model="dropbox_integration"
								/>
								<label
									class="form-check-label"
									for="dropbox-integration"
								>
									"Dropbox" like Client
								</label>
							</div>

							<div class="form-check mb-4">
								<input
									class="form-check-input"
									type="checkbox"
									value="desktop-integration"
									id="desktop-integration"
									v-model="desktop_integration"
								/>
								<label
									class="form-check-label"
									for="desktop-integration"
								>
									Desktop Client
								</label>
							</div>

							<div class="form-check mb-4">
								<input
									class="form-check-input"
									type="checkbox"
									value="web-integration"
									id="web-integration"
									v-model="web_integration"
								/>
								<label
									class="form-check-label"
									for="web-integration"
								>
									Web Interface
								</label>
							</div>

							<div class="form-check mb-4">
								<input
									class="form-check-input"
									type="checkbox"
									value="other-integration"
									id="other-integration"
									v-model="other_integration"
								/>
								<label
									class="form-check-label"
									for="other-integration"
								>
									Other
								</label>
							</div>
						</div>
					</div>
					<input
						v-if="displayOtherIntegrationField"
						type="text"
						class="form-control"
						id="other"
						placeholder="Other"
						v-model="other"
					/>
				</div>

				<div class="mb-4">
					<label for="comments" class="form-label"
						>Other thoughts or comments</label
					>
					<input
						type="text"
						class="form-control"
						id="comments"
						placeholder="Your answer"
						v-model="comments"
					/>
				</div>

				<button
					type="button"
					class="btn btn-primary btn-lg"
					v-bind:disabled="!submitEnabled"
					v-on:click="submit"
				>
					Submit
				</button>
			</div>
		</div>

		<div class="col-1"></div>
	</div>
</template>

<script>
export default {
	name: "UpgradeForm",
	data: () => ({
		name: "",
		use_case: "",
		storage: "",
		bandwidth: "",
		rclone_integration: "",
		restic_integration: "",
		filezilla_integration: "",
		duplicati_integration: "",
		qnap_integration: "",
		synology_integration: "",
		dropbox_integration: "",
		desktop_integration: "",
		web_integration: "",
		other_integration: "",
		other: "",
		comments: ""
	}),
	computed: {
		displayOtherIntegrationField() {
			return this.other_integration;
		},

		submitEnabled() {
			return (
				this.name &&
				this.use_case &&
				this.storage &&
				this.bandwidth &&
				(this.rclone_integration ||
					this.restic_integration ||
					this.filezilla_integration ||
					this.duplicati_integration ||
					this.qnap_integration ||
					this.synology_integration ||
					this.dropbox_integration ||
					this.desktop_integration ||
					this.web_integration ||
					this.other_integration)
			);
		}
	},
	methods: {
		submit() {
			if (!this.submitEnabled) {
				return;
			}
		}
	}
};
</script>
