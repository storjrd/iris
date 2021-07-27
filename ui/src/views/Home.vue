<style>
h1,
h2,
h3,
h4,
h5,
h6 {
	font-family: Inter;
	font-weight: 700;
}
p,
span {
	line-height: 1.7;
	font-weight: 300;
}
a {
	font-weight: 400;
}
.logo {
	width: 100%;
	max-width: 140px;
}

.navbar {
	z-index: 2;
}

@media only screen and (max-width: 991px) {
	.navbar {
		text-align: center;
	}
}

@media only screen and (max-width: 991px) {
	.navbar-collapse {
		margin-top: 20px;
	}
}

.navbar-nav .login-link-mobile {
	display: none;
}

@media only screen and (max-width: 991px) {
	.navbar-nav {
		text-align: left;
	}

	.navbar-nav .blog-link {
		margin-top: 12px;
	}

	.navbar-nav .login-link {
		display: none;
	}

	.navbar-nav .login-link-mobile {
		display: block;
		text-align: center;
		margin-top: 10px;
	}
}

.navlink {
	font-style: normal;
	font-weight: 500;
	font-size: 16px;
	line-height: 20px;
	color: #fff;
	cursor: pointer;
}

.navlink a {
	color: #fff;
}

.navlink:hover {
	color: #fff;
	text-decoration: underline;
}

@media only screen and (max-width: 991px) {
	button.navlink {
		padding: 0;
		margin: 5px 0 20px;
	}
}

.dropdown-item {
	margin: 10px 0;
	font-size: 16px;
}

.dropdown-item:hover {
	background: 0 0;
}

.dropdown-item:hover img {
	display: block;
}

.dropdown-menu {
	font-size: 16px;
	width: 300px;
}

@media only screen and (max-width: 991px) {
	.dropdown-menu {
		margin: 0 auto;
		margin-bottom: 12px;
	}
}

.dropdown-menu .dropdown-divider {
	width: 85%;
	margin: 10px auto;
	opacity: 0.5;
	border-top: 1px solid #93a1af;
	margin-left: 24px;
}

@media only screen and (max-width: 991px) {
	.dropdown-menu .dropdown-divider {
		width: 43%;
		float: left;
	}
}

.dropdown-menu img {
	float: right;
	position: relative;
	top: 7px;
	display: none;
}

.dropdown-menu .dropdown-header {
	color: #93a1af;
	font-weight: 700;
	text-transform: uppercase;
	letter-spacing: 1px;
	font-size: 12px;
	line-height: 16px;
	margin-top: 8px;
}

.dropdown-menu .dropdown-header:hover {
	color: #93a1af;
}

.dropdown-menu .dropdown-item {
	font-size: 14px;
}

.dropdown-menu .dropdown-item:hover {
	color: #2582ff;
	font-weight: 700;
}

.navbar-toggler-icon {
	background-image: url(data:image/svg+xml;
 charset=utf8;
	base64,
	PHN2ZyB2aWV3Qm94PSIwIDAgMzIgMzIiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggc3Ryb2tlPSJyZ2JhKDI1NSwyNTUsMjU1LCAxKSIgc3Ryb2tlLXdpZHRoPSIyIiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1taXRlcmxpbWl0PSIxMCIgZD0iTTQgOGgyNE00IDE2aDI0TTQgMjRoMjQiLz48L3N2Zz4=) !important;
}

.navbar-toggler {
	background-color: transparent !important;
	visibility: visible;
}

@media only screen and (max-width: 991px) {
	.navbar-toggler {
		margin-bottom: 17px;
	}
}

.navbar-toggler .exit-icon {
	visibility: hidden;
	display: none;
}

.navbar-toggler[aria-expanded="true"] {
	visibility: hidden;
}

.navbar-toggler[aria-expanded="true"] .exit-icon {
	visibility: visible;
	display: block;
	position: absolute;
	right: 17px;
	top: 10px;
	width: 18px;
}

.navbar-toggler.bg-light:focus,
.navbar-toggler.bg-light:hover {
	background-color: transparent !important;
}

@media only screen and (max-width: 991px) {
	.dropdown-menu.show {
		display: inline;
	}

	.dropdown-menu.show .dropdown-item {
		color: #fff;
	}

	.dropdown-menu.show .dropdown-header {
		color: #93a1af;
		margin-bottom: 0;
		margin-top: 0;
	}

	.dropdown-menu.show .dropdown-icon {
		display: none;
	}
}

.btn-group {
	display: inline;
	position: relative;
	top: 2px;
}

