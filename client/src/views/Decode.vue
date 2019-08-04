<template>
  <div class="decode">
    <section class="inputs flex">
      <field
        class="input"
        name="Ciphertext"
        placeholder="dEprEssIoNDEpressIONdePReSSIoNdEpREsSIoNdeprEsSi..."
        v-model="text"
        :theme="theme.DARK"
        multiline
        mono
      />
      <custom-button :theme="theme.DARK" @click="handleSubmit">
        decode
      </custom-button>
    </section>
    <section class="results flex" v-if="result">
      <field
        id="decode-result"
        class="result"
        name="Plaintext"
        v-model="result"
        multiline
        readonly
      />
      <copy-button target="#decode-result.input" />
    </section>
  </div>
</template>

<script>
import Field from "@/components/Field";
import Button, { CopyButton } from "@/components/Button";

import theme from "@/types/theme";
import bincode from "@/services/bincode";

export default {
  data: () => ({ text: "", result: "", theme }),
  methods: {
    async handleSubmit() {
      if (!this.text) return;
      this.result = await bincode.decode(this.text);
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
  &.input { height: 391px; }
  &.result {
    height: 200px;
    margin-top: 16px;
  }
}

.button {
  margin: 12px 0 6px 0;
  align-self: flex-end;
}
</style>
