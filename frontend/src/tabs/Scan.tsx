import "./Scan.css"
import { createPolled } from "@solid-primitives/timer"
import { type Component, createSignal, For, Show } from "solid-js"
import AP from "../components/AP.tsx"
import Error from "../components/Error.tsx"
import Loading, { type loadingData } from "../components/Loading.tsx"
import Option, { type optionData } from "../components/Option.tsx"
import { type bettercapWifiAPData, getBettercapWifi } from "../lib/api/bettercap.ts"

export const [bettercapWifi, setBettercapWifi] = createSignal<bettercapWifiAPData[]>([])

const Scan: Component = () => {
	createPolled(async () => setBettercapWifi(await getBettercapWifi()), 5000)

	const [loading, setLoading] = createSignal<loadingData>({ title: "Loading", pending: false, msg: "" })
	const [option, setOption] = createSignal<optionData>({ open: false, title: "Attack", message: "", children: null, buttonFirst: null, buttonSecond: null })

	return (
		<div id="scan">
			<Option data={option()} />
			<Show when={bettercapWifi().length !== 0} fallback={<Error title="No data!" msg="You probably don't have any devices in monitoring mode." />}>
				<For each={bettercapWifi() as bettercapWifiAPData[]}>
					{(ap, _) => <AP ap={ap} option={option} setOption={setOption} setLoading={setLoading} />}
				</For>
			</Show>
			<Loading data={loading} setData={setLoading} />
		</div>
	)
}

export default Scan
