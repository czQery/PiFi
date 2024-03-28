import {resolve} from "path"
import {defineConfig} from "vite"
import solidPlugin from "vite-plugin-solid"

const root: string = resolve(__dirname, "src")

export default defineConfig({
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
