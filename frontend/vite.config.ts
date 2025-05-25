import {resolve} from "path"
import {defineConfig, loadEnv} from "vite"
import {browserslistToTargets} from "lightningcss"
import solid from "vite-plugin-solid"
import browserslist from "browserslist"

const root: string = resolve(__dirname, "src")

export default ({mode}) => {
	Object.assign(process.env, loadEnv(mode, process.cwd()))
	return defineConfig({
		root,
		base: "/pifi",
		plugins: [
			solid()
		],
		css: {
			transformer: "lightningcss",
			lightningcss: {
				targets: browserslistToTargets(browserslist(">= 0.25%"))
			}
		},
		server: {
			port: 3000
		},
		build: {
			cssMinify: "lightningcss",
			target: "esnext",
			outDir: "../dist",
			emptyOutDir: true
		},
		envDir: "../"
	})
}
