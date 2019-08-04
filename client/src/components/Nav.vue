<template>
  <section class="nav">
    <router-link
      class="link"
      v-for="section in sections"
      :to="section"
      :key="section"
    >
      <div class="section" :class="{ active: isSectionActive(section) }">
        <h2 class="name">
          {{ section }}
        </h2>
        <div class="underline" v-if="isSectionActive(section)" />
      </div>
    </router-link>
    <div class="logo-container flex center-children">
      <logo />
    </div>
  </section>
</template>

<script>
import Logo from "@/components/Logo";

export default {
  data: () => ({ sections: ["encode", "decode"] }),
  methods: {
    /**
     * @param {string} section
     * @returns {boolean}
     */
    isSectionActive(section) {
      const active = this.$route.path.slice(1);
      return active === section;
    },
  },
  components: {
    logo: Logo,
  },
};
</script>

<style lang="scss" scoped>
.nav {
  display: flex;
  margin-top: 4px;
  padding: 4px 130px 0 35px;
  border-bottom: solid 3px #dde5f2;
  position: relative;
}

// prettier-ignore
.link { text-decoration: none; }

.section {
  position: relative;
  padding: 12px;
  cursor: pointer;
  transition: color 200ms ease-in-out;

  .name {
    color: #a2b1ca;
    font-size: 14pt;
    font-weight: 600;
  }

  // prettier-ignore
  .underline {
    position: absolute;
    bottom: -3px; left: 0; right: 0;
    height: 3px;
    background: #4d7ecd;
  }

  // prettier-ignore
  &:hover .name { color: lighten(#4d7ecd, 20%); }

  // prettier-ignore
  &.active .name { color: #4d7ecd; }
}

// prettier-ignore
.logo-container {
  position: absolute;
  top: 0; right: 0; bottom: 0;
  padding-right: 10px;
}
</style>
