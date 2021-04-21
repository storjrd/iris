export async function getMtCredentials({ access }) {
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
		accessKey: access_key_id,
		secretKey: secret_key
	};

	return credentials;
}
