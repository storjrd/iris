import { getAccessKeys, deleteAccessKeys } from "../lib/satellite";

export default {
	namespaced: true,
	state: () => ({
		accessKeys: [],
		headingSorted: "name",
		orderBy: "asc",
		asc: 2,
		desc: 1,
		pageCount: 1,
		nameHeadingOrder: 1,
		dateCreatedHeadingOrder: 2,
		openedDropdown: null,
		accessCreationModal: null,
		accessKeysBeingDeleted: [],
		accessTableDropdown: null,
		selectedAccessKeys: [],
		satellite: {
			limit: 10,
			page: 1,
			search: "",
			order: 1,
			orderDirection: 2
		}
	}),
	getters: {},
	mutations: {
		setAccessKey(state, key) {
			state.accessKeys = [...state.accessKeys, key];
		},
		updateAccessKeys(state, keys) {
			state.accessKeys = keys;
		},
		setDropdown(state, id) {
			state.openedDropdown = id;
		},
		setAccessTableDropdown(state, value) {
			state.accessTableDropdown = value;
		},
		setHeadingSorted(state, heading) {
			state.headingSorted = heading;
		},
		setOrderBy(state, order) {
			state.orderBy = order;
		},
		setAccessCreationModal(state, value) {
			state.accessCreationModal = value;
		},
		setAccessKeysBeingDeleted(state, keys) {
			state.accessKeysBeingDeleted = [
				...state.accessKeysBeingDeleted,
				...keys
			];
		},
		removeAllAccessKeyBeingDeleted(state) {
			state.accessKeysBeingDeleted = [];
		},
		setSelectedAccessKey(state, key) {
			state.selectedAccessKeys = [...state.selectedAccessKeys, key];
		},
		removeSelectedAccessKey(state, key) {
			state.selectedAccessKeys = state.selectedAccessKeys.filter(
				(accessKey) => accessKey !== key
			);
		},
		removeAllSelectedAccessKeys(state) {
			state.selectedAccessKeys = [];
		},
		setPageCount(state, count) {
			state.pageCount = count;
		},
		updateSatelliteRequestObject(state) {
			if (state.headingSorted === "name") {
				state.satellite.order = state.nameHeadingOrder;
			} else {
				state.satellite.order = state.dateCreatedHeadingOrder;
			}

			if (state.orderBy === "asc") {
				state.satellite.orderDirection = state.asc;
			} else {
				state.satellite.orderDirection = state.desc;
			}
		}
	},
	actions: {
		async list({ commit, state, rootState }) {
			const { token, projectId } = rootState.account;
			const { limit, search, page, order, orderDirection } =
				state.satellite;

			const data = await getAccessKeys({
				token,
				projectId,
				limit,
				search,
				page,
				order,
				orderDirection
			});
			const { apiKeys, pageCount } = data;

			commit("updateAccessKeys", apiKeys);
			commit("setPageCount", pageCount);
		},
		async createAccessKey({ commit }, { key }) {
			// create access key
			// commit("setAccessKey", key);
		},
		async deleteAccessKeys({ commit, rootState, dispatch }, keyIDs) {
			const { token } = rootState.account;

			if (!Array.isArray(keyIDs)) {
				keyIDs = [keyIDs];
			}

			commit("setAccessKeysBeingDeleted", keyIDs);
			await deleteAccessKeys({ token, id: keyIDs });
			commit("removeAllAccessKeyBeingDeleted");
			dispatch("list");
		},
		addSelectedAccessKey({ commit }, key) {
			commit("setSelectedAccessKey", key);
		},
		deleteSelectedAccessKey({ commit }, key) {
			commit("removeSelectedAccessKey", key);
		},
		async deleteSelectedAccessKeys({ dispatch, commit, state }) {
			dispatch("deleteAccessKeys", [...state.selectedAccessKeys]);
			commit("removeAllSelectedAccessKeys");
		},
		openDropdown({ commit }, id) {
			commit("setDropdown", id);
		},
		closeDropdown({ commit }) {
			commit("setDropdown", null);
		},
		openAccessTableDropdown({ commit }) {
			commit("setAccessTableDropdown", true);
		},
		closeAccessTableDropdown({ commit }) {
			commit("setAccessTableDropdown", null);
		},
		sort({ commit, state, dispatch }, heading) {
			const flip = (orderBy) => (orderBy === "asc" ? "desc" : "asc");

			const order =
				state.headingSorted === heading ? flip(state.orderBy) : "asc";

			commit("setOrderBy", order);
			commit("setHeadingSorted", heading);
			commit("updateSatelliteRequestObject");
			dispatch("list");
		},
		closeAllInteractions({ state, dispatch }) {
			if (state.openedDropdown) {
				dispatch("closeDropdown");
			}

			if (state.accessTableDropdown) {
				dispatch("closeAccessTableDropdown");
			}
		}
	}
};
