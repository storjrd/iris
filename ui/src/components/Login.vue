<style scoped>
.login {
	font-weight: bold;
}
</style>

<template>
	<div>
		<h5 class="mb-4">Login</h5>

		<div v-if="errorExists" class="alert alert-danger" role="alert">
			{{ errorMessage }}
		</div>

		<label for="emailAddress">Email Address</label>
		<input
			type="email"
			class="form-control input email"
			placeholder="example@email.com"
			v-on:keyup.enter="signUp"
			id="emailAddress"
			v-model="email"
		/>

		<label for="password">Password</label>
		<input
			type="password"
			class="form-control input password"
			placeholder="••••••••••••"
			id="password"
			v-model="password"
		/>

		<button
			class="btn btn-primary signup-btn btn-block my-4"
			@click="login"
		>
			<span
				v-if="loggingIn"
				class="spinner-border text-light align-middle"
				role="status"
			></span>
			<span v-else class="login">Login</span>
		</button>

		<p>
			Don't have an account?

			<router-link to="/">Sign Up</router-link>
		</p>
	</div>
</template>

<script>
export default {
	data: () => ({
		email: "",
		password: ""
	}),
	methods: {
		async login() {
			await this.$store.dispatch("account/login", {
				email: this.email,
				password: this.password
			});

			this.routeToBucketsView();
		},

		routeToBucketsView() {
			this.$router.push({ path: "/app/buckets" });
		}
	},
	computed: {
		errorMessage() {
			return this.$store.state.account.errorMessage;
		},

		emailInputFromSignupForm() {
			return this.$store.state.account.email;
		},

		errorExists() {
			return this.errorMessage && this.errorMessage.length > 0;
		},

		emailFromSignupExists() {
			return (
				this.emailInputFromSignupForm &&
				this.emailInputFromSignupForm.length > 0
			);
		},

		loggingIn() {
			return this.$store.state.account.loggingIn;
		}
	},
	created() {
		if (this.$store.state.account.token) {
			this.routeToBucketsView();
		}

		if (this.emailFromSignupExists) {
			this.email = this.emailInputFromSignupForm;
		}
	}
};
</script>
