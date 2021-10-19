<template>
  <div class="home">
    <containers-list :containers="containers"></containers-list>
    <volumes-list :volumes="volumes"></volumes-list>
  </div>
</template>

<script>
import { ref } from "vue";

import ContainersList from '../components/containers/List.vue'
import VolumesList from '../components/volumes/List.vue'

export default {
  name: "Home",
  components: {
    ContainersList,
    VolumesList
  },
  mounted() {
    this.getContainers()
    this.getVolumes()
  },
  setup() {
    let containers = ref([])
    let volumes = ref([])

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

    const getVolumes = () => {
      window.backend
        .Beluga
        .GetAllVolumes()
        .then(res => {
          volumes.value = res
        })
        .catch((err) => {
          console.log(err);
        });
    };

    return {
      containers,
      getContainers,
      volumes,
      getVolumes
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
