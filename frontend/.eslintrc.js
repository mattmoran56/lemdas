module.exports = {
	extends: [
		"react-app",
		"react-app/jest",
		"airbnb",
	],
	rules: {
		"react/jsx-filename-extension": [
			"error",
			{
				extensions: [".js", ".jsx"],
			},
		],
		quotes: [
			"error",
			"double",
		],
		indent: [
			"error",
			"tab",
			{
				SwitchCase: 1,
			},
		],
		"react/jsx-indent": [
			"error",
			"tab",
		],
		"react/jsx-indent-props": [
			"error",
			"tab",
		],
		"react/function-component-definition": [
			"error",
			{
				namedComponents: "arrow-function",
				unnamedComponents: "arrow-function",
			},
		],
		"react/no-unknown-property": [
			"error", { ignore: ["test-id"] },
		],
		"import/extensions": "off",
		"import/no-unresolved": "off",
		"react/require-default-props": "off",
		"react/jsx-one-expression-per-line": "off",
		"arrow-body-style": ["error", "always"],
		"max-len": [
			"error",
			120,
		],
		"no-tabs": "off",
		"react/prop-types": "off",
	},
};
