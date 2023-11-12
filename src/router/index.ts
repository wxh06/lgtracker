import type { RouterOptions } from "vite-ssg";
import ContestView from "@/views/ContestView.vue";

const routes: RouterOptions["routes"] = [
  {
    path: "/",
    name: "contests",
    component: ContestView,
  },
];

export default routes;
