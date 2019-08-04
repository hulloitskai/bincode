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
        ref="copy"
      />
      <copy-button target="#decode-result.input" ref="copy" />
    </section>
  </div>
</template>

<script>
import Field from "@/components/Field";
import Button, { CopyButton } from "@/components/Button";

import theme from "@/types/theme";
import bincode from "@/services/bincode";
import { scrollIntoView } from "@/utils/dom";
import { setTimeoutFrom } from "@/utils/time";

export default {
  data: () => ({ text: "", result: "", theme }),
  methods: {
    async handleSubmit() {
      const start = new Date();
      if (!this.text) return;
      this.result = await bincode.decode(this.text);
      setTimeoutFrom(start, () => scrollIntoView(this.$refs.copy.$el), 200);
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
