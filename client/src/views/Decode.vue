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
      <custom-button
        class="copy"
        ref="copy"
        :theme="theme.DARK"
        data-clipboard-target="#decode-result.input"
      >
        copy
      </custom-button>
    </section>
  </div>
</template>

<script>
import Field from "@/components/Field";
import Button from "@/components/Button";

import theme from "@/types/theme";
import bincode from "@/services/bincode";

export default {
  data: () => ({ text: "", result: "", theme }),
  components: {
    field: Field,
    "custom-button": Button,
  },
  methods: {
    async handleSubmit() {
      if (!this.text) return;
      this.result = await bincode.decode(this.text);
    },
  },
};
</script>

<style lang="scss" scoped>
// prettier-ignore
.field {
  &.input { height: 391px; }
  &.result { margin-top: 16px; }
}

.button {
  margin: 12px 0 6px 0;
  align-self: flex-end;
}
</style>
