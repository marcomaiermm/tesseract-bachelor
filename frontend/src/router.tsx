import { createReactRouter, createRouteConfig } from "@tanstack/react-router";
import Layout from "@pages/layout";
import Index from "@pages/index";
import Complaints from "@pages/complaints";

const rootRoute = createRouteConfig();

const indexRoute = rootRoute.createRoute({ path: "/", component: Index });
const complaintsRoute = rootRoute.createRoute({
  path: "/complaints",
  component: Complaints,
});

const layoutRoute = rootRoute.createRoute({
  id: "layout",
  component: Layout,
});

const routeConfig = rootRoute.addChildren([
  layoutRoute.addChildren([indexRoute, complaintsRoute]),
]);

export const router = createReactRouter({ routeConfig });
console.log(router);