.nav-join {
	border: 2px solid #fff;
	border-radius: 4px;
	color: #fff;
	font-weight: 700;
	font-size: 14px;
	width: 150px;
}

.nav-join:hover {
	color: #2582ff;
	background: #fff;
}

@media only screen and (max-width: 991px) {
	.nav-join {
		width: 100%;
	}
}

.second-nav.navbar {
	padding-top: 20px;
	padding-bottom: 10px;
}

.second-nav.navbar .navlink {
	color: #fff;
}

.second-nav.navbar .navlink:hover {
	color: #fff;
	text-decoration: underline;
}

.second-nav.navbar .nav-join {
	color: #fff;
	border: 2px solid #fff;
}

.second-nav.navbar .nav-join:hover {
	color: #2582ff;
}

@media only screen and (max-width: 991px) {
	.second-nav.navbar .dropdown-menu .dropdown-item {
		color: #fff;
	}
}

@media only screen and (max-width: 991px) {
	.second-nav.navbar .dropdown-menu.show {
		border: none;
	}
}

.second-nav.navbar .navbar-toggler-icon {
	background-image: url(data:image/svg+xml;
 charset=utf8;
	base64,
	PHN2ZyB2aWV3Qm94PSIwIDAgMzIgMzIiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggc3Ryb2tlPSJyZ2JhKDE0NywxNjEsMTc1LCAxKSIgc3Ryb2tlLXdpZHRoPSIyIiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1taXRlcmxpbWl0PSIxMCIgZD0iTTQgOGgyNE00IDE2aDI0TTQgMjRoMjQiLz48L3N2Zz4=) !important;
}

.second-nav.navbar .navbar-toggler[aria-expanded="true"] {
	visibility: visible;
}

.second-nav.navbar .logo {
	margin-top: -14px;
	height: 42px;
}
.input {
	width: 100%;
	height: 48px;
	border: 1px solid rgba(56, 75, 101, 0.4);
	box-sizing: border-box;
	border-radius: 6px;
	font-family: Inter;
	font-style: normal;
	font-weight: normal;
	font-size: 16px;
	color: #000;
	padding: 10px 15px;
	margin-bottom: 16px;
	padding-left: 45px !important;
}
.input:focus {
	color: #000;
}
.email {
	background: url("~@/assets/tar-ico-email.svg") no-repeat 16px 16px;
	padding-left: 48px;
}
.password {
	background: url("~@/assets/tar-ico-lock.svg") no-repeat 16px 16px;
	background-position-x: 20px;
	padding-left: 48px;
}
::-webkit-input-placeholder {
	/* WebKit, Blink, Edge */
	color: #afb7c1;
}
:-moz-placeholder {
	/* Mozilla Firefox 4 to 18 */
	color: #afb7c1;
	opacity: 1;
}
::-moz-placeholder {
	/* Mozilla Firefox 19+ */
	color: #afb7c1;
	opacity: 1;
}
:-ms-input-placeholder {
	/* Internet Explorer 10-11 */
	color: #afb7c1;
}
::-ms-input-placeholder {
	/* Microsoft Edge */
	color: #afb7c1;
}
::placeholder {
	/* Most modern browsers support this now. */
	color: #afb7c1;
}
.video {
	box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.15);
	border: 1px solid #adb4be;
	border-radius: 6px;
}
input {
	height: 48px;
	font-size: 16px;
	padding: 10px 15px;
	border: 1px solid #adb4be;
	color: #474e59;
	border-radius: 6px;
	cursor: text;
}
.button {
	padding: 14px 0;
	background-color: #376fff;
	border-radius: 6px;
	font-weight: bold;
	font-size: 16px;
	line-height: 19px;
	text-align: center;
	text-decoration: none;
	color: #ffffff;
	margin-bottom: 16px;
	transition: all 100ms ease-in-out;
}
.button-no-bg {
	padding: 14px 0;
	border-radius: 6px;
	font-weight: bold;
	font-size: 16px;
	line-height: 19px;
	text-align: center;
	text-decoration: none;
	color: #ffffff;
	margin-bottom: 16px;
	transition: all 100ms ease-in-out;
}
.button:hover {
	background: #0059d0;
}
.btn {
	font-weight: 700;
	border-radius: 50em;
	line-height: 2.4;
}
.btn-primary {
	background: #376fff;
}
.btn-light {
	background: #e6e9ef;
}
.btn-success {
	background: #00a858;
}
.btn-primary.disabled,
.btn-primary:disabled {
	color: #fff;
	background-color: #001030;
	border-color: #001030;
}
.custom-select {
	width: 100%;
	height: 48px;
	border: 1px solid rgba(56, 75, 101, 0.4);
	box-sizing: border-box;
	border-radius: 6px;
	font-family: Inter;
	font-style: normal;
	font-weight: normal;
	font-size: 16px;
	color: rgba(27, 37, 51, 0.8);
	padding: 10px 15px;
	margin-bottom: 16px;
}
.how {
	background: #fff;
	padding: 40px 0;
}
.keys .title {
	font-weight: bold;
	font-size: 29px;
	line-height: 43px;
	text-align: center;
	color: #252525;
}
.keys .label {
	font-weight: bold;
	font-size: 16px;
	color: #354049;
}
.keys .copy {
	background: #0059d0;
	border: 1px solid #0059d0;
	border-radius: 4px;
	font-weight: bold;
	font-size: 14px;
	text-align: center;
	color: #fff;
	cursor: pointer;
	position: absolute;
	bottom: 8px;
	right: 26px;
	padding: 4px 14px;
	line-height: 1.6;
	transition: all 100ms ease-in-out;
}
.keys .copy:hover {
	background: #2582ff;
	color: #fff;
}
.keys a {
	font-weight: bold;
}
.keys input {
	height: 48px;
	font-size: 16px;
	padding: 10px 15px;
	border: 1px solid #adb4be;
	color: #474e59;
	border-radius: 6px;
	cursor: text;
}
.form-control:disabled,
.form-control[readonly] {
	background-color: #f7f8fb;
}
.keys .alert-warning {
	color: #000;
	background-color: #fff9f6;
	border-color: #f9a482;
}
.spinner-border {
	color: #376fff;
}
.nav-tabs {
	border-bottom: none;
}
.nav-item svg {
	margin-top: -2px;
	margin-right: 3px;
}
.nav-tabs .nav-link {
	background: #eff1f5;
	color: #000;
	border-color: #dde1e6;
	padding: 12px;
	border-bottom: 2px solid #dde1e6;
	min-height: 100%;
	transition: all 100ms ease-in-out;
}
.nav-tabs .nav-link.active {
	font-weight: bold;
	color: #000;
}
.nav-tabs .nav-link:hover {
	color: #376fff;
	background: #f8f9fb;
}
.nav-tabs .nav-link.active:hover {
	background: #fff;
	color: #000;
}
.card-top-flat {
	border-top-left-radius: 0;
	border-top-right-radius: 0;
}

