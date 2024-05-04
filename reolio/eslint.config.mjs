import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";
import pluginReactConfig from "eslint-plugin-react/configs/recommended.js";


export default [
  {
    languageOptions: { globals: globals.browser },
  },
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
  pluginReactConfig,
  {
    "settings": {
      "react": {
        "version": "detect"
      }
    },
    "fix": true
  }
];

// import globals from "globals";
// import pluginJs from "@eslint/js";
// import tseslint from "typescript-eslint";
// import pluginReactConfig from "eslint-plugin-react/configs/recommended.js";

// export default {
//   root: true,
//   parserOptions: {
//     ecmaVersion: 2021,
//     sourceType: "module",
//   },
//   plugins: [
//     "react", // 通常、プラグイン名だけで十分です
//     "@typescript-eslint"
//   ],
//   extends: [
//     "eslint:recommended",
//     "plugin:@typescript-eslint/recommended",
//     "plugin:react/recommended",
//     // 他の推奨設定を含める
//   ],
//   rules: {
//     // 特定のルールを上書き
//   },
//   globals: globals.browser
// };
