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
	methods: {},
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
