<style scoped>
.file-browser {
	min-height: 500px;

	user-select: none;
	-moz-user-select: none;
	-khtml-user-select: none;
	-webkit-user-select: none;
	-o-user-select: none;
}

tbody {
	user-select: none;
}

.table-heading {
	color: #768394;
	border-top: 0;
	border-bottom: 1px solid #dee2e6;
	padding-left: 0;
	cursor: pointer;
}

.path {
	font-size: 18px;
	font-weight: 700;
}

.upload-help {
	font-size: 1.75rem;
	text-align: center;
	margin-top: 1.5rem;
	color: #93a1ae;
	border: 2px dashed #bec4cd;
	border-radius: 10px;
	padding: 80px 20px;
	background: #fafafb;
}

.metric {
	color: #444;
}

.div-responsive {
	min-height: 400px;
}
</style>

<template>
	<div>
		<div class="col-sm-12">
			<div class="card card-top-flat border-0 p-4 p-lg-5">
				<div class="div-responsive" v-cloak>
					<div class="row mb-2">
						<div class="col-sm-12 col-md-4 col-xl-8 mb-3">
							<h2>
								Buckets
								<!-- <span class="trash-can mx-1" v-if="areThereFilesToDelete()" v-on:click="deleteSelected">
								<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-trash" viewBox="0 0 16 16">
									<path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>
									<path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4L4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>
								</svg>
							</span> -->
							</h2>
						</div>
						<div class="col-sm-12 col-md-4 col-xl-1 mb-3"></div>
						<div
							class="col-sm-12 col-md-4 col-xl-3 mb-3 d-flex justify-content-lg-end justify-content-center"
						>
							<input
								ref="fileInput"
								type="file"
								hidden
								multiple
							/>
							<button class="btn btn-primary btn-block w-75" @click="createBucket">
								<svg
									width="22"
									height="20"
									viewBox="0 0 22 20"
									fill="none"
									xmlns="http://www.w3.org/2000/svg"
								>
									<path
										fill-rule="evenodd"
										clip-rule="evenodd"
										d="M11.9414 6.37625C11.9801 6.08967 12 5.79716 12 5.5C12 4.27554 11.6614 3.13015 11.0728 2.15236C11.8914 2.054 12.7826 2 13.7149 2C17.7162 2 20.9599 2.99466 20.9599 4.22165C20.9599 5.44863 17.7162 6.44329 13.7149 6.44329C13.1031 6.44329 12.509 6.42003 11.9414 6.37625ZM6.5 11.9236C8.71728 11.5812 10.5646 10.1182 11.4453 8.13147C12.1494 8.19493 12.8858 8.22901 13.6386 8.23117L13.7149 8.23128C15.8794 8.23128 17.914 7.96736 19.4379 7.50003C20.2135 7.26217 20.8458 6.97514 21.3041 6.64285C21.3439 6.61398 21.3826 6.58466 21.42 6.55488L20.3732 18.0185L20.3704 18.0186C20.2668 19.1183 17.3267 20 13.7149 20C10.1031 20 7.16303 19.1183 7.05939 18.0186L7.05656 18.0185L6.5 11.9236Z"
										fill="white"
									/>
									<circle
										cx="5.5"
										cy="5.5"
										r="5.5"
										fill="white"
									/>
									<path
										d="M7.60876 5.69568H3.39124M5.5 3.58691V7.80444"
										stroke="#0068DC"
										stroke-linecap="round"
										stroke-linejoin="round"
									/>
								</svg>

								New Bucket
							</button>
						</div>
					</div>

					<div>
						<table class="table table-hover">
							<th class="table-heading" scope="col">Name</th>
							<!--<th class="table-heading" scope="col">
								Date Added
							</th>-->
							<th class="table-heading"></th>

							<tbody>
								<bucket
									v-for="bucket in buckets"
									v-bind:bucket="bucket"
									v-bind:key="bucket"
								></bucket>
							</tbody>
						</table>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
import { inject, ref } from "vue";
import S3 from "aws-sdk/clients/s3";

import Bucket from "./Bucket";

import { generateAccess } from "../lib/access";
import { getCredentials } from "../lib/mt";

export default {
	name: "BucketList",
	components: {
		Bucket
	},
	setup() {
		const buckets = ref([]);

		const store = inject("store");

		console.log(store.getters["gateway/rootClient"]);

		buckets.value = [];

		async function listBuckets() {
			buckets.value = await store.dispatch("gateway/getBuckets");
		}

		async function createBucket() {
			const name = prompt("Bucket name");

			await store.dispatch("gateway/createBucket", { name });
			await listBuckets();
		}

		listBuckets();

		return {
			buckets,
			createBucket
		};
	},
	createBucket() {

	}
};
</script>
