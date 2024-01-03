/** @type {import('tailwindcss').Config} */
module.exports = {
	content: [
		"./src/**/*.{js,jsx,ts,tsx}",
	],
	theme: {
		extend: {
			colors: {
				offwhite: "#FEFFFE",
				indianred: {
					DEFAULT: "#C95D63",
				},
				oxfordblue: {
					DEFAULT: "#0C1B33",
				},
				lightlavender: {
					DEFAULT: "#E9EBF8",
				},
				frenchgrey: {
					DEFAULT: "#B4B8C5",
				},
			},
		},
	},
	plugins: [],
};
