import { getAccessKeys } from "../lib/satellite";

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
		removeAccessKeyBeingDeleted(state, key) {
			state.accessKeysBeingDeleted = state.accessKeysBeingDeleted.filter(
				(accessKey) => accessKey !== key
			);
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
		async deleteAccessKey({ commit }, key) {
			// find key by name
			// commit("setAccessKeysBeingDeleted", [key]);
			// delete access key
			// commit("removeAccessKeyBeingDeleted", key);
		},
		async deleteSelectedAccessKeys({ dispatch }, keys) {
			await Promise.all(keys.map((key) => dispatch("deleteAccessKey")));
		},
		openDropdown({ commit }, id) {
			commit("setDropdown", id);
		},
		closeDropdown({ commit }) {
			commit("setDropdown", null);
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
		}
	}
};
