import "./assets/main.css";

import { ViteSSG } from "vite-ssg";
import { createPinia } from "pinia";
import PrimeVue from "primevue/config";
import Ripple from "primevue/ripple";
import Tooltip from "primevue/tooltip";

import App from "./App.vue";
import routes from "./router";

export const createApp = ViteSSG(
  // the root component
  App,
  // vue-router options
  { routes },
  // function to have custom setups
  ({ app }) => {
    // install plugins etc.
    app.use(createPinia());
    app.use(PrimeVue, { ripple: true });
    app.directive("tooltip", Tooltip);
    app.directive("ripple", Ripple);
  },
);
