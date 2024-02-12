/** @type {import('tailwindcss').Config} */
module.exports = {
	content: [
		"./src/**/*.{js,jsx,ts,tsx}",
	],
	theme: {
		extend: {
			animation: {
				'spin-slow': 'spin 0.75s linear infinite',
			},
			colors: {
				offwhite: "#FEFFFE",
				indianred: {
					DEFAULT: "#C95D63",
				},
				oxfordblue: {
					DEFAULT: "#0C1B33",
					200: "#acc5ec",
					extralight: "#eaf0fa",
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
