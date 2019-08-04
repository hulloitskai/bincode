<template>
  <div class="home center-children" :class="{ locked }">
    <main class="card" ref="card" :style="cardStyle">
      <custom-nav />
      <section class="content flex">
        <keep-alive>
          <router-view />
        </keep-alive>
      </section>
    </main>
  </div>
</template>

<script>
import Clipboard from "clipboard";
import Nav from "@/components/Nav";

import { screenSizes, getScreenSize } from "@/styles/breakpoints";

export default {
  data: () => ({ offset: 0, locked: false, cardStyle: {} }),

  mounted() {
    // FIXME: Top ten anime jankiest layout tricks.
    // DO NOT EVER DO ANYTHING LIKE THIS EVER AGAIN.
    setTimeout(this.lockCardOffset, 1000);

    this.clipboard = new Clipboard(".copy-button");
  },
  // prettier-ignore
  beforeDestroy() { this.clipboard.destroy(); },

  methods: {
    lockCardOffset() {
      if (getScreenSize() <= screenSizes.PHABLET) return;
      const { top } = this.$refs.card.getBoundingClientRect();
      const margin = `${Math.max(top, 40)}px`;
      this.cardStyle = { marginTop: margin, marginBottom: margin };
      this.locked = true;
    },
  },
  components: {
    "custom-nav": Nav,
  },
};
</script>

<style lang="scss" scoped>
@import "@/styles/breakpoints.scss";

.home {
  width: 100%;
  min-height: 100vh;
  display: flex;
  background: linear-gradient(134deg, #b792ff 0%, #7dbff3 100%);

  // prettier-ignore
  &.locked { align-items: flex-start !important; }
}

.card {
  flex: 1;
  align-self: stretch;
  background: white;
  box-shadow: 0 4px 30px 0 rgba(0, 0, 0, 0.3);

  @include breakpoint(phablet) {
    align-self: auto;
    margin: 40px;
    max-width: 700px;
    border-radius: 8px;
  }
}

.content {
  padding: 22px 30px;
}
</style>
