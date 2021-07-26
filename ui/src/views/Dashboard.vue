<style scoped>
.white-background {
	background-color: white;
}
.logo {
	margin-top: -4px;
	margin-bottom: 20px;
	margin-left: 12px;
}
.nav-item {
	margin-top: 10px;
	margin-bottom: 10px;
	border-radius: 46px;
	padding: 10px;
	width: 100%;
}
.nav-item.active {
	background: #fff;
	font-weight: 700;
}
.nav-item.active .nav-link {
	color: #001030;
}
.nav-link {
	color: #c0c4cf;
}
.nav-link:hover {
	color: #fff;
}
.nav-icon {
	width: 24px;
	height: 24px;
	margin-right: 4px;
	margin-left: 14px;
	display: inline-block;
}
.community-icon {
	position: relative;
	left: -3px;
}
.nav-icon svg path {
	fill: #768394;
}
.nav-item.active .nav-link .nav-icon svg path {
	fill: #376fff;
}
.nav-item:hover .nav-link .nav-icon svg path {
	fill: #fff;
}
.nav-item.active:hover .nav-link .nav-icon svg path {
	fill: #376fff;
}
.position-sticky {
	top: 0;
}
.navbar-nav.bottom {
	position: sticky;
	bottom: 0;
	margin-top: 430px;
}
.project-title {
	color: white;
	font-weight: lighter;
}
.project-heading {
	color: white;
	font-weight: bold;
	white-space: nowrap;
	font-size: 0.9rem;
}
</style>

