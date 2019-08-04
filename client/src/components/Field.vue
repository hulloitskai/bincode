<template>
  <div
    class="field flex"
    :class="[{ focused, readonly }, theme]"
    @click="handleClick"
  >
    <h5 class="label">{{ name }}</h5>
    <textarea
      :class="inputClass"
      v-if="multiline"
      v-model="model"
      v-bind="inputProps"
      @focus="handleInputFocus"
      @blur="handleInputBlur"
      ref="input"
    />
    <input
      :class="inputClass"
      type="text"
      v-else
      v-model="model"
      v-bind="inputProps"
      @focus="handleInputFocus"
      @blur="handleInputBlur"
      ref="input"
    />
  </div>
</template>

<script>
import theme, { themeProp } from "@/types/theme";

export default {
  props: {
    value: {
      type: String,
      required: true,
    },
    name: {
      type: String,
      required: true,
    },
    theme: {
      default: theme.LIGHT,
      ...themeProp,
    },
    multiline: {
      type: Boolean,
      default: false,
    },
    readonly: {
      type: Boolean,
      default: false,
    },
    mono: {
      type: Boolean,
      default: false,
    },
  },

  data: () => ({ focused: false }),
  computed: {
    // prettier-ignore
    model: {
      get() { return this.value; },
      set(val) { this.$emit("input", val); },
    },

    // prettier-ignore
    inputClass() { return ["input", { mono: this.mono }]; },
    inputProps() {
      const { $attrs, readonly } = this;
      return { ...$attrs, readonly };
    },
  },

  // prettier-ignore
  methods: {
    handleClick() {
      if (!window.getSelection().isCollapsed) return;
      this.$refs.input.focus();
    },
    handleInputFocus() { this.focused = true; },
    handleInputBlur() { this.focused = false; },
  }
};
</script>

<style lang="scss" scoped>
.field {
  padding: 8px 10px;
  border: 1px solid #aab7cd;
  border-radius: 4px;
  cursor: text;

  // prettier-ignore
  &:not(.readonly) {
    transition: border 200ms ease-in-out;
    &:hover { border-color: #8ba8db; }
    &.focused { border-color: #4579d4; }
  }

  .label {
    margin-bottom: 5px;
    font-weight: 500;
    font-size: 9pt;
    color: #aab7cd;
  }

  .input {
    flex: 1;
    width: 100%;
    box-sizing: border-box;
    outline: none;
    border: none;
    resize: none;

    color: #6691d8;
    background: inherit;
    font-size: 11pt;
    font-weight: 500;

    // prettier-ignore
    &::placeholder { color: #b8cff4; }
  }

  &.dark {
    background: #3b4658;

    // prettier-ignore
    .label { color: #9eb6dd; }

    // prettier-ignore
    .input {
      color: #c9ddff;
      &::placeholder { color: #61789c; }
    }
  }
}
</style>
