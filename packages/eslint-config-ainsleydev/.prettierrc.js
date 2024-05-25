/** @type {import("prettier").Config} */
module.exports = {
	useTabs: true,
	tabWidth: 4,
	singleQuote: true,
	trailingComma: 'all',
	printWidth: 100,
	editorconfig: true,
	overrides: [
		{
			files: '*.yaml',
			options: {
				singleQuote: false,
			},
		},
		{
			files: '*.yml',
			options: {
				singleQuote: false,
			},
		},
	],
};