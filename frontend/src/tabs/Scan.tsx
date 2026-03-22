import "./Scan.css"
import { createPolled } from "@solid-primitives/timer"
import { type Component, createSignal, For, onMount, Show } from "solid-js"
import AP from "../components/AP.tsx"
import DialogLoading, { type dialogLoadingData } from "../components/dialog/DialogLoading.tsx"
import DialogWrapper, { type dialogWrapperData } from "../components/dialog/DialogWrapper.tsx"
import Error from "../components/Error.tsx"
import InputBool from "../components/input/InputBool.tsx"
import { getAttackDeauth, setAttackDeauth } from "../lib/api/attack.ts"
import { type bettercapWifiAPData, getBettercapWifi } from "../lib/api/bettercap.ts"

export const [bettercapWifi, setBettercapWifi] = createSignal<bettercapWifiAPData[]>([])
export const [attackDeauthList, setAttackDeauthList] = createSignal<string[]>([])

const Scan: Component = () => {
	createPolled(async () => setBettercapWifi(await getBettercapWifi()), 5000)

	onMount(async () => {
		setAttackDeauthList(await getAttackDeauth())
	})

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
			<DialogWrapper data={option()}>
				<InputBool
					name="Deauth"
					value={attackDeauthList()?.includes(option().helper) ?? false}
					callback={async e => {
						setAttackDeauth(e.checked, option().helper as string).then()
						setAttackDeauthList(await getAttackDeauth())
					}}
				/>
			</DialogWrapper>
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
