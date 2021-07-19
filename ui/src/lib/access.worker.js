import * as Comlink from "comlink";
import "./wasm_exec.js";

const go = new Go();

const wasmInit = WebAssembly.instantiateStreaming(
	fetch("/access.wasm"),
	go.importObject
);

wasmInit.then((result) => {
	go.run(result.instance);
});

export async function generateAccess({
	key,
	projectId,
	buckets = [],
	satelliteUrl = "12tRQrMTWUWwzwGh18i7Fqs67kmdhH9t6aToeiwbo5mfS2rUmo@us2.storj.io:7777",
	passphrase = ""
}) {
	await wasmInit;

	const permission = await global.newPermission();

	if (permission.error) {
		throw new Error(permission.error);
	}

	permission.value.AllowDownload = true;
	permission.value.AllowUpload = true;
	permission.value.AllowDelete = true;
	permission.value.AllowList = true;

	console.log(permission);

	const apiKeyPermissionParameters = [
		key, // api key
		buckets,
		permission.value
	];

	console.log({ apiKeyPermissionParams: apiKeyPermissionParameters });

	const restricted = await global.setAPIKeyPermission(
		...apiKeyPermissionParameters
	);

	if (restricted.error) {
		throw new Error(restricted.error);
	}

	const accessGrantParameters = [
		satelliteUrl,
		restricted.value,
		passphrase,
		projectId
	];

	console.log({ accessGrantParams: accessGrantParameters });

	const access = await global.generateAccessGrant(...accessGrantParameters);

	if (access.error) {
		throw new Error(access.error);
	}

	return {
		access: access.value
	};
}

Comlink.expose({
	generateAccess
});