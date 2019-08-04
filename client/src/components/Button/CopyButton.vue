<template>
  <custom-button
    class="copy-button tooltipped-n tooltipped-no-delay"
    :class="{ tooltipped: clicked }"
    :data-clipboard-target="target"
    aria-label="Copied!"
    v-bind="$attrs"
    v-on="$listeners"
    @click="handleClick"
  >
    copy
  </custom-button>
</template>

<script>
import Button from "./Button";

export default {
  props: {
    target: {
      type: String,
      required: true,
    },
  },
  data: () => ({ clicked: false }),
  methods: {
    // prettier-ignore
    handleClick() { this.clicked = true; }
  },
  components: {
    "custom-button": Button,
  },
};
</script>
