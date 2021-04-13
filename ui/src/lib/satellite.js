export class SatelliteError extends Error {};
export class AccountAlreadyExistsError extends SatelliteError {};

function throwIfSatelliteError(json) {
	if(typeof json.error === "string") {
		if(json.error.includes("This email is already in use, try another")) {
			throw new AccountAlreadyExistsError(json.error);
		}

		throw new SatelliteError(json.error);
	}
}

export async function register({ fullName, shortName, email, password }) {
	const response = await fetch("/api/v0/auth/register", {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify({
			fullName,
			shortName,
			email,
			password,
			isProfessional: false,
			position: "dj",
			secret: ""
		})
	});

	const token = await response.json();

	throwIfSatelliteError(token);

	if (typeof token !== "string") {
		throw new SatelliteError(
			`Satellite register returned bad value ${JSON.stringify(token)}`
		);
	}
}

export async function login({ email, password }) {
	const response = await fetch("/api/v0/auth/token", {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify({
			email,
			password
		})
	});

	const token = await response.json();

	throwIfSatelliteError(token);

	return {token};
}

export async function getProjects({ token }) {
	const response = await fetch("/api/v0/graphql", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
			"Cookie": `_tokenKey=${token}`
		},
		body: JSON.stringify({
			query: `{
				myProjects {
					name
					id
					description
					createdAt
					ownerId
					__typename
				}
			}`
		})
	});

	const { data: { myProjects } } = await response.json();

	return {
		myProjects
	};
}

export async function createApiKey({ token, projectId, name }) {
	const response = await fetch("/api/v0/graphql", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
			"Cookie": `_tokenKey=${token}`
		},
		body: JSON.stringify({
			query: `
				mutation ($projectId: String!, $name: String!) {
					createAPIKey(projectID: $projectId, name: $name) {
						key
						__typename
					}
				}
			  `,
			variables: {
				projectId,
				name
			}
		})
	});

	const {
		data: {
			createAPIKey: {
				key
			}
		}
	} = await response.json();

	return { key };
}

