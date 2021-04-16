<style scoped>
.white-background {
	background-color: white;
	border-radius: 5px;
}
</style>

<template>
	<div>
		<div class="container">
			<div class="white-background py-5">
				<router-view></router-view>
			</div>
		</div>
		<app-footer></app-footer>
	</div>
</template>

<script>
import AppFooter from "../components/AppFooter.vue";

export default {
	components: {
		AppFooter
	}
};
</script>
