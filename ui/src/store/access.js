import * as R from "ramda";

export default {
  namespaced: true,
  state: () => ({
    accessKeys: [],
    headingSorted: "name",
    orderBy: "asc",
    openedDropdown: null,
    accessCreationModal: null,
    accessKeysBeingDeleted: []
  }),
  getters: {
    sortedAccessKeys(state) {
      // temporary for testing purposes
      return state.accessKeys;

      const duration = (a, b) => {
        if (a === "forever") {
          return 1;
        } else if (b === "forever") {
          return -1;
        }

        // find difference between dates and return accordingly
      };

      // key-specific sort cases
      const fns = {
        name: (a, b) => a.name.localeCompare(b.name),
        duration: duration,
        dateCreated: (a, b) => new Date(a.dateCreated) - new Date(b.dateCreated)
      };

      // sort by appropriate function
			const sortedKeys = R.sort(fns[state.headingSorted], state.accessKeys);

      // reverse if descending order
			const orderedAccessKeys =
      state.orderBy === "asc" ? sortedKeys : R.reverse(sortedKeys);

      return orderedAccessKeys;
    }
  },
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
      state.accessKeysBeingDeleted = [...state.accessKeysBeingDeleted, ...keys];
    },
    removeAccessKeyBeingDeleted(state, key) {
      state.accessKeysBeingDeleted = state.accessKeysBeingDeleted.filter((accessKey) => accessKey !== key);
    }
  },
  actions: {
    async list({ commit }) {
      // fetch access keys

      const demoAccessObjects = [
        {
          name: "FoodApp",
          permissions: "Read, Write",
          duration: "Forever",
          buckets: "All",
          dateCreated: "07/06/2021 12:27PM EST"
        },
        {
          name: "Iris",
          permissions: "Read, Write, List, Delete",
          duration: "06/21/2021 - 08/01/2021",
          buckets: "5",
          dateCreated: "06/21/2021 10:04AM EST"
        },
        {
          name: "NFT Kitties",
          permissions: "Read, Write, List",
          duration: "05/14/2021 - 09/22/2021",
          buckets: "2",
          dateCreated: "05/14/2021 3:11PM EST"
        },
        {
          name: "Todo iOS App",
          permissions: "List, Delete",
          duration: "04/29/2021 - 08/05/2021",
          buckets: "8",
          dateCreated: "04/29/2021 8:02AM EST"
        },
        {
          name: "Watermelon Picker",
          permissions: "Write",
          duration: "Forever",
          buckets: "All",
          dateCreated: "03/19/2021 7:12PM EST"
        },
        {
          name: "Accounting",
          permissions: "Read, Write, Delete",
          duration: "Forever",
          buckets: "13",
          dateCreated: "01/30/2021 11:20PM EST"
        }
      ];

      commit("updateAccessKeys", demoAccessObjects);
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
    sort({ commit, state }, heading) {
      const flip = (orderBy) => (orderBy === "asc" ? "desc" : "asc");

      const order = state.headingSorted === heading ?
        flip(state.orderBy) :
        "asc";

      commit("setOrderBy", order);
      commit("setHeadingSorted", heading);
    }
  }
}
