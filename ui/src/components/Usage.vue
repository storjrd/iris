<style media="screen" scoped>
.progress {
	border-radius: 1rem;
}
.progress-bar {
	background-color: #0059d0;
}

.plan-table {
	border-collapse: separate;
	border-spacing: 0 10px;
}
.plan-table th {
	border-top: 0;
	border-bottom: 1px solid #dee2e6;
	color: #768394;
	font-size: 14px;
}
.plan-table td {
	height: 80px;
	border: 0;
}
.plan-row {
	padding: 10px;
}
.plan-selected {
	color: #fff;
	background: #0059d0;
	border-radius: 10px;
}
.plan-selected td {
	border: solid 1px #0059d0;
	border-style: solid none;
	padding: 10px;
}
.plan-selected td:first-child {
	border-left-style: solid;
	border-top-left-radius: 10px;
	border-bottom-left-radius: 10px;
}
.plan-selected td:last-child {
	border-right-style: solid;
	border-bottom-right-radius: 10px;
	border-top-right-radius: 10px;
}
.plan-name {
	font-weight: bold;
}
.btn-upgrade {
	background: #eff1f3;
	color: #0068d0;
	min-width: 164px;
}
.usage-mute {
	color: #aaa;
}
</style>
<template>
	<div class="card card-top-flat border-0 p-4 p-lg-5">
		<div class="bars" v-if="usage != null">
			<div class="row mb-4">
				<div class="col">
					<h5 class="mb-4">
						Usage <span class="usage-mute">(updates daily)</span>
					</h5>
					<hr />
				</div>
			</div>

			<div class="row mb-4">
				<div class="col-md-12 col-lg-4 mb-4 mb-lg-0">
					<div class="row">
						<div class="col">
							<p class="text-small">
								<svg
									width="1em"
									height="1em"
									viewBox="0 0 16 16"
									class="bi bi-hdd mr-1"
									fill="currentColor"
									xmlns="http://www.w3.org/2000/svg"
								>
									<path
										fill-rule="evenodd"
										d="M14 9H2a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1zM2 8a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-1a2 2 0 0 0-2-2H2z"
									/>
									<path
										d="M5 10.5a.5.5 0 1 1-1 0 .5.5 0 0 1 1 0zm-2 0a.5.5 0 1 1-1 0 .5.5 0 0 1 1 0z"
									/>
									<path
										fill-rule="evenodd"
										d="M4.094 4a.5.5 0 0 0-.44.26l-2.47 4.532A1.5 1.5 0 0 0 1 9.51v.99H0v-.99c0-.418.105-.83.305-1.197l2.472-4.531A1.5 1.5 0 0 1 4.094 3h7.812a1.5 1.5 0 0 1 1.317.782l2.472 4.53c.2.368.305.78.305 1.198v.99h-1v-.99a1.5 1.5 0 0 0-.183-.718L12.345 4.26a.5.5 0 0 0-.439-.26H4.094z"
									/>
								</svg>
								<strong>Storage</strong>
							</p>
						</div>
						<div class="col">
							<p class="text-right">
								<span class="metric"
									>{{ bytesUploadedPrettyBytes }}/{{
										bytesUploadedQuotaPrettyBytes
									}}</span
								>
							</p>
						</div>
					</div>
					<div class="progress">
						<div
							class="progress-bar"
							role="progressbar"
							v-bind:style="`width: ${bytesUploadedPercent}%;`"
						></div>
					</div>
				</div>

				<div class="col-md-12 col-lg-4 mb-4 mb-lg-0">
					<div class="row">
						<div class="col">
							<p class="text-small">
								<svg
									width="1em"
									height="1em"
									viewBox="0 0 16 16"
									class="bi bi-files mr-1"
									fill="currentColor"
									xmlns="http://www.w3.org/2000/svg"
								>
									<path
										fill-rule="evenodd"
										d="M4 2h7a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2zm0 1a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4z"
									/>
									<path
										d="M6 0h7a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2v-1a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1H4a2 2 0 0 1 2-2z"
									/>
								</svg>
								<strong>Files</strong>
							</p>
						</div>
						<div class="col">
							<p class="text-right">
								<span class="metric"
									>{{ filesUploadedToHumanString }}/{{
										filesUploadedQuotaToHumanString
									}}</span
								>
							</p>
						</div>
					</div>
					<div class="progress">
						<div
							class="progress-bar"
							role="progressbar"
							v-bind:style="`width: ${filesUploadedPercent}%;`"
						></div>
					</div>
				</div>

				<div class="col-md-12 col-lg-4 mb-4 mb-lg-0">
					<div class="row">
						<div class="col">
							<p class="text-small">
								<svg
									width="1em"
									height="1em"
									viewBox="0 0 16 16"
									class="bi bi-arrow-down-up mr-1"
									fill="currentColor"
									xmlns="http://www.w3.org/2000/svg"
								>
									<path
										fill-rule="evenodd"
										d="M11.5 15a.5.5 0 0 0 .5-.5V2.707l3.146 3.147a.5.5 0 0 0 .708-.708l-4-4a.5.5 0 0 0-.708 0l-4 4a.5.5 0 1 0 .708.708L11 2.707V14.5a.5.5 0 0 0 .5.5zm-7-14a.5.5 0 0 1 .5.5v11.793l3.146-3.147a.5.5 0 0 1 .708.708l-4 4a.5.5 0 0 1-.708 0l-4-4a.5.5 0 0 1 .708-.708L4 13.293V1.5a.5.5 0 0 1 .5-.5z"
									/>
								</svg>
								<strong>Bandwidth</strong>
							</p>
						</div>
						<div class="col">
							<p class="text-right">
								<span class="metric"
									>{{ bytesDownloadedPrettyBytes }}/{{
										bytesDownloadedQuotaPrettyBytes
									}}</span
								>
							</p>
						</div>
					</div>
					<div class="progress">
						<div
							class="progress-bar"
							role="progressbar"
							v-bind:style="`width: ${bytesDownloadedPercent}%;`"
						></div>
					</div>
				</div>
			</div>

			<div class="row">
				<div class="col">
					<h5 class="mt-5 mb-3">Plan Details</h5>
				</div>
			</div>

			<div class="row">
				<div class="col">
					<div class="table-responsive">
						<table class="table plan-table">
							<thead>
								<tr>
									<th scope="col"></th>
									<th scope="col" class="w-25">Plan</th>
									<th scope="col">Storage</th>
									<th scope="col">Bandwidth</th>
									<th scope="col">File Limit</th>
									<th scope="col"></th>
								</tr>
							</thead>
							<tbody>
								<tr
									v-bind:class="{
										'plan-row': true,
										'plan-selected': id == planId
									}"
									v-for="(plan, id) in plans"
									v-bind:key="id"
								>
									<td class="text-right">
										<img
											src="@/assets/checkmark.svg"
											alt="Checkmark"
											width="14"
										/>
									</td>
									<td class="plan-name w-25">
										{{ plan.name }}
									</td>
									<td>
										{{
											formatQuota(plan.storageBytesQuota)
										}}
									</td>
									<td>
										{{
											formatQuota(plan.downloadBytesQuota)
										}}
									</td>
									<td>
										{{
											formatStorageFilesQuota(
												plan.storageFilesQuota
											)
										}}
									</td>
									<td>
										<button
											v-if="free100G(id)"
											type="button"
											class="btn btn-light btn-sm"
										>
											Verify Email
										</button>

										<router-link
											to="/app/usage/upgrade-form"
										>
											<button
												v-if="free1Tb(id)"
												type="button"
												class="btn btn-light btn-sm"
											>
												Upgrade - Free!
											</button>
										</router-link>
									</td>
								</tr>
							</tbody>
						</table>
					</div>
				</div>
			</div>

			<div class="row">
				<div class="col">
					<p class="text-center mt-4">
						Need more?
						<a href="mailto:iris@storj.io" class="font-weight-bold"
							>Contact Us</a
						>
					</p>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
