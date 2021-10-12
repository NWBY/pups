<template>
  <div class="home">
    <p @click="getMessage">Click</p>
  </div>
</template>

<script>
import { ref } from "vue";

export default {
  name: "Home",
  components: {
  },
  setup() {
    const getMessage = () => {
      window.backend
        .Beluga
        .GetAllContainers()
        .then(console.log)
        .catch((err) => {
          console.log(err);
        });
    };

    return {
      message: message,
      getMessage: getMessage,
    };
  },
};
</script>
