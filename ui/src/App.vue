<style scoped>
#app {
	background: #001030;
	font-family: "Inter", sans-serif;
	min-height: 100vh;
}

.about-tardigrade {
	font-weight: bold;
	font-size: 16px;
	color: #376fff;
}

@media (max-width: 575.98px) {
	.about-tardigrade {
		font-size: 14px;
	}
}
</style>

<template>
	<div id="app" v-on:click="closeAllInteractions">
		<div
			class="alert alert-warning text-center rounded-0 mb-0 border-bottom"
			role="alert"
		>
			Welcome! Please be aware that this is an alpha development version.
			Data and accounts will be deleted every few days.
		</div>
		<router-view></router-view>
	</div>
</template>

<script>
export default {
	computed: {
		isLoggedIn() {
			return this.$store.getters["account/isLoggedIn"];
		}
	},
	methods: {
		navigateLogin() {
			if (this.isLoggedIn === true) {
				this.$router.push({
					path: "/app/buckets"
				});
			} else {
				this.$router.push({
					path: "/"
				});
			}
		},

		logout() {
			this.$store.dispatch("account/logout");
		},

		closeAllInteractions() {
			this.$store.dispatch("files/closeAllInteractions");
			this.$store.dispatch("access/closeAllInteractions");

			// close dropdrowns in access page and buckets if it has any
		}
	},
	watch: {
		isLoggedIn() {
			this.navigateLogin();
		}
	},
	created() {
		this.navigateLogin();
	}
};
</script>
