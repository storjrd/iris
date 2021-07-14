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
			<th></th>
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
		dateCreatedAscFill: ascFill("dateCreated")
		// durationDescFill: descFill("duration"),
		// durationAscFill: ascFill("duration")
	},
	methods: {
		sortByName: sortBy("name"),
		sortByDateCreated: sortBy("dateCreated")
		// sortByDuration: sortBy("duration")
	}
};
</script>
