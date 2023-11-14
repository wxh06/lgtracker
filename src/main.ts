import "primevue/resources/themes/lara-light-teal/theme.css";
import "primeicons/primeicons.css";

import { ViteSSG } from "vite-ssg";
import { createPinia } from "pinia";
import PrimeVue from "primevue/config";
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
    app.use(PrimeVue);
    app.directive("tooltip", Tooltip);
  },
);
