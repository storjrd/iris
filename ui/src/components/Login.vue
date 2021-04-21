<template>
	<div>
		<h5 class="mb-4">Login</h5>

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

		<button class="btn btn-primary button signup-btn btn-block" @click="login">
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
		password: ""
	}),
	methods: {
		async login() {
			await this.$store.dispatch("account/login", { email: this.email, password: this.password });
			this.routeToBucketsView();
		},

		routeToBucketsView() {
			this.$router.push({ path: "/app/buckets" });
		}
	},
	created() {
		if (this.$store.state.account.token) {
			this.routeToBucketsView();
		}
	},
}
</script>
