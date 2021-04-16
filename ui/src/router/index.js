import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";
import Dashboard from "../views/Dashboard";
import BucketCreation from "../components/BucketCreation";
import BucketAccess from "../components/BucketAccess";
import BucketList from "../components/BucketList";

const routes = [
	{
		path: "/",
		name: "Home",
		component: Home
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
			}
		]
	}
];

const router = createRouter({
	history: createWebHistory(process.env.BASE_URL),
	routes
});

export default router;
