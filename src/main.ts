import "./assets/main.css";

import { ViteSSG } from "vite-ssg";
import { createPinia } from "pinia";

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
  },
);
