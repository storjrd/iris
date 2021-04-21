import { createStore } from "vuex";

import account from "./account";
import gateway from "./gateway";

export default createStore({
	state: {},
	mutations: {},
	actions: {},
	modules: {
		account,
		gateway
	}
});
