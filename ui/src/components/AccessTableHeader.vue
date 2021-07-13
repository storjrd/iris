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
</style>

<template>
  <thead>
        <tr>
          <th class="name d-flex align-items-center">
            NAME
            <div class="container">
              <svg width="9" height="6" fill="none" xmlns="http://www.w3.org/2000/svg" class="">
                <path d="M4.737 0L9 6H0l4.737-6z" fill="#354049" class="arrow-svg-path"></path>
              </svg>
              <svg width="9" height="6" fill="none" xmlns="http://www.w3.org/2000/svg" class="active">
                <path d="M4.263 6L0 0h9L4.263 6z" fill="#354049" class="arrow-svg-path"></path>
              </svg>
            </div>
          </th>
          <th>
            PERMISSIONS
          </th>
          <th class="d-flex align-items-center">
            DURATION
            <div class="container">
              <svg width="9" height="6" fill="none" xmlns="http://www.w3.org/2000/svg" class="">
                <path d="M4.737 0L9 6H0l4.737-6z" fill="#354049" class="arrow-svg-path"></path>
              </svg>
              <svg width="9" height="6" fill="none" xmlns="http://www.w3.org/2000/svg" class="active">
                <path d="M4.263 6L0 0h9L4.263 6z" fill="#354049" class="arrow-svg-path"></path>
              </svg>
            </div>
          </th>
          <th>
            BUCKETS
          </th>
          <th class="d-flex align-items-center">
            DATE CREATED
            <div class="container">
              <svg width="9" height="6" fill="none" xmlns="http://www.w3.org/2000/svg" class="">
                <path d="M4.737 0L9 6H0l4.737-6z" fill="#354049" class="arrow-svg-path"></path>
              </svg>
              <svg width="9" height="6" fill="none" xmlns="http://www.w3.org/2000/svg" class="active">
                <path d="M4.263 6L0 0h9L4.263 6z" fill="#354049" class="arrow-svg-path"></path>
              </svg>
            </div>
          </th>
          <th></th>
        </tr>
      </thead>
</template>

<script>
export default {
  name: "AccessTableHeader"
}
</script>
