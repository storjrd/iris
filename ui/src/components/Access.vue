<style scoped>
tbody {
  border-top: 1px solid #dee2e6;
}
</style>

<template>
  <h5>Access Keys</h5>

  <div class="d-flex justify-content-end mb-4">
    <button type="button" class="btn btn-primary btn-sm">Create Access +</button>
  </div>

  <div>
		<table class="table table-hover">

      <access-table-header></access-table-header>

      <tbody>
        <access-table-entry v-for="access in accessKeys" v-bind:access="access"
							v-bind:key="access.name"></access-table-entry>
      </tbody>
    </table>
  </div>
</template>

<script>
import AccessTableEntry from './AccessTableEntry.vue';
import AccessTableHeader from "./AccessTableHeader.vue";

export default {
  name: "Access",
  data: () => ({
    loadingAccessSpinner: true
  }),
  computed: {
    accessKeys() {
      return this.$store.getters["access/sortedAccessKeys"];
    }
  },
  async created() {
    // fetch access keys
		await this.$store.dispatch("access/list");
		this.loadingAccessSpinner = false;
	},
  components: {
    AccessTableHeader,
    AccessTableEntry,
    AccessTableEntry
  }
}
</script>
