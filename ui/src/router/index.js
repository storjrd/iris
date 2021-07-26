import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";
import Dashboard from "../views/Dashboard";

import Login from "../components/Login";
import SignUp from "../components/SignUp";

import BucketCreation from "../components/BucketCreation";
import BucketList from "../components/BucketList";
import Unlock from "../components/Unlock";
import Browse from "../components/Browse";
import Backup from "../components/Backup";
import Usage from "../components/Usage";
import Access from "../components/Access";
import UpgradeForm from "../components/UpgradeForm";

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
				name: "backup",
				path: "backup",
				component: Backup
			},
			{
				name: "usage",
				path: "usage",
				component: Usage
			},
			{
				name: "access",
				path: "access",
				component: Access
			},
			{
				name: "upgrade-form",
				path: "usage/upgrade-form",
				component: UpgradeForm
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
