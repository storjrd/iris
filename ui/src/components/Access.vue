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
      // retrieve access keys

      const demoAccessObjects = [
        {
          name: "FoodApp",
          permissions: "Read, Write",
          duration: "Forever",
          buckets: "All",
          dateCreated: "07/06/2021 12:27PM EST"
        },
        {
          name: "Iris",
          permissions: "Read, Write, List, Delete",
          duration: "06/21/2021 - 08/01/2021",
          buckets: "5",
          dateCreated: "06/21/2021 10:04AM EST"
        },
        {
          name: "NFT Kitties",
          permissions: "Read, Write, List",
          duration: "05/14/2021 - 09/22/2021",
          buckets: "2",
          dateCreated: "05/14/2021 3:11PM EST"
        },
        {
          name: "Todo iOS App",
          permissions: "List, Delete",
          duration: "04/29/2021 - 08/05/2021",
          buckets: "8",
          dateCreated: "04/29/2021 8:02AM EST"
        },
        {
          name: "Watermelon Picker",
          permissions: "Write",
          duration: "Forever",
          buckets: "All",
          dateCreated: "03/19/2021 7:12PM EST"
        },
        {
          name: "Accounting",
          permissions: "Read, Write, Delete",
          duration: "Forever",
          buckets: "13",
          dateCreated: "01/30/2021 11:20PM EST"
        }
      ];

      return demoAccessObjects;
    },
    sortAccessKeys() {

    }
  },
  async created() {
    // fetch access keys
		// await this.$store.dispatch("");
		this.loadingAccessSpinner = false;
	},
  components: {
    AccessTableHeader,
    AccessTableEntry,
    AccessTableEntry
  }
}
</script>
