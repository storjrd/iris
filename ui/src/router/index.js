import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";
import Dashboard from "../views/Dashboard";

import Login from "../components/Login";
import SignUp from "../components/SignUp";

import BucketCreation from "../components/BucketCreation";
import BucketAccess from "../components/BucketAccess";
import BucketList from "../components/BucketList";
import Unlock from "../components/Unlock";
import Browse from "../components/Browse";

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
				path: "create",
				component: BucketCreation
			},
			{
				path: "access",
				component: BucketAccess
			},
			{
				path: "buckets",
				component: BucketList
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
			},
		]
	}
];

const router = createRouter({
	history: createWebHistory(process.env.BASE_URL),
	routes
});

export default router;