<template>
	<div class="container-fluid">
		<div class="row">
			<div class="col-4 col-md-3 col-lg-3 col-xl-2 p-3">
				<div class="position-sticky">
					<a class="navbar-brand mt-4" href="javascript:null">
						<img
							class="logo"
							alt="Storj logo"
							src="@/assets/storj-logo.svg"
						/>
					</a>

					<ul class="navbar-nav top">
						<li class="nav-item">
							<div class="project-title">Project</div>
							<a
								class="
									nav-link
									project-heading
									d-flex
									justify-content-between
								"
								href="#"
							>
								{{ "My Storage Project" }}
								<span class="align-middle">
									<svg
										xmlns="http://www.w3.org/2000/svg"
										width="16"
										height="16"
										fill="currentColor"
										class="bi bi-chevron-down"
										viewBox="0 0 16 16"
									>
										<path
											fill-rule="evenodd"
											d="M1.646 4.646a.5.5 0 0 1 .708 0L8 10.293l5.646-5.647a.5.5 0 0 1 .708.708l-6 6a.5.5 0 0 1-.708 0l-6-6a.5.5 0 0 1 0-.708z"
										/>
									</svg>
								</span>
							</a>
						</li>
						<li
							class="nav-item"
							v-bind:class="{ active: bucketsPath }"
						>
							<router-link class="nav-link" to="/app/buckets">
								<div class="nav-icon">
									<svg
										width="16"
										height="18"
										viewBox="0 0 16 18"
										fill="none"
										xmlns="http://www.w3.org/2000/svg"
									>
										<path
											fill-rule="evenodd"
											clip-rule="evenodd"
											d="M15.42 4.55488L14.3732 16.0185L14.3704 16.0186C14.2668 17.1183 11.3267 18 7.7149 18C4.10314 18 1.16304 17.1183 1.05939 16.0186L1.05656 16.0185L0.00976562 4.55488C0.0472251 4.58465 0.0858874 4.61398 0.125716 4.64285C0.584003 4.97514 1.21629 5.26217 1.9919 5.50003C3.49783 5.96186 5.5026 6.22504 7.63856 6.23117L7.7149 6.23128C9.8794 6.23128 11.914 5.96736 13.4379 5.50003C14.2135 5.26217 14.8458 4.97514 15.3041 4.64285C15.3439 4.61398 15.3826 4.58466 15.42 4.55488ZM7.7149 0C11.7162 0 14.9599 0.994665 14.9599 2.22165C14.9599 3.44863 11.7162 4.44329 7.7149 4.44329C3.71361 4.44329 0.469924 3.44863 0.469924 2.22165C0.469924 0.994665 3.71361 0 7.7149 0Z"
											fill="#768394"
										/>
									</svg>
								</div>
								Buckets
							</router-link>
						</li>
						<li
							class="nav-item"
							v-bind:class="{ active: backupPath }"
						>
							<router-link class="nav-link" to="/app/backup">
								<div class="nav-icon">
									<svg
										xmlns="http://www.w3.org/2000/svg"
										width="16"
										height="16"
										fill="currentColor"
										class="bi bi-clock-fill"
										viewBox="0 0 16 16"
									>
										<path
											d="M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0zM8 3.5a.5.5 0 0 0-1 0V9a.5.5 0 0 0 .252.434l3.5 2a.5.5 0 0 0 .496-.868L8 8.71V3.5z"
										/>
									</svg>
								</div>
								Backup
							</router-link>
						</li>
						<li
							class="nav-item"
							v-bind:class="{ active: accessPath }"
						>
							<router-link class="nav-link" to="/app/access">
								<div class="nav-icon">
									<svg
										class="svg"
										width="19"
										height="19"
										viewBox="0 0 19 19"
										fill="none"
										xmlns="http://www.w3.org/2000/svg"
									>
										<path
											class="navigation-svg-path"
											opacity="0.8"
											fill-rule="evenodd"
											clip-rule="evenodd"
											d="M19 0L18.8777 3.30167L17.3804 3.33154L17.3114 4.86798L15.548 5.164L15.5552 6.62419L14.0737 6.63827L13.8958 8.28354L12.3349 9.84422L13.2397 13.2214C13.3996 13.8182 13.229 14.4549 12.7921 14.8918L9.19044 18.4934C8.75356 18.9303 8.1168 19.1009 7.52001 18.941L2.60009 17.6227C2.00331 17.4628 1.53716 16.9967 1.37725 16.3999L0.0589658 11.48C-0.100943 10.8832 0.0696777 10.2464 0.506556 9.80956L4.10819 6.20793C4.54507 5.77105 5.18183 5.60043 5.77862 5.76034L9.15552 6.66484L15.6983 0.122284L19 0ZM6.52703 12.473C5.78414 11.7301 4.57967 11.7301 3.83678 12.473C3.09389 13.2159 3.09389 14.4203 3.83678 15.1632C4.57967 15.9061 5.78414 15.9061 6.52703 15.1632C7.26992 14.4203 7.26992 13.2159 6.52703 12.473Z"
											fill="#768394"
										/>
									</svg>
								</div>
								Keys
							</router-link>
						</li>
					</ul>

					<ul class="navbar-nav bottom pb-4">
						<li
							class="nav-item"
							v-bind:class="{ active: usagePath }"
						>
							<router-link class="nav-link" to="/app/usage">
								<div class="nav-icon">
									<svg
										width="19"
										height="14"
										viewBox="0 0 19 14"
										fill="none"
										xmlns="http://www.w3.org/2000/svg"
									>
										<path
											d="M4 0C1.79086 0 0 1.79086 0 4V7.77285L0.232979 7.56317C2.03668 5.93984 4.75516 5.87583 6.63327 7.41246C7.79445 8.36252 9.48052 8.30255 10.5713 7.2724L13.3261 4.67062C14.7828 3.29486 16.7699 2.81912 18.5933 3.23348C18.2357 1.39103 16.6136 0 14.6667 0H4Z"
											fill="#768394"
										/>
										<path
											d="M18.6667 5.26758C17.3261 4.76332 15.7593 5.0291 14.6522 6.07472L11.8974 8.6765C10.1006 10.3735 7.3231 10.4723 5.41028 8.90723C4.27017 7.97441 2.61991 8.01327 1.52497 8.99872L0.0157177 10.357C0.196332 12.3989 1.91119 14 4 14H14.6667C16.8758 14 18.6667 12.2091 18.6667 10V5.26758Z"
											fill="#768394"
										/>
									</svg>
								</div>
								Plan & Usage
							</router-link>
						</li>
						<li class="nav-item">
							<a
								class="nav-link"
								href="https://forum.storj.io"
								target="_blank"
							>
								<div class="nav-icon community-icon">
									<svg
										width="23"
										height="17"
										viewBox="0 0 29 17"
										fill="none"
										xmlns="http://www.w3.org/2000/svg"
									>
										<g clip-path="url(#clip0)">
											<path
												fill-rule="evenodd"
												clip-rule="evenodd"
												d="M9.06257 10.3421C7.08318 10.3421 5.27165 11.0609 3.87491 12.2519C2.60132 13.3378 1.67231 14.8162 1.27557 16.5H12.2809L14.5 12.474C13.0728 11.1509 11.1621 10.3421 9.06257 10.3421ZM9.06257 0.5C8.0888 0.5 7.20722 0.894697 6.56908 1.53283C5.93094 2.17097 5.53625 3.05255 5.53625 4.02632C5.53625 5.00008 5.93094 5.88166 6.56908 6.5198C7.20722 7.15794 8.0888 7.55264 9.06257 7.55264C10.0363 7.55264 10.9179 7.15794 11.5561 6.5198C12.1942 5.88166 12.5889 5.00008 12.5889 4.02632C12.5889 3.05255 12.1942 2.17097 11.5561 1.53283C10.9179 0.894697 10.0363 0.5 9.06257 0.5Z"
												fill="#768394"
											/>
											<path
												fill-rule="evenodd"
												clip-rule="evenodd"
												d="M21.7279 10.3421C19.7485 10.3421 17.937 11.0609 16.5403 12.2519C15.2667 13.3378 14.3377 14.8162 13.9409 16.5H29.5149C29.1182 14.8162 28.1892 13.3378 26.9156 12.2519C25.5188 11.0609 23.7073 10.3421 21.7279 10.3421ZM21.7279 0.5C20.7541 0.5 19.8726 0.894697 19.2344 1.53283C18.5963 2.17097 18.2016 3.05255 18.2016 4.02632C18.2016 5.00008 18.5963 5.88166 19.2344 6.5198C19.8726 7.15794 20.7541 7.55264 21.7279 7.55264C22.7017 7.55264 23.5833 7.15794 24.2214 6.5198C24.8595 5.88166 25.2542 5.00008 25.2542 4.02632C25.2542 3.05255 24.8595 2.17097 24.2214 1.53283C23.5833 0.894697 22.7017 0.5 21.7279 0.5Z"
												fill="#768394"
											/>
										</g>
										<defs>
											<clipPath id="clip0">
												<rect
													width="27"
													height="17"
													fill="white"
												/>
											</clipPath>
										</defs>
									</svg>
								</div>
								Community
							</a>
						</li>
						<!-- <li v-on:click="logout" class="nav-item">
							<a
								class="nav-link d-flex justify-content-between"
								href="#"
							>
								<span>
									<div class="nav-icon">
										<svg
											xmlns="http://www.w3.org/2000/svg"
											width="16"
											height="16"
											fill="currentColor"
											class="bi bi-person-circle"
											viewBox="0 0 16 16"
										>
											<path
												d="M11 6a3 3 0 1 1-6 0 3 3 0 0 1 6 0z"
											/>
											<path
												fill-rule="evenodd"
												d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8zm8-7a7 7 0 0 0-5.468 11.37C3.242 11.226 4.805 10 8 10s4.757 1.225 5.468 2.37A7 7 0 0 0 8 1z"
											/>
										</svg>
									</div>
									Account
								</span>
								<span class="align-middle">
									<svg
										xmlns="http://www.w3.org/2000/svg"
										width="16"
										height="16"
										fill="currentColor"
										class="bi bi-chevron-right"
										viewBox="0 0 16 16"
									>
										<path
											fill-rule="evenodd"
											d="M4.646 1.646a.5.5 0 0 1 .708 0l6 6a.5.5 0 0 1 0 .708l-6 6a.5.5 0 0 1-.708-.708L10.293 8 4.646 2.354a.5.5 0 0 1 0-.708z"
										/>
									</svg>
								</span>
							</a>
						</li> -->
						<li v-on:click="logout" class="nav-item">
							<a class="nav-link" href="#">
								<div class="nav-icon">
									<svg
										xmlns="http://www.w3.org/2000/svg"
										width="18"
										height="16"
										fill="none"
										class="bi bi-lock-fill"
										viewBox="0 0 18 16"
									>
										<path
											d="M8 1a2 2 0 0 1 2 2v4H6V3a2 2 0 0 1 2-2zm3 6V3a3 3 0 0 0-6 0v4a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z"
										/>
									</svg>
								</div>
								Logout
							</a>
						</li>
					</ul>
				</div>
			</div>
			<div
				class="
					col-8 col-md-9 col-lg-9 col-xl-10
					p-4 p-lg-5
					white-background
				"
			>
				<router-view></router-view>
			</div>
		</div>
	</div>
</template>

<script>
import AppFooter from "../components/AppFooter.vue";

export default {
	components: {
		AppFooter
	},
	computed: {
		backupPath() {
			return this.$route.path === "/app/backup";
		},
		bucketsPath() {
			return this.$route.path === "/app/buckets";
		},
		usagePath() {
			return (
				this.$route.path === "/app/usage" ||
				this.$route.path === "/app/usage/upgrade-form"
			);
		},
		accessPath() {
			return this.$route.path === "/app/access";
		}
	},
	methods: {
		logout() {
			this.$store.dispatch("account/logout");
		}
	}
};
</script>
