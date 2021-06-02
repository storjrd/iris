<style scoped>
.white-background {
	background-color: white;
}

.title {
	color: #4b93fe;
	font-weight: 600;
}

.back-btn {
	margin-left: 6%;
}
</style>

<template>
	<div>
		<button
			v-on:click="back"
			class="
				btn btn-success
				d-flex
				align-items-center
				back-btn
				btn-sm
				mb-1
			"
		>
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
		<file-browser v-if="ready"></file-browser>
	</div>
</template>

<script>
import FileBrowser from "../../browser/src/components/FileBrowser";

export default {
	data: () => ({
		ready: false
	}),
	computed: {
		bucket() {
			return this.$route.params.bucket;
		}
	},
	methods: {
		back() {
			this.$router.push({
				name: "unlock",
				params: {
					bucket: this.bucket
				}
			});
		}
	},
	components: {
		FileBrowser
	},
	async created() {
		const { accessKeyId, secretAccessKey } = await this.$store.dispatch(
			"gateway/getBucketCredentials",
			{
				name: this.bucket,
				passphrase: this.$store.state.buckets.passphrases[this.bucket]
			}
		);

		this.$store.commit("files/init", {
			endpoint: "gateway.tardigradeshare.io",
			accessKey: accessKeyId,
			secretKey: secretAccessKey,
			bucket: this.bucket,
			browserRoot: `/app/buckets/${this.bucket}/browse/`,
			openModalOnFirstUpload: false
		});

		this.ready = true;

		this.$store.dispatch("files/list");
	}
};
</script>
