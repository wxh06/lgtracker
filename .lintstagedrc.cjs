const eslint = "eslint --fix";
const prettier = "prettier --write";

module.exports = {
  "*.{js,mjs,cjs,ts,vue}": [eslint, prettier],
  "*.{md,html,css,scss,json,yml,yaml}": prettier,
};
