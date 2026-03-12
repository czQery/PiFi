import "./Scan.css"
import { createPolled } from "@solid-primitives/timer"
import { type Component, createSignal, For, Show } from "solid-js"
import AP from "../components/AP.tsx"
import Error from "../components/Error.tsx"
import { type bettercapWifiAPData, getBettercapWifi } from "../lib/api/bettercap.ts"

export const [bettercapWifi, setBettercapWifi] = createSignal<bettercapWifiAPData[]>([])

const Scan: Component = () => {
	createPolled(async () => setBettercapWifi(await getBettercapWifi()), 5000)

	return (
		<div id="scan">
			<Show when={bettercapWifi().length !== 0} fallback={<Error title="No data!" msg="You probably don't have any devices in monitoring mode." />}>
				<For each={bettercapWifi() as bettercapWifiAPData[]}>{(ap, _) => <AP ap={ap} />}</For>
			</Show>
		</div>
	)
}

export default Scan
