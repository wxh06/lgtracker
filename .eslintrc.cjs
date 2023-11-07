require("@rushstack/eslint-patch/modern-module-resolution");

module.exports = {
  root: true,
  extends: [
    "plugin:vue/vue3-essential",
    "eslint:recommended",
    "@vue/eslint-config-typescript/recommended",
    "@vue/eslint-config-prettier/skip-formatting",
  ],
  parserOptions: {
    ecmaVersion: "latest",
  },
  overrides: [
    {
      files: [".*rc.*", "*.config.*"],
      env: { node: true },
    },
  ],
};
