import { createRouter, createWebHistory } from "vue-router";
import { useStore } from "../stores/store";

const routes = [
  {
    path: "/",
    name: "login",
    component: () => import("@/views/pages/auth/Login.vue"),
  },
  {
    path: "/auth/callback",
    name: "auth-callback",
    component: () => import("@/views/pages/auth/AuthCallback.vue"),
  },
  {
    path: "/success",
    name: "success",
    component: () => import("@/views/pages/Success.vue"),
    beforeEnter: [guard], 
  },
  
];

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { left: 0, top: 0 };
  },
});

function guard(to, from, next) {
  const store = useStore();
  let TokenStorage = store.token;
  
  if (TokenStorage) {
    next();
  } else {
    next("/");
  }
}

export default router;
