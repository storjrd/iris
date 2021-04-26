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
		<file-browser></file-browser>
	</div>
</template>

<script>
import { FileBrowser } from "browser";

export default {
	computed: {
		bucket() {
			return this.$route.params.bucket;
		}
	},
	methods: {},
	components: {
		FileBrowser
	}
};
</script>
