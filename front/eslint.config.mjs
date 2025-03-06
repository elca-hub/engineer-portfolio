import { FlatCompat } from "@eslint/eslintrc";
import typescriptEslintParser from "@typescript-eslint/parser";
import importAccess from "eslint-plugin-import-access/flat-config";
import { dirname } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const compat = new FlatCompat({
  baseDirectory: __dirname,
});

const eslintConfig = [
  ...compat.extends("next/core-web-vitals", "next/typescript", "prettier"),
  {
    // set up typescript-eslint
    languageOptions: {
      parser: typescriptEslintParser,
      parserOptions: {
        project: true,
        sourceType: "module",
      },
    },
  },
  {
    plugins: {
      "import-access": importAccess,
      "prettier": {
        rules: {
          "prettier/prettier": "error",
        },
      },
    },
  },
  {
    rules: {
      "import-access/jsdoc": ["error"],
      "eqeqeq": ["error", "always"],
      "react-hooks/exhaustive-deps": "off",
    },
  },
  {
    ignores: [
      ".next/*",
      "node_modules/*",
      "out/*",
      "public/*",
      "build/*"
    ]
  }
];

export default eslintConfig;
