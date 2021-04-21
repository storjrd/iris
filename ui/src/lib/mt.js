export async function getCredentials({ access }) {
	const response = await fetch("https://auth.tardigradeshare.io/v1/access", {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify({
			access_grant: access,
			public: false
		})
	});

	const { access_key_id, secret_key } = await response.json();

	const credentials = {
		accessKeyId: access_key_id,
		secretAccessKey: secret_key
	};

	return credentials;
}
