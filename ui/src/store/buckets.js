export default {
	state: () => ({
		names: [],
		passphrases: {}
	}),

	getters: {
		// is the bucket unlocked?
		isUnlocked: (state) => (bucket) => bucket in state.passphrases,
		s3: (state) => (bucket) => getS3ClientForBucketAndPassphrase()
	},

	mutations: {
		// called with the list of all buckets in a user's project
		setNames(state, names) {
			state.names = names;
		},

		// called when the user sets a passphrase
		unlock(state, { name, passphrase }) {
			state.passphrases[name] = passphrase;
		}
	},

	actions: {
		async getBuckets({ dispatch, commit }) {
			// get s3 client with access to all buckets [but no passphrase set]
			const s3 = await dispatch("gateway/getRootClient");

			const { Buckets } = await s3.listBuckets({}).promise();

			const names = Buckets.map(({ Name }) => Name);

			// update list of buckets
			commit("setNames", names);
		}
	}
};
