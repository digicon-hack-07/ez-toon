import { createRouter, createWebHistory } from "vue-router";
import MenuPage from "./pages/MenuPage.vue";
import SelectPage from "./pages/SelectPage.vue";
import NotFound from "./pages/NotFound.vue";
import PageViewer from "./pages/PageViewer.vue";

const routes = [
  { path: "/project/:id", name: "ProjectPage", component: SelectPage},
  { path: "/page/:id", name: "PageViewer", component: PageViewer},
  { path: "/", name: "Menu",component: MenuPage },
  { path: "/:path(.*)",name: "not found", component: NotFound },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;