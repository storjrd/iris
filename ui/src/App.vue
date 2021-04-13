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
import { AccountAlreadyExistsError, register, login, getProjects, createApiKey } from "./lib/satellite.js";
import { generateAccess } from "./lib/access.js";
import { getMtCredentials } from "./lib/mt.js";

export default {
	async setup() {
		const fullName = "Monty Anderson";
		const shortName = "Monty";
		const email = "monty+iris-testing-5@storj.io";
		
		/*
		try {
			await register({
				fullName,
				shortName,
				email,
				password
			});
		} catch(error) { 
			// throw if failed for any other reason than already existing
			if(!(error instanceof AccountAlreadyExistsError)) {
				throw error;
			}
		}
		*/

		// get login token
		const { token } = await login({
			email,
			password
		});

		console.log({token});

		// find primary project
		const { myProjects: [ project ] } = await getProjects({ token });
		console.log({ project });

		if(typeof project !== "object") {
			throw new Error("Couldn't find project");
		}

		// create new api key for this session
		// todo: overwrite previous api key
		const { key } = await createApiKey({
			token,
			projectId: project.id,
			name: `Iris API Key: ${Date.now()}`
		});

		console.log({ key });

		// create access grant with access to all buckets
		const { access } = await generateAccess({
			key,
			projectId: project.id,
			satelliteUrl: "12tRQrMTWUWwzwGh18i7Fqs67kmdhH9t6aToeiwbo5mfS2rUmo@us2.tardigrade.io:7777",
			buckets: [],
			passphrase: "testestsetset"
		});

		console.log({ access });

		const { accessKey, secretKey } = await getMtCredentials({
			access
		});

		console.log({ accessKey, secretKey });
	}
};
</script>
