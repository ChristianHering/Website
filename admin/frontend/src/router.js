import { createRouter, createWebHashHistory } from "vue-router";
import Home from "./views/Home.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/Errors",
    name: "Errors",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ "./views/Errors.vue")
  },
  {
    path: "/Posts",
    name: "Posts",
    component: () => import(/* webpackChunkName: "posts" */ "./views/Posts.vue")
  },
  {
    path: "/PostCreate",
    name: "PostCreate",
    component: () => import(/* webpackChunkName: "postcreate" */ "./views/PostCreate.vue")
  }
];

const router = createRouter({
  history: createWebHashHistory(),
  routes
});

export default router;
