import { RouteRecordRaw } from "vue-router";
const routes: RouteRecordRaw[] = [
  { path: "/", redirect: "/home" },
  {
    path: "/",
    component: () => import("layouts/MainLayout.vue"),
    children: [
      {
        path: "/home",
        name: "home",
        component: () => import("pages/Index.vue"),
      },
      {
        path: "/list",
        component: () => {
          return import("pages/FileList.vue");
        },
      },
      {
        path: "/transfer",
        name: "transfer",
        component: () => import("pages/Index.vue"),
      },
      {
        path: "/chat",
        name: "chat",
        component: () => import("pages/Index.vue"),
      },
      {
        path: "/text",
        name: "text",
        component: () => import("pages/Index.vue"),
      },   {
        path: "/ws-page",
        name: "ws",
        component: () => import("pages/WsPage.vue"),
      },
      // Always leave this as last one,
      // but you can also remove it
      {
        path: "/:catchAll(.*)*",
        component: () => import("pages/Error404.vue"),
      },
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: "/:catchAll(.*)*",
    component: () => import("pages/Error404.vue"),
  },
];

export default routes;
