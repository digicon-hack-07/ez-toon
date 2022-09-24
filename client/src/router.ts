import { createRouter, createWebHistory } from "vue-router";
import MenuPage from "./pages/MenuPage.vue";

const routes = [
  { path: "/", name:"menu",component: MenuPage },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;