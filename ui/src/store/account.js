import { register, login, createApiKey } from "../lib/satellite";

export default {
    state: () => ({
        email: null,
        token: null,
        apiKey: null
    }),
    mutations: {
        setSession(state, { email, token }) {
            state.email = email;
            state.token = token;
        }
    },
    actions: {
        async signUp({ commit }, { email, password }) {
            const { token } = register({
                fullName: "Iris User",
                shortName: "Iris User",
                email,
                password
            });

            commit("setSession", {
                email,
                token
            });
        },

        async login({ commit }, { email, password }) {
            const { token } = login({
                email,
                password
            });

            commit("setSession", {
                email,
                token
            });
        }
    },
    getters: {
        isLoggedIn: state => typeof state.token === "string"
    }
}