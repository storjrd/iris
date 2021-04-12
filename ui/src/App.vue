<template>
	<div id="nav">
		<router-link to="/">Home</router-link>
	</div>
	<router-view />
</template>

<style>
#app {
	font-family: Avenir, Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	color: #2c3e50;
}

#nav {
	padding: 30px;
}

#nav a {
	font-weight: bold;
	color: #2c3e50;
}

#nav a.router-link-exact-active {
	color: #42b983;
}
</style>

<script>
export default {
	async setup() {
		const response = await fetch("/api/v0/auth/register", {
			method: "POST",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify({
				fullName: "Monty Anderson",
				shortName: "Monty",
				email: `monty+iris-testing-${Math.random().toString().slice(2)}@storj.io`,
				password: "testtest123",
				isProfessional: false,
				position: "dj"
			})
		});

		console.log('response', response.json());
	},
};
</script>
