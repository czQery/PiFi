import browserslist from "browserslist"
import { browserslistToTargets } from "lightningcss"
import { resolve } from "path"
import { defineConfig, loadEnv } from "vite"
import lucidePreprocess from "vite-plugin-lucide-preprocess"
import { ViteMinifyPlugin } from "vite-plugin-minify"
import solid from "vite-plugin-solid"

const root: string = resolve(__dirname, "src")

// @ts-ignore
export default ({ mode }) => {
	Object.assign(process.env, loadEnv(mode, process.cwd()))
	return defineConfig({
		root,
		base: "/pifi",
		plugins: [lucidePreprocess(), solid(), ViteMinifyPlugin()],
		css: { transformer: "lightningcss", lightningcss: { targets: browserslistToTargets(browserslist(">= 0.25%")) } },
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
	})
}
