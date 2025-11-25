import "./Scan.css"
import { type Component, createSignal, For, onMount } from "solid-js"
import AP from "../components/AP.tsx"
import { type bettercapWifiData, getBettercapWifi } from "../lib/bettercap.ts"
export const [bettercapWifi, setBettercapWifi] = createSignal<bettercapWifiData[]>([])

const Scan: Component = () => {
	onMount(async () => {
		setBettercapWifi(await getBettercapWifi())
	})

	return (
		<div id="scan">
			<For each={bettercapWifi() as bettercapWifiData[]}>{(ap, _) => <AP ap={ap} />}</For>
		</div>
	)
}

export default Scan
