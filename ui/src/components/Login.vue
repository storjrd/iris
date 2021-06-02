<template>
	<div>
		<h5 class="mb-4">Login</h5>

		<div v-if="loginError" class="alert alert-danger" role="alert">
			{{ loginErrorMessage }}
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
			class="btn btn-primary button signup-btn btn-block"
			@click="login"
		>
			Login
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
		password: "",
		loginError: false
	}),
	methods: {
		async login() {
			try {
				await this.$store.dispatch("account/login", {
					email: this.email,
					password: this.password
				});

				this.routeToBucketsView();
			} catch(e) {
				this.loginError = true;
				this.setErrorMessage(e);
				this.password = "";
			}

		},

		routeToBucketsView() {
			this.$router.push({ path: "/app/buckets" });
		},

		setErrorMessage(e) {
			this.loginErrorMessage = e.message.split(":")[1].trim();
		}
	},
	created() {
		if (this.$store.state.account.token) {
			this.routeToBucketsView();
		}
	}
};
</script>
