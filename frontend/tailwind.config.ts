import { join } from "path";
import type { Config } from "tailwindcss";
import type { CustomThemeConfig } from "@skeletonlabs/tw-plugin";
import { skeleton } from "@skeletonlabs/tw-plugin";

export const customtheme: CustomThemeConfig = {
	name: "custom-theme",
	properties: {
		// =~= Theme Properties =~=
		"--theme-font-family-base": "'Inter', sans-serif",
		"--theme-font-family-heading": "'Abril Fatface', sans-serif",
		"--theme-font-color-base": "32 32 32",
		"--theme-font-color-dark": "212 212 212",
		"--theme-rounded-base": "4px",
		"--theme-rounded-container": "8px",
		"--theme-border-base": "1px",
		// =~= Theme On-X Colors =~=
		"--on-primary": "255 255 255",
		"--on-secondary": "0 0 0",
		"--on-tertiary": "0 0 0",
		"--on-success": "0 0 0",
		"--on-warning": "0 0 0",
		"--on-error": "0 0 0",
		"--on-surface": "255 255 255",
		// =~= Theme Colors  =~=
		// primary | #4d53fe
		"--color-primary-50": "228 229 255", // #e4e5ff
		"--color-primary-100": "219 221 255", // #dbddff
		"--color-primary-200": "211 212 255", // #d3d4ff
		"--color-primary-300": "184 186 255", // #b8baff
		"--color-primary-400": "130 135 254", // #8287fe
		"--color-primary-500": "77 83 254", // #4d53fe
		"--color-primary-600": "69 75 229", // #454be5
		"--color-primary-700": "58 62 191", // #3a3ebf
		"--color-primary-800": "46 50 152", // #2e3298
		"--color-primary-900": "38 41 124", // #26297c
		// secondary | #5c92ff
		"--color-secondary-50": "231 239 255", // #e7efff
		"--color-secondary-100": "222 233 255", // #dee9ff
		"--color-secondary-200": "214 228 255", // #d6e4ff
		"--color-secondary-300": "190 211 255", // #bed3ff
		"--color-secondary-400": "141 179 255", // #8db3ff
		"--color-secondary-500": "92 146 255", // #5c92ff
		"--color-secondary-600": "83 131 230", // #5383e6
		"--color-secondary-700": "69 110 191", // #456ebf
		"--color-secondary-800": "55 88 153", // #375899
		"--color-secondary-900": "45 72 125", // #2d487d
		// tertiary | #8af7ff
		"--color-tertiary-50": "237 254 255", // #edfeff
		"--color-tertiary-100": "232 253 255", // #e8fdff
		"--color-tertiary-200": "226 253 255", // #e2fdff
		"--color-tertiary-300": "208 252 255", // #d0fcff
		"--color-tertiary-400": "173 249 255", // #adf9ff
		"--color-tertiary-500": "138 247 255", // #8af7ff
		"--color-tertiary-600": "124 222 230", // #7cdee6
		"--color-tertiary-700": "104 185 191", // #68b9bf
		"--color-tertiary-800": "83 148 153", // #539499
		"--color-tertiary-900": "68 121 125", // #44797d
		// success | #9bff99
		"--color-success-50": "240 255 240", // #f0fff0
		"--color-success-100": "235 255 235", // #ebffeb
		"--color-success-200": "230 255 230", // #e6ffe6
		"--color-success-300": "215 255 214", // #d7ffd6
		"--color-success-400": "185 255 184", // #b9ffb8
		"--color-success-500": "155 255 153", // #9bff99
		"--color-success-600": "140 230 138", // #8ce68a
		"--color-success-700": "116 191 115", // #74bf73
		"--color-success-800": "93 153 92", // #5d995c
		"--color-success-900": "76 125 75", // #4c7d4b
		// warning | #ffcf99
		"--color-warning-50": "255 248 240", // #fff8f0
		"--color-warning-100": "255 245 235", // #fff5eb
		"--color-warning-200": "255 243 230", // #fff3e6
		"--color-warning-300": "255 236 214", // #ffecd6
		"--color-warning-400": "255 221 184", // #ffddb8
		"--color-warning-500": "255 207 153", // #ffcf99
		"--color-warning-600": "230 186 138", // #e6ba8a
		"--color-warning-700": "191 155 115", // #bf9b73
		"--color-warning-800": "153 124 92", // #997c5c
		"--color-warning-900": "125 101 75", // #7d654b
		// error | #ff9999
		"--color-error-50": "255 240 240", // #fff0f0
		"--color-error-100": "255 235 235", // #ffebeb
		"--color-error-200": "255 230 230", // #ffe6e6
		"--color-error-300": "255 214 214", // #ffd6d6
		"--color-error-400": "255 184 184", // #ffb8b8
		"--color-error-500": "255 153 153", // #ff9999
		"--color-error-600": "230 138 138", // #e68a8a
		"--color-error-700": "191 115 115", // #bf7373
		"--color-error-800": "153 92 92", // #995c5c
		"--color-error-900": "125 75 75", // #7d4b4b
		// surface | #1e1e2e
		"--color-surface-50": "221 221 224", // #dddde0
		"--color-surface-100": "210 210 213", // #d2d2d5
		"--color-surface-200": "199 199 203", // #c7c7cb
		"--color-surface-300": "165 165 171", // #a5a5ab
		"--color-surface-400": "98 98 109", // #62626d
		"--color-surface-500": "30 30 46", // #1e1e2e
		"--color-surface-600": "27 27 41", // #1b1b29
		"--color-surface-700": "23 23 35", // #171723
		"--color-surface-800": "18 18 28", // #12121c
		"--color-surface-900": "15 15 23", // #0f0f17
	},
};

const config = {
	darkMode: "class",
	content: [
		"./src/**/*.{html,js,svelte,ts}",
		join(
			require.resolve("@skeletonlabs/skeleton"),
			"../**/*.{html,js,svelte,ts}",
		),
	],
	theme: {
		extend: {
			boxShadow: {
				"3xl": "0 35px 60px -15px rgba(0, 0, 0, 0.3)",
			},
		},
	},
	plugins: [
		skeleton({
			themes: {
				custom: [customtheme],
			},
		}),
	],
} satisfies Config;

export default config;
