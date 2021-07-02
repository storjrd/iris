import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";
import Dashboard from "../views/Dashboard";

import Login from "../components/Login";
import SignUp from "../components/SignUp";

import BucketCreation from "../components/BucketCreation";
import BucketList from "../components/BucketList";
import Unlock from "../components/Unlock";
import Browse from "../components/Browse";
<<<<<<< HEAD
import Backup from "../components/Backup";
=======
import Usage from "../components/Usage";
>>>>>>> usage-page

const routes = [
	{
		path: "/",
		name: "Home",
		component: Home,
		children: [
			{
				path: "login",
				component: Login
			},
			{
				path: "",
				component: SignUp
			}
		]
	},
	{
		path: "/app",
		component: Dashboard,
		children: [
			{
<<<<<<< HEAD
				name: "backup",
				path: "backup",
				component: Backup
=======
				name: "usage",
				path: "usage",
				component: Usage
>>>>>>> usage-page
			},
			{
				name: "buckets",
				path: "buckets",
				component: BucketList
			},
			{
				name: "create",
				path: "buckets/create",
				component: BucketCreation
			},
			{
				name: "unlock",
				path: "buckets/:bucket/unlock",
				component: Unlock
			},
			{
				name: "browse",
				path: "buckets/:bucket/browse/:pathMatch*",
				component: Browse
			}
		]
	}
];

const router = createRouter({
	history: createWebHistory(process.env.BASE_URL),
	routes
});

export default router;