import prettyBytes from "pretty-bytes";
import { toHumanString } from "human-readable-numbers";

export default {
	methods: {
		async updateUsage() {
			// const { data } = await axios.post("/api/usage");
			// this.usage = data;
		},
		formatQuota(quota) {
			return prettyBytes(quota);
		},
		formatStorageFilesQuota(quota) {
			return quota.toLocaleString();
		},
		free100G(id) {
			return id === "free-100g";
		},
		free1Tb(id) {
			return id === "free-1tb";
		}
	},
	computed: {
		usage() {
			return this.$store.state.account.usage;
		},
		plans() {
			return this.$store.state.account.plans;
		},
		planId() {
			return this.$store.state.account.planId;
		},
		bytesUploadedPrettyBytes() {
			return prettyBytes(this.usage.bytesUploaded);
		},
		bytesUploadedQuotaPrettyBytes() {
			return prettyBytes(this.usage.bytesUploadedQuota);
		},
		bytesDownloadedPrettyBytes() {
			return prettyBytes(this.usage.bytesDownloaded);
		},
		bytesDownloadedQuotaPrettyBytes() {
			return prettyBytes(this.usage.bytesDownloadedQuota);
		},
		filesUploadedToHumanString() {
			return toHumanString(this.usage.filesUploaded);
		},
		filesUploadedQuotaToHumanString() {
			return toHumanString(this.usage.filesUploadedQuota);
		},
		bytesUploadedPercent() {
			return this.usage !== null
				? (
						(this.usage.bytesUploaded /
							this.usage.bytesUploadedQuota) *
						100
				  ).toFixed(2)
				: 0;
		},

		filesUploadedPercent() {
			return this.usage !== null
				? (
						(this.usage.filesUploaded /
							this.usage.filesUploadedQuota) *
						100
				  ).toFixed(2)
				: 0;
		},

		bytesDownloadedPercent() {
			return this.usage !== null
				? (
						(this.usage.bytesDownloaded /
							this.usage.bytesDownloadedQuota) *
						100
				  ).toFixed(2)
				: 0;
		}
	},
	async created() {
		this.$store.dispatch("account/getUsage");
	}
};
</script>
