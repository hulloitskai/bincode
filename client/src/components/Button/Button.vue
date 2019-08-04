<template>
  <button class="button" :class="theme" v-bind="$attrs" v-on="$listeners">
    <slot />
  </button>
</template>

<script>
import theme, { themeProp } from "@/types/theme";

export default {
  props: {
    theme: {
      default: theme.LIGHT,
      ...themeProp,
    },
  },
};
</script>

<style lang="scss" scoped>
.button {
  padding: 8px 12px;
  border: none;
  border-radius: 4px;
  outline: none;

  background: #dde5f2;
  color: #4d7ecd;
  cursor: pointer;
  font-weight: 500;

  transition: all 200ms ease-in-out;
  &:hover {
    background: #d0dff5;
    color: #346ecc;
  }

  &.dark {
    background: #456497;
    color: #bbd6ff;

    &:hover {
      background: #5375ad;
      color: #d5e5ff;
    }
  }
}
</style>
