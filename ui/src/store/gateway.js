import S3 from "aws-sdk/clients/s3";

import { generateAccess } from "../lib/access";
import { getCredentials } from "../lib/mt";

export default {
	namespaced: true,
	state: () => ({
		rootCredentials: null
	}),
	mutations: {
		setRootCredentials(state, credentials) {
			this.rootCredentials = credentials;
		}
	},
	actions: {
		async getRootCredentials({
			commit,
			state,
			rootState: {
				account: { apiKey, projectId }
			}
		}) {
			if (state.rootCredentials !== null) {
				return state.rootCredentials;
			}

			const { access } = await generateAccess({
				key: apiKey,
				projectId
			});

			const { accessKeyId, secretAccessKey } = await getCredentials({
				access
			});

			commit("setRootCredentials", {
				accessKeyId,
				secretAccessKey
			});

			return {
				accessKeyId,
				secretAccessKey
			};
		},

		async getBucketCredentials({
			rootState: {
				account: { apiKey, projectId }
			}
		}, { name, passphrase }) {
			const { access } = await generateAccess({
				key: apiKey,
				projectId,
				passphrase,
				buckets: [ name ]
			});

			const { accessKeyId, secretAccessKey } = await getCredentials({
				access
			});

			return {
				accessKeyId,
				secretAccessKey
			};
		},

		async getRootClient({ dispatch }) {
			const { accessKeyId, secretAccessKey } = await dispatch(
				"getRootCredentials"
			);

			const s3 = new S3({
				accessKeyId,
				secretAccessKey,
				endpoint: "https://gateway.tardigradeshare.io",
				s3ForcePathStyle: true,
				signatureVersion: "v4"
			});

			return s3;
		},

		async getBuckets({ dispatch }) {
			const s3 = await dispatch("getRootClient");

			const { Buckets } = await s3.listBuckets({}).promise();

			const buckets = Buckets.map(({ Name }) => Name);

			return buckets;
		},

		async createBucket({ dispatch }, { name }) {
			const s3 = await dispatch("getRootClient");

			await s3
				.createBucket({
					Bucket: name
				})
				.promise();
		}
	},
	getters: {}
};
