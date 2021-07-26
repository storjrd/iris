<style scoped>
.form-check {
	display: flex;
	align-items: center;
}
.form-check-input {
	margin-top: 0px;
	cursor: pointer;
}
label {
	font-weight: bold;
	margin-left: 1rem;
}
td {
	border: none;
}
.dropdown-item {
	color: #212529 !important;
}
</style>

<template>
	<tr>
		<td>
			<div class="form-check">
				<input
					class="form-check-input"
					type="checkbox"
					id="{{access.id}}"
					v-model="selectionCheckbox"
				/>
				<label class="form-check-label" for="{{access.id}}">
					{{ name }}
				</label>
			</div>
		</td>
		<!-- <td>
    {{ permissions }}
  </td>
  <td>
    {{ duration }}
  </td>
  <td>
    {{ buckets }}
  </td> -->
		<td>
			{{ date }}
		</td>
		<td>
			<div class="dropleft">
				<div
					v-if="accessBeingDeleted"
					class="spinner-border"
					role="status"
				></div>
				<button
					v-else
					class="btn btn-white btn-actions"
					type="button"
					aria-haspopup="true"
					aria-expanded="false"
					v-on:click="toggleDropdown"
				>
					<svg
						width="4"
						height="16"
						viewBox="0 0 4 16"
						fill="none"
						xmlns="http://www.w3.org/2000/svg"
					>
						<path
							d="M3.2 1.6C3.2 2.48366 2.48366 3.2 1.6 3.2C0.716344 3.2 0 2.48366 0 1.6C0 0.716344 0.716344 0 1.6 0C2.48366 0 3.2 0.716344 3.2 1.6Z"
							fill="#7C8794"
						/>
						<path
							d="M3.2 8C3.2 8.88366 2.48366 9.6 1.6 9.6C0.716344 9.6 0 8.88366 0 8C0 7.11634 0.716344 6.4 1.6 6.4C2.48366 6.4 3.2 7.11634 3.2 8Z"
							fill="#7C8794"
						/>
						<path
							d="M1.6 16C2.48366 16 3.2 15.2837 3.2 14.4C3.2 13.5163 2.48366 12.8 1.6 12.8C0.716344 12.8 0 13.5163 0 14.4C0 15.2837 0.716344 16 1.6 16Z"
							fill="#7C8794"
						/>
					</svg>
				</button>
				<div class="dropdown-menu shadow show" v-if="dropdownOpen">
					<button
						v-if="!deleteConfirmation"
						class="dropdown-item action p-3"
						v-on:click="confirmDeletion"
					>
						<svg
							width="1.5em"
							height="1.5em"
							viewBox="0 0 16 16"
							class="bi bi-x mr-1"
							fill="currentColor"
							xmlns="http://www.w3.org/2000/svg"
						>
							<path
								fill-rule="evenodd"
								d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z"
							/>
						</svg>
						Delete
					</button>
					<div v-else>
						<p class="deletion-confirmation mx-5 pt-3">
							Are you sure?
						</p>
						<div class="d-flex">
							<button
								class="dropdown-item trash p-3 action"
								v-on:click="finalDelete"
							>
								<svg
									xmlns="http://www.w3.org/2000/svg"
									width="15"
									height="15"
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
								v-on:click="cancelDeletion"
							>
								<svg
									width="1.5em"
									height="1.5em"
									viewBox="0 0 16 16"
									class="bi bi-x mr-1"
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
		</td>
	</tr>
</template>

<script>
export default {
	name: "AccessTableEntry",
	props: ["access"],
	data: () => ({
		deleteConfirmation: false,
		selectionCheckbox: null
	}),
	computed: {
		name() {
			return this.access.name;
		},
		date() {
			return new Date(this.access.createdAt).toLocaleString();
		},
		accessBeingDeleted() {
			return this.$store.state.access.accessKeysBeingDeleted.find(
				(access) => access === this.access.id
			);
		},
		dropdownOpen() {
			return this.$store.state.access.openedDropdown === this.access.name;
		}
	},
	watch: {
		selectionCheckbox() {
			if (this.selectionCheckbox) {
				this.$store.dispatch(
					"access/addSelectedAccessKey",
					this.access.id
				);
			} else {
				this.$store.dispatch(
					"access/deleteSelectedAccessKey",
					this.access.id
				);
			}
		}
	},
	methods: {
		toggleDropdown(event) {
			event.stopPropagation();

			if (this.dropdownOpen) {
				this.$store.dispatch("access/closeDropdown");
			} else {
				this.$store.dispatch("access/openDropdown", this.access.name);
			}

			this.deleteConfirmation = false;
		},
		confirmDeletion(event) {
			event.stopPropagation();
			this.deleteConfirmation = true;
		},
		finalDelete(event) {
			event.stopPropagation();
			this.$store.dispatch("access/closeDropdown");
			this.$store.dispatch("access/deleteAccessKeys", this.access.id);
			this.deleteConfirmation = false;
		},
		cancelDeletion(event) {
			event.stopPropagation();
			// close dropdown
			this.dropdownOpen = false;
			this.deleteConfirmation = false;
		}
	}
};
</script>
