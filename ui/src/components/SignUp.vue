<template>
	<div @keyup.enter="signUp">
		<h5 class="mb-4">Get Started</h5>

		<label for="emailAddress">Email Address</label>
		<input
			required
			type="email"
			class="form-control input email"
			placeholder="example@email.com"
			id="emailAddress"
			v-model="email"
		/>

		<label for="password">Password</label>
		<input
			required
			type="password"
			class="form-control input password"
			placeholder="••••••••••••"
			id="password"
			v-model="password"
		/>

		<div class="custom-control custom-checkbox mb-3">
			<input
				required
				type="checkbox"
				class="custom-control-input"
				id="termsCheck"
			/>
			<label class="custom-control-label" for="termsCheck"
				>Accept the
				<a
					href="https://tardigrade.io/terms-of-service/"
					target="_blank"
					>Terms &amp; Conditions</a
				></label
			>
		</div>

		<button class="btn btn-primary btn-block" @click="signUp">
			Try Storj
		</button>

		<hr />

		<router-link to="/login">
			<button class="btn btn-success btn-block">Login to Storj</button>
		</router-link>
	</div>
</template>

<script>
export default {
	data: () => ({
		email: "",
		password: ""
	}),
	methods: {
		async signUp() {
			await this.$store.dispatch("account/signUp", {
				email: this.email,
				password: this.password
			});

			this.routeToBucketsView();
		},
		routeToBucketsView() {
			console.log("going to buckets");
			this.$router.push({ path: "/app/buckets" });
		}
	},
	watch: {
		email() {
			this.$store.commit("account/setEmail", { email: this.email });
		}
	},
	created() {
		if (this.$store.state.account.token) {
			this.routeToBucketsView();
		}
	}
};
</script>
