import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

const { BASE_URL } = process.env;
const router = new Router({
  mode: "history",
  base: BASE_URL,
  routes: [
    {
      path: "/",
      redirect: "/encode",
    },
    {
      path: "/encode",
      name: "encode",
      component: () =>
        import(/* webpackChunkName: "encode" */ "@/views/Encode.vue"),
    },
    {
      path: "/decode",
      name: "decode",
      component: () =>
        import(/* webpackChunkName: "encode" */ "@/views/Decode.vue"),
    },
  ],
});

export default router;
