import browserslist from "browserslist"
import { browserslistToTargets } from "lightningcss"
import { resolve } from "path"
import { defineConfig, loadEnv } from "vite"
import lucidePreprocess from "vite-plugin-lucide-preprocess"
import { ViteMinifyPlugin } from "vite-plugin-minify"
import { VitePWA } from "vite-plugin-pwa"
import solid from "vite-plugin-solid"

const root: string = resolve(__dirname, "src")

// @ts-ignore
export default ({ mode }) => {
	Object.assign(process.env, loadEnv(mode, process.cwd()))
	return defineConfig({
		root,
		base: "/pifi",
		css: { transformer: "lightningcss", lightningcss: { targets: browserslistToTargets(browserslist("defaults")) } },
		server: { port: 3000 },
		build: {
			minify: "terser",
			cssMinify: "lightningcss",
			target: "esnext",
			outDir: "../dist",
			emptyOutDir: true,
			sourcemap: false,
			terserOptions: { mangle: { toplevel: true }, compress: { passes: 2 }, format: { comments: false } },
		},
		envDir: "../",
		plugins: [
			lucidePreprocess(),
			solid(),
			ViteMinifyPlugin(),
			VitePWA({
				devOptions: { enabled: true },
				injectRegister: "inline",
				registerType: "autoUpdate",
				workbox: {
					clientsClaim: true,
					skipWaiting: true,
					cleanupOutdatedCaches: true,
					globPatterns: ["**/*.{js,css,html,ico,txt,woff2,webp,png,svg}"],
					runtimeCaching: undefined,
				},
				manifestFilename: "site.webmanifest",
				manifest: {
					name: "PiFi",
					description: "Wardriving toolkit",
					short_name: "PiFi",
					start_url: "/pifi/",
					id: "/",
					lang: "en-US",
					theme_color: "#cc2f5e",
					display: "standalone",
					orientation: "portrait-primary",
					background_color: "#1a1c1d",
					shortcuts: [],
					prefer_related_applications: false,
					related_applications: [],
					categories: ["utilities"],
					icons: [{ src: "/pifi/img/logo.png", sizes: "512x512", type: "image/png", purpose: "any" }, {
						src: "/pifi/img/logo-maskable.png",
						sizes: "512x512",
						type: "image/png",
						purpose: "maskable",
					}],
				},
			}),
		],
	})
}
