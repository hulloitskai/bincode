<template>
  <div id="app">
    <home />
  </div>
</template>

<script>
import Home from "./views/Home";

export default {
  metaInfo: {
    titleTemplate: "Bincode – %s",
  },
  components: {
    home: Home,
  },
};
</script>

<style lang="scss">
@import "@/styles/fonts.scss";

// Global style defaults:
body {
  font-family: $sans;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

// prettier-ignore
h1, h2, h3, h4, h5, h6, p { margin: 0; }

// Flexbox shortcut classes:
.flex {
  display: flex;
  flex-direction: column;
}

.center-children {
  align-items: center;
  justify-content: center;
}

// Font overrides:
.sans {
  font-family: $sans;
}
.mono {
  font-family: $mono;
}
</style>
