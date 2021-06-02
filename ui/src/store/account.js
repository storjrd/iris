import {
	register,
	login,
	createProject,
	createApiKey,
	getProjects
} from "../lib/satellite";

export default {
	namespaced: true,

	state: () => {
		const sessionJson = null; // localStorage.getItem("session");

		if (typeof sessionJson === "string") {
			const cachedLogin = JSON.parse(sessionJson);

			return cachedLogin;
		}

		return {
			email: "",
			token: null,
			apiKey: null,
			projectId: null
		};
	},
	mutations: {
		setSession(state, { email, token, apiKey, projectId }) {
			console.log("setSession", { email, token, apiKey, projectId });

			state.email = email;
			state.token = token;
			state.apiKey = apiKey;
			state.projectId = projectId;

			/*
			localStorage.setItem(
				"session",
				JSON.stringify({
					email,
					token,
					apiKey,
					projectId
				})
			);
			*/
		},

		setEmail(state, { email }) {
			state.email = email;
		}
	},
	actions: {
		async signUp({ commit, dispatch }, { email, password }) {
			await register({
				fullName: "Iris User",
				shortName: "Iris User",
				email,
				password
			});

			return dispatch("login", {
				email,
				password
			});
		},

		async login({ commit }, { email, password }) {
			const { token } = await login({
				email,
				password
			});

			// find and create Iris project
			const projectName = "Iris";

			const { myProjects } = await getProjects({ token });

			let projectId = myProjects.find(
				(project) => project.name === projectName
			).id;

			// create new project if it doesn't exist
			if (typeof projectId !== "string") {
				console.log("couldn't find Iris project, creating new one");

				const { id } = await createProject({
					token,
					name: projectName,
					description: "Your buckets created and managed by Iris."
				});

				projectId = id;
			}

			// create new api key
			const { key } = await createApiKey({
				token,
				projectId,
				name: `${projectName} Key ${Date.now()}`
			});

			commit("setSession", {
				email,
				token,
				projectId,
				apiKey: key
			});
		},

		async logout({ commit }) {
			commit("setSession", {
				email: null,
				token: null,
				projectId: null,
				apiKey: null
			});
		}
	},
	getters: {
		isLoggedIn: (state) => typeof state.token === "string"
	}
};
