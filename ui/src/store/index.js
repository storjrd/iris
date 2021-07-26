import { createStore } from "vuex";

import account from "./account";
import gateway from "./gateway";
import buckets from "./buckets";
import access from "./access";
import files from "../../browser/src/store/files";

export default createStore({
	state: {},
	mutations: {},
	actions: {},
	modules: {
		account,
		gateway,
		buckets,
		access,
		files
	}
});
