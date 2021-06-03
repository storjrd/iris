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
	margin-top: 300px;
}
</style>

<template>
	<div class="container-fluid">
		<div class="row">
			<div class="col-4 col-md-3 col-lg-3 col-xl-2 p-3">
				<div class="position-sticky">
					<a class="navbar-brand" href="javascript:null">
						<img
							class="logo"
							alt="Storj logo"
							src="@/assets/storj-logo.svg"
						/>
					</a>

					<ul class="navbar-nav top">
						<li class="nav-item active">
							<a class="nav-link" href="https://storj.io/blog">
								<div class="nav-icon">
									<svg
										width="18"
										height="16"
										viewBox="0 0 18 16"
										fill="none"
										xmlns="http://www.w3.org/2000/svg"
									>
										<path
											d="M0 3C0 1.34315 1.34315 0 3 0H7.875C8.49632 0 9 0.50368 9 1.125V1.125C9 1.74632 9.50368 2.25 10.125 2.25H15C16.6569 2.25 18 3.59315 18 5.25V12.75C18 14.4069 16.6569 15.75 15 15.75H3C1.34315 15.75 0 14.4069 0 12.75V3Z"
											fill="#376FFF"
										/>
									</svg>
								</div>
								Files
							</a>
						</li>
						<li class="nav-item">
							<a class="nav-link" href="">
								<div class="nav-icon">
									<svg
										width="18"
										height="24"
										viewBox="0 0 18 24"
										fill="none"
										xmlns="http://www.w3.org/2000/svg"
									>
										<path
											fill-rule="evenodd"
											clip-rule="evenodd"
											d="M9 23.087C9 23.087 18 13.9706 18 9C18 4.02944 13.9706 0 9 0C4.02944 0 0 4.02944 0 9C0 13.9706 9 23.087 9 23.087ZM9.0001 12.5217C10.9451 12.5217 12.5218 10.945 12.5218 8.99998C12.5218 7.05498 10.9451 5.47825 9.0001 5.47825C7.0551 5.47825 5.47836 7.05498 5.47836 8.99998C5.47836 10.945 7.0551 12.5217 9.0001 12.5217Z"
											fill="#768394"
										/>
									</svg>
								</div>
								Endpoint
							</a>
						</li>
						<li class="nav-item">
							<a class="nav-link" href="">
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
								Usage
							</a>
						</li>
					</ul>

					<ul class="navbar-nav bottom pb-4">
						<li class="nav-item">
							<a class="nav-link" href="" target="_blank">
								<div class="nav-icon">
									<svg
										width="18"
										height="18"
										viewBox="0 0 18 18"
										fill="none"
										xmlns="http://www.w3.org/2000/svg"
									>
										<path
											fill-rule="evenodd"
											clip-rule="evenodd"
											d="M2.18477 6.89616L0.92135 7.11049C0.383646 7.20073 -0.00741218 7.66323 0.000106588 8.20846V9.79149C0.000106588 10.3367 0.391166 10.7992 0.928871 10.8895L2.19229 11.1C2.32014 11.5212 2.4931 11.931 2.69991 12.3258L1.9554 13.3712C1.63954 13.8149 1.69219 14.4203 2.07572 14.8038L3.19625 15.9243C3.40682 16.1349 3.68508 16.2515 3.98213 16.2515C4.21526 16.2515 4.44087 16.18 4.62888 16.0447L5.67045 15.3001C6.06151 15.507 6.47136 15.6762 6.89627 15.8078L7.11059 17.0712C7.20084 17.6089 7.66334 17.9999 8.20857 17.9999H9.79536C10.3406 17.9999 10.8031 17.6089 10.8933 17.0712L11.1039 15.8078C11.5138 15.6799 11.9123 15.5182 12.2921 15.3189L13.3487 16.071C13.5367 16.2063 13.7586 16.2778 13.9917 16.2778C14.2888 16.2778 14.567 16.1612 14.7776 15.9507L15.8981 14.8301C16.2854 14.4466 16.3343 13.8412 16.0184 13.3975L15.2777 12.3672C15.4883 11.9762 15.6612 11.5663 15.7928 11.1452L17.0713 10.9271C17.609 10.8368 18.0001 10.3743 18.0001 9.82909V8.24606C18.0001 7.70083 17.609 7.23833 17.0713 7.14809L15.8229 6.93752C15.6951 6.51262 15.5296 6.10276 15.3228 5.70794L16.0748 4.65133C16.3907 4.20763 16.3381 3.60225 15.9545 3.21871L14.8302 2.09818C14.6234 1.88761 14.3414 1.77104 14.0444 1.77104C13.8112 1.77104 13.5856 1.84248 13.3976 1.97785L12.3636 2.72237C11.9725 2.5118 11.5626 2.33883 11.1415 2.20722L10.9234 0.928763C10.8332 0.391058 10.3707 0 9.82544 0H8.24241C7.69718 0 7.23468 0.391058 7.14443 0.928763L6.93387 2.17714C6.49393 2.30875 6.06903 2.48547 5.66293 2.6998L4.6176 1.95529C4.42959 1.81992 4.20774 1.74848 3.97461 1.74848C3.67756 1.74848 3.39931 1.86505 3.18874 2.07562L2.0682 3.19615C1.6809 3.57968 1.63202 4.18507 1.94788 4.62877L2.69239 5.67034C2.48558 6.0614 2.31637 6.47126 2.18477 6.89616ZM9.01323 6.01634C10.6538 6.01634 11.9838 7.3463 11.9838 8.98688C11.9838 10.6275 10.6538 11.9574 9.01323 11.9574C7.37265 11.9574 6.04269 10.6275 6.04269 8.98688C6.04269 7.3463 7.37265 6.01634 9.01323 6.01634Z"
											fill="#768394"
										/>
									</svg>
								</div>
								Settings
							</a>
						</li>
						<li class="nav-item">
							<a
								class="nav-link"
								href="https://forum.storj.io/"
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
						<li class="nav-item">
							<a
								class="nav-link"
								href="https://support.storj.io/hc/en-us"
								target="_blank"
							>
								<div class="nav-icon">
									<svg
										width="20"
										height="17"
										viewBox="0 0 20 17"
										fill="none"
										xmlns="http://www.w3.org/2000/svg"
									>
										<path
											fill-rule="evenodd"
											clip-rule="evenodd"
											d="M6.5 0C2.91015 0 0 2.91015 0 6.5C0 9.6597 2.25453 12.2928 5.24267 12.8785L5.24264 17L12.5695 13H13.5C17.0899 13 20 10.0899 20 6.5C20 2.91015 17.0899 0 13.5 0H6.5Z"
											fill="#768394"
										/>
									</svg>
								</div>
								Support
							</a>
						</li>
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
	methods: {
		logout() {
			this.$store.dispatch("account/logout");
		}
	}
};
</script>
