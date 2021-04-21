<template>
	<div>
		<h5 class="mb-4">Get Started</h5>

		<label for="emailAddress">Email Address</label>
		<input
			required
			type="email"
			class="form-control input email"
			placeholder="example@email.com"
			id="emailAddress"
		/>

		<label for="password">Password</label>
		<input
			required
			type="password"
			class="form-control input password"
			placeholder="••••••••••••"
			id="password"
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

		<button class="btn btn-primary button btn-block">Try Storj</button>

		<hr />

		<router-link to="/login">
			<button class="button-no-bg btn btn-success btn-block">
				Login to Storj
			</button>
		</router-link>
	</div>
</template>
