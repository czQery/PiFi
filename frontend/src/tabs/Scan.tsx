import "./Scan.css"
import { createPolled } from "@solid-primitives/timer"
import { type Component, createEffect, createSignal, For, onMount, Show } from "solid-js"
import AP from "../components/AP.tsx"
import DialogLoading, { type dialogLoadingData } from "../components/dialog/DialogLoading.tsx"
import DialogWrapper, { type dialogWrapperData } from "../components/dialog/DialogWrapper.tsx"
import Error from "../components/Error.tsx"
import InputBool from "../components/input/InputBool.tsx"
import InputNumber from "../components/input/InputNumber.tsx"
import { channelMax, channelMin } from "../components/SettingsInterface.tsx"
import { getAttackChannelSwitch, getAttackDeauth, setAttackChannelSwitch, setAttackDeauth } from "../lib/api/attack.ts"
import { type bettercapWifiAPData, getBettercapWifi } from "../lib/api/bettercap.ts"

export const [bettercapWifi, setBettercapWifi] = createSignal<bettercapWifiAPData[]>([])
export const [attackDeauthList, setAttackDeauthList] = createSignal<string[]>([])
export const [attackChannelSwitchList, setAttackChannelSwitchList] = createSignal<Record<string, number>>({})

const Scan: Component = () => {
	createPolled(async () => setBettercapWifi(await getBettercapWifi()), 5000)

	onMount(async () => {
		setAttackDeauthList(await getAttackDeauth())
		setAttackChannelSwitchList(await getAttackChannelSwitch())
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

	const [attackChannelSwitchChannel, setAttackChannelSwitchChannel] = createSignal<number>(channelMin)
	createEffect(() => {
		setAttackChannelSwitchChannel(attackChannelSwitchList()[option().helper as string] ?? channelMin)
	})

	return (
		<div id="scan">
			<DialogWrapper data={option()}>
				<div class="scan-flex">
					<InputBool
						name="Deauth"
						value={attackDeauthList()?.includes(option().helper) ?? false}
						callback={async e => {
							await setAttackDeauth(e.checked, option().helper as string)
							setAttackDeauthList(await getAttackDeauth())
						}}
					/>
				</div>
				<div class="scan-flex">
					<InputBool
						name="Channel-switch"
						value={!!attackChannelSwitchList()[option().helper as string]}
						callback={async e => {
							await setAttackChannelSwitch(e.checked, option().helper as string, attackChannelSwitchChannel())
							setAttackChannelSwitchList(await getAttackChannelSwitch())
						}}
					/>
					<InputNumber
						name="Channel"
						value={attackChannelSwitchChannel()}
						min={channelMin}
						max={channelMax}
						callback={async e => {
							if (!!attackChannelSwitchList()[option().helper as string]) {
								await setAttackChannelSwitch(true, option().helper as string, e.valueAsNumber)
								setAttackChannelSwitchList(await getAttackChannelSwitch())
							} else {
								setAttackChannelSwitchChannel(e.valueAsNumber)
							}
						}}
					/>
				</div>
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
