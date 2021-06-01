<style scoped>
.no-border {
	border: none;
}

.unlocked-bucket {
	fill: #376fff;
}

.locked-bucket {
	fill: #768295;
}

.bucket-name,
.bucket-name span {
	cursor: pointer;
	font-weight: 700;
	color: #091c45;
}
.bucket-name:hover,
.bucket-name:hover span {
	color: #366fff;
}
.bucket-date {
	color: grey;
	font-weight: 300;
}
</style>

<template>
	<tr scope="row">
		<td v-on:click="unlock" class="bucket-name">
			<svg
				data-v-765080aa=""
				width="16"
				height="18"
				fill="none"
				xmlns="http://www.w3.org/2000/svg"
				class="ml-3"
				v-bind:class="{
					'unlocked-bucket': isUnlocked,
					'locked-bucket': !isUnlocked
				}"
			>
				<path
					data-v-765080aa=""
					fill-rule="evenodd"
					clip-rule="evenodd"
					d="M15.42 4.555l-1.047 11.463h-.003C14.267 17.119 11.327 18 7.715 18s-6.552-.882-6.656-1.981h-.002L.01 4.554c.037.03.076.059.116.088.458.332 1.09.62 1.866.857 1.506.462 3.51.725 5.647.731h.076c2.164 0 4.199-.264 5.723-.731.775-.238 1.408-.525 1.866-.857.04-.029.079-.058.116-.088zM7.715 0c4.001 0 7.245.995 7.245 2.222 0 1.227-3.244 2.221-7.245 2.221C3.714 4.443.47 3.45.47 2.222S3.714 0 7.715 0z"
				></path>
			</svg>
			<span class="ml-3">
				<span v-if="isUnlocked" v-on:click="browse" class="bucket-name">
					{{ bucket }}
				</span>

				<span>
					{{ bucket }}
				</span>
			</span>
		</td>
		<td class="bucket-date">4/16/2021, 12:32:52 PM</td>
		<td class="d-flex justify-content-lg-end">
			<div class="d-inline-flex">
				<div class="dropleft">
					<div
						v-if="false"
						class="spinner-border"
						role="status"
					></div>
					<button
						v-else
						class="btn btn-white btn-actions"
						type="button"
						aria-haspopup="true"
						aria-expanded="false"
						v-on:click=""
					>
						<svg
							width="4"
							height="16"
							viewBox="0 0 4 16"
							fill="none"
							xmlns="http://www.w3.org/2000/svg"
						>
							<path
								d="M3.2 1.6C3.2 2.48366 2.48366 3.2 1.6 3.2C0.716344 3.2 0 2.48366 0 1.6C0 0.716344 0.716344 0 1.6 0C2.48366 0 3.2 0.716344 3.2 1.6Z"
								fill="#7C8794"
							/>
							<path
								d="M3.2 8C3.2 8.88366 2.48366 9.6 1.6 9.6C0.716344 9.6 0 8.88366 0 8C0 7.11634 0.716344 6.4 1.6 6.4C2.48366 6.4 3.2 7.11634 3.2 8Z"
								fill="#7C8794"
							/>
							<path
								d="M1.6 16C2.48366 16 3.2 15.2837 3.2 14.4C3.2 13.5163 2.48366 12.8 1.6 12.8C0.716344 12.8 0 13.5163 0 14.4C0 15.2837 0.716344 16 1.6 16Z"
								fill="#7C8794"
							/>
						</svg>
					</button>
					<div class="dropdown-menu shadow show" v-if="false">
						<button
							v-if="false"
							class="dropdown-item action p-3"
							v-on:click=""
						>
							<svg
								width="1.5em"
								height="1.5em"
								viewBox="0 0 16 16"
								class="bi bi-x mr-1"
								fill="currentColor"
								xmlns="http://www.w3.org/2000/svg"
							>
								<path
									fill-rule="evenodd"
									d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z"
								/>
							</svg>
							Delete
						</button>
						<div v-if="false">
							<p class="deletion-confirmation mx-5 pt-3">
								Are you sure?
							</p>
							<div class="d-flex">
								<button
									class="dropdown-item trash p-3 action"
									v-on:click=""
								>
									<svg
										xmlns="http://www.w3.org/2000/svg"
										width="15"
										height="15"
										fill="red"
										class="bi bi-trash"
										viewBox="0 0 16 16"
									>
										<path
											d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"
										/>
										<path
											fill-rule="evenodd"
											d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4L4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"
										/>
									</svg>
									Yes
								</button>
								<button
									class="dropdown-item p-3 action"
									v-on:click=""
								>
									<svg
										width="1.5em"
										height="1.5em"
										viewBox="0 0 16 16"
										class="bi bi-x mr-1"
										fill="green"
										xmlns="http://www.w3.org/2000/svg"
									>
										<path
											fill-rule="evenodd"
											d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z"
										/>
									</svg>
									No
								</button>
							</div>
						</div>
					</div>
				</div>
			</div>
		</td>
	</tr>
</template>

<script>
import { toRefs } from "vue";

export default {
	name: "Bucket",
	props: ["bucket"],
	setup(props) {
		const { bucket } = toRefs(props);
		return {
			bucket
		};
	},
	computed: {
		isUnlocked() {
			return this.$store.getters["buckets/isUnlocked"](this.bucket);
		},

		bucketFill() {
			console.log("fill");
			return this.isUnlocked ? "#28a745" : "#768394";
		}
	},
	methods: {
		unlock() {
			this.$router.push({
				name: "unlock",
				params: {
					bucket: this.bucket
				}
			});
		},
		browse() {
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
