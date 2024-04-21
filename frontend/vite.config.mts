import {resolve} from "path"
import {defineConfig, loadEnv} from "vite"
import solidPlugin from "vite-plugin-solid"

const root: string = resolve(__dirname, "src")

export default ({mode}) => {
    Object.assign(process.env, loadEnv(mode, process.cwd()))
    return defineConfig({
        root,
        plugins: [
            solidPlugin(),
        ],
        server: {
            port: 3000,
        },
        build: {
            target: "esnext",
            outDir: "../dist",
            emptyOutDir: true,
        },
        envDir: "../"
    })
}
