import "./Scan.css"
import { createPolled } from "@solid-primitives/timer"
import { type Component, createSignal, For, Show } from "solid-js"
import AP from "../components/AP.tsx"
import DialogLoading, { type dialogLoadingData } from "../components/dialog/DialogLoading.tsx"
import DialogWrapper, { type dialogWrapperData } from "../components/dialog/DialogWrapper.tsx"
import Error from "../components/Error.tsx"
import { type bettercapWifiAPData, getBettercapWifi } from "../lib/api/bettercap.ts"

export const [bettercapWifi, setBettercapWifi] = createSignal<bettercapWifiAPData[]>([])

const Scan: Component = () => {
	createPolled(async () => setBettercapWifi(await getBettercapWifi()), 5000)

	const [loading, setLoading] = createSignal<dialogLoadingData>({ title: "Loading", pending: false, msg: "" })
	const [option, setOption] = createSignal<dialogWrapperData>({
		open: false,
		title: "Attack",
		message: "",
		helper: null,
		buttonFirst: null,
		buttonSecond: null,
	})

	return (
		<div id="scan">
			<DialogWrapper data={option()} />
			<Show when={bettercapWifi().length !== 0} fallback={<Error title="No data!" msg="You probably don't have any devices in monitoring mode." />}>
				<For each={bettercapWifi() as bettercapWifiAPData[]}>
					{(ap, _) => <AP ap={ap} option={option} setOption={setOption} setLoading={setLoading} />}
				</For>
			</Show>
			<DialogLoading data={loading} setData={setLoading} />
		</div>
	)
}

export default Scan
