<style scoped>
thead th {
	font-weight: normal;
	border-top: none;
	border-bottom: none;
}
.container {
	display: flex;
	-webkit-box-orient: vertical;
	-webkit-box-direction: normal;
	-ms-flex-direction: column;
	flex-direction: column;
	-webkit-box-pack: justify;
	-ms-flex-pack: justify;
	justify-content: space-between;
	height: 17px;
}
.name {
	font-weight: bold;
}
.arrow {
	cursor: pointer;
}
</style>

<template>
	<thead>
		<tr>
			<th v-on:click="sortByName">
				<div class="name d-flex align-items-center">
					NAME
					<div class="container">
						<svg
							width="9"
							height="6"
							fill="none"
							xmlns="http://www.w3.org/2000/svg"
							class="arrow"
						>
							<path
								d="M4.737 0L9 6H0l4.737-6z"
								v-bind:fill="nameAscFill"
								class="arrow-svg-path"
							></path>
						</svg>
						<svg
							width="9"
							height="6"
							fill="none"
							xmlns="http://www.w3.org/2000/svg"
							class="arrow"
						>
							<path
								d="M4.263 6L0 0h9L4.263 6z"
								v-bind:fill="nameDescFill"
								class="arrow-svg-path"
							></path>
						</svg>
					</div>
				</div>
			</th>
			<!-- <th>
        PERMISSIONS
      </th>
      <th class="d-flex align-items-center" v-on:click="sortByDuration">
        DURATION
        <div class="container">
          <svg width="9" height="6" fill="none" xmlns="http://www.w3.org/2000/svg" class="arrow">
            <path d="M4.737 0L9 6H0l4.737-6z" v-bind:fill="durationAscFill" class="arrow-svg-path"></path>
          </svg>
          <svg width="9" height="6" fill="none" xmlns="http://www.w3.org/2000/svg" class="arrow">
            <path d="M4.263 6L0 0h9L4.263 6z" v-bind:fill="durationDescFill" class="arrow-svg-path"></path>
          </svg>
        </div>
      </th>
      <th>
        BUCKETS
      </th> -->
			<th v-on:click="sortByDateCreated">
				<div class="d-flex align-items-center">
					DATE CREATED
					<div class="container">
						<svg
							width="9"
							height="6"
							fill="none"
							xmlns="http://www.w3.org/2000/svg"
							class="arrow"
						>
							<path
								d="M4.737 0L9 6H0l4.737-6z"
								v-bind:fill="dateCreatedAscFill"
								class="arrow-svg-path"
							></path>
						</svg>
						<svg
							width="9"
							height="6"
							fill="none"
							xmlns="http://www.w3.org/2000/svg"
							class="arrow"
						>
							<path
								d="M4.263 6L0 0h9L4.263 6z"
								v-bind:fill="dateCreatedDescFill"
								class="arrow-svg-path"
							></path>
						</svg>
					</div>
				</div>
			</th>
			<th class="table-heading" scope="col">
				<div class="dropleft">
					<a
						class="d-flex justify-content-end"
						v-if="keysToDelete"
						v-on:click="deleteSelectedDropdown"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							class="bi bi-trash"
							viewBox="0 0 16 16"
						>
							<path
								d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"
							/>
							<path
								fill-rule="evenodd"
								d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4L4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"
							/>
						</svg>
					</a>
					<div
						v-if="displayDropdown"
						class="dropdown-menu shadow show"
					>
						<div>
							<p class="deletion-confirmation px-5 pt-3">
								Are you sure?
							</p>
							<div class="d-flex">
								<button
									class="dropdown-item trash p-3 action"
									v-on:click="confirmDeleteSelection"
								>
									<svg
										xmlns="http://www.w3.org/2000/svg"
										width="16"
										height="16"
										fill="red"
										class="bi bi-trash"
										viewBox="0 0 16 16"
									>
										<path
											d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"
										/>
										<path
											fill-rule="evenodd"
											d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4L4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"
										/>
									</svg>
									Yes
								</button>
								<button
									class="dropdown-item p-3 action"
									v-on:click="cancelDeleteSelection"
								>
									<svg
										width="2em"
										height="2em"
										viewBox="0 0 16 16"
										class="bi bi-x mr-1"
										fill="green"
										xmlns="http://www.w3.org/2000/svg"
									>
										<path
											fill-rule="evenodd"
											d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z"
										/>
									</svg>
									No
								</button>
							</div>
						</div>
					</div>
				</div>
			</th>
		</tr>
	</thead>
</template>

<script>
// Computed property creators

const fromAccessStore = (prop) =>
	function () {
		return this.$store.state.access[prop];
	};

const colorFill = (condition) => {
	if (condition) {
		return "#007bff";
	}

	return "#354049";
};

const descFill = (heading) =>
	function () {
		return colorFill(
			this.headingSorted === heading && this.orderBy === "desc"
		);
	};

const ascFill = (heading) =>
	function () {
		return colorFill(
			this.headingSorted === heading && this.orderBy === "asc"
		);
	};

// Method creators
const sortBy = (heading) =>
	function () {
		this.$store.dispatch("access/sort", heading);
	};

export default {
	name: "AccessTableHeader",
	computed: {
		headingSorted: fromAccessStore("headingSorted"),
		orderBy: fromAccessStore("orderBy"),

		nameDescFill: descFill("name"),
		nameAscFill: ascFill("name"),
		dateCreatedDescFill: descFill("dateCreated"),
		dateCreatedAscFill: ascFill("dateCreated"),
		// durationDescFill: descFill("duration"),
		// durationAscFill: ascFill("duration"),

		keysToDelete() {
			return this.$store.state.access.selectedAccessKeys.length > 0;
		},
		displayDropdown() {
			return this.$store.state.access.accessTableDropdown;
		}
	},
	methods: {
		sortByName: sortBy("name"),
		sortByDateCreated: sortBy("dateCreated"),
		// sortByDuration: sortBy("duration"),

		deleteSelectedDropdown(event) {
			event.stopPropagation();
			this.$store.dispatch("access/openAccessTableDropdown");
		},
		confirmDeleteSelection() {
			this.$store.dispatch("access/deleteSelectedAccessKeys");
			this.$store.dispatch("access/closeAccessTableDropdown");
		},
		cancelDeleteSelection() {
			this.$store.dispatch("access/closeAccessTableDropdown");
		}
	}
};
</script>
