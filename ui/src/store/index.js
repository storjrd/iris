import { createStore } from "vuex";
import { files } from "browser";

import account from "./account";
import gateway from "./gateway";
import buckets from "./buckets";

export default createStore({
	state: {},
	mutations: {},
	actions: {},
	modules: {
		account,
		gateway,
		buckets,
		files
	}
});
