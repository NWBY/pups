<template>
  <div class="home">
    <containers-list :containers="containers"></containers-list>
  </div>
</template>

<script>
import { ref } from "vue";

import ContainersList from '../components/containers/List.vue'

export default {
  name: "Home",
  components: {
    ContainersList
  },
  mounted() {
    this.getContainers()
  },
  setup() {
    let containers = ref([])

    const getContainers = () => {
      window.backend
        .Beluga
        .GetAllContainers()
        .then(res => {
          containers.value = res
        })
        .catch((err) => {
          console.log(err);
        });
    };

    return {
      containers,
      getContainers
    };
  },
};
</script>

<style scoped>
.home {
  width: 75vw;
  margin: 0 auto;
}
</style>
