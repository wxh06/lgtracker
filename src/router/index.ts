import type { RouterOptions } from "vite-ssg";
import ContestsView from "@/views/ContestsView.vue";

const routes: RouterOptions["routes"] = [
  {
    path: "/",
    name: "contests",
    component: ContestsView,
  },
];

export default routes;