@media (max-width: 575.98px) {
	.nav-item {
		text-align: center;
		padding: 0;
	}
	.nav-item .nav-link .nav-icon {
		display: block;
		margin: 0 auto;
	}
}
</style>

<template>
	<div>
		<div class="container">
			<div class="row pb-3 pt-4">
				<div class="col">
					<nav class="navbar navbar-expand-sm second-nav p-0">
						<a class="navbar-brand" href="javascript:null"
							><img
								class="logo"
								alt="Storj logo"
								src="@/assets/storj-logo.svg"
						/></a>
						<ul class="navbar-nav ml-auto">
							<li class="nav-item mr-3 d-none d-sm-block">
								<a
									class="nav-link"
									href="https://storj.io/blog"
									target="_blank"
								>
									<p class="navlink">Blog</p>
								</a>
							</li>
							<li class="nav-item mr-3 d-none d-sm-block">
								<a
									class="nav-link"
									href="https://forum.storj.io/"
									target="_blank"
								>
									<p class="navlink">Forum</p>
								</a>
							</li>

							<li v-if="isLoggedIn" class="nav-item active">
								<button
									class="btn my-2 my-sm-0 nav-join"
									v-on:click="logout"
								>
									Logout
								</button>
							</li>

							<!--<li class="nav-item active login-link-mobile">
								<a class="nav-link" href="https://tardigrade.io/login">
									<p class="navlink">Sign In <span class="sr-only">(current)</span></p>
								</a>
							</li>-->
						</ul>
						<!--</div>-->
					</nav>
				</div>
			</div>
		</div>

		<div>
			<div class="container">
				<div class="row">
					<div class="col-sm-12 col-md-6 col-lg-5 pt-5 mt-0 mt-lg-5">
						<div>
							<h1 class="hero-title text-white">
								Decentralized Compute is here
							</h1>

							<p class="description my-4 text-white">
								Get safer, faster, decentralized object storage
								at a fraction of the cost. Experience it for
								yourself and start with 1 TB free when you sign
								up.
							</p>

							<img
								src="@/assets/mascot.png"
								alt="Mascot"
								class="w-100 mt-4 mascot"
							/>
						</div>
					</div>

					<div class="col-sm-12 col-md-6 col-lg-6 offset-lg-1">
						<div class="card border-0 p-4 p-lg-5 mb-5">
							<div
								v-if="error"
								class="alert alert-danger"
								role="alert"
							>
								{{ error }}
							</div>

							<!-- sign up / login -->
							<router-view></router-view>
						</div>
					</div>
				</div>
			</div>

			<div class="how">
				<div class="container">
					<div class="row py-5">
						<div class="col-sm-12 col-md-5 pb-5">
							<div
								class="
									video
									embed-responsive embed-responsive-16by9
									mb-5
								"
							>
								<iframe
									class="embed-responsive-item"
									src="https://www.youtube.com/embed/BpMbr1U6Z8I"
									frameborder="0"
									allowfullscreen
								></iframe>
							</div>
							<h2 class="mb-4">How It Works</h2>
							<p>
								This app is powered by the Tardigrade
								decentralized cloud service, the world’s first
								decentralized cloud object storage service. The
								Tardigrade cloud storage platform utilizes spare
								disk drive space and bandwidth shared by its
								community members worldwide.
								<br /><br />Tardigrade splits files into
								encrypted pieces, then distributes them on a
								global network, and recompiles each piece on
								download. This means your data isn’t being
								stored in a vulnerable data center; it’s
								securely stored all over the world.
							</p>
						</div>

						<div class="col-sm-12 col-md-6 offset-md-1">
							<div class="row mb-5">
								<div class="col-1">
									<img
										src="@/assets/tar-ico-secure.svg"
										alt="100% Secure"
										height="32"
										width="32"
									/>
								</div>
								<div class="col-11 px-4">
									<h5 class="mb-3">100% Secure</h5>
									<p>
										Store your data with confidence, knowing
										Storj uses end-to-end encryption and
										decentralization to offer unparalled
										security.
									</p>
								</div>
							</div>

							<div class="row mb-5">
								<div class="col-1">
									<img
										src="@/assets/tar-ico-durable.svg"
										alt="Remarkably Durable"
										height="32"
										width="32"
									/>
								</div>
								<div class="col-11 px-4">
									<h5 class="mb-3">Remarkably Durable</h5>
									<p>
										Don't worry about downtime and outages.
										Files on Storj are split into pieces and
										distributed on a global network, which
										is why we've never lost a file.
									</p>
								</div>
							</div>

							<div class="row mb-5">
								<div class="col-1">
									<img
										src="@/assets/tar-ico-money.svg"
										alt="Predictable Pricing"
										height="32"
										width="32"
									/>
								</div>
								<div class="col-11 px-4">
									<h5 class="mb-3">Predictable Pricing</h5>
									<p>
										Storj is a fraction of the cost of other
										storage solutions, and pricing is
										predictable so you don't run into
										surprises when looking at your bill.
									</p>
								</div>
							</div>

							<div class="row mb-5">
								<div class="col-1">
									<img
										src="@/assets/tar-ico-global.svg"
										alt="Globally Distributed"
										height="32"
										width="32"
									/>
								</div>
								<div class="col-11 px-4">
									<h5 class="mb-3">Globally Distributed</h5>
									<p>
										Most storage services run on data
										centers. Storj never puts your file in
										one place, so downtime and breaches are
										a thing of the past.
									</p>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>

			<main-footer></main-footer>
		</div>
	</div>
</template>

<script>
import MainFooter from "../components/MainFooter";
import { ref, watch, computed } from "vue";

export default {
	name: "Home",
	components: {
		MainFooter
	},
	setup() {
		const apiKey = ref("");
		const email = ref("");
		const password = ref("");
		const termsAndConditions = ref(false);
		const message = ref("");
		const showLogin = ref(false);
		const destination = ref("/app");

		const signUp = async () => {};

		const login = async () => {};

		const error = computed(() => {
			//return this.$store.state.userError;
		});

		const isLoggedIn = computed(() => {
			//return this.$store.getters.isLoggedIn;
		});

		watch(isLoggedIn, (loggedIn) => {
			if (loggedIn === true) {
				this.$router.push({ path: destination });
			}
		});

		// redirection
		//if (this.$route.query.app) this.destination = "/app/apps";
		//else this.destination = "/app/browser";

		return {
			apiKey,
			email,
			password,
			termsAndConditions,
			message,
			showLogin,
			error,
			isLoggedIn,
			signUp,
			login
		};
	}
};
</script>
