const eslint = "eslint --fix";
const prettier = "prettier --write";
const gofmt = "gofmt -w";

module.exports = {
  "*.{js,mjs,cjs,ts,vue}": [eslint, prettier],
  "*.{md,html,css,scss,json,yml,yaml}": prettier,
  "*.go": gofmt,
};
