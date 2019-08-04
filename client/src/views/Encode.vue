<template>
  <div class="encode">
    <section class="inputs flex">
      <field class="key" name="Key" v-model="key" mono />
      <field
        class="text"
        name="Plaintext"
        placeholder="Lorem ipsum dolor sit amet..."
        v-model="text"
        multiline
      />
      <custom-button class="encode" @click="handleSubmit">encode</custom-button>
    </section>
    <section class="results flex" v-if="result">
      <field
        id="encode-result"
        class="result"
        name="Ciphertext"
        v-model="result"
        :theme="theme.DARK"
        multiline
        readonly
        mono
      />
      <copy-button target="#encode-result.input" :theme="theme.DARK" />
    </section>
  </div>
</template>

<script>
import Field from "@/components/Field";
import Button, { CopyButton } from "@/components/Button";

import theme from "@/types/theme";
import bincode from "@/services/bincode";

export default {
  data: () => ({
    key: "depression",
    text: "",
    result: "",
    theme,
  }),
  methods: {
    async handleSubmit() {
      if (!this.key || !this.text) return;
      this.result = await bincode.encode(this.key, { plaintext: this.text });
    },
  },
  components: {
    field: Field,
    "custom-button": Button,
    "copy-button": CopyButton,
  },
};
</script>

<style lang="scss" scoped>
.field {
  // prettier-ignore
  &.key { width: 250px; }

  &.text {
    height: 325px;
    margin-top: 12px;
  }

  &.result {
    margin-top: 16px;
    height: 200px;
  }
}

.button {
  margin: 12px 0 6px 0;
  align-self: flex-end;
}
</style>
