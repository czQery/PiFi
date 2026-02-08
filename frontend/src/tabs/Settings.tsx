import type { Component } from "solid-js"
import { createSignal, For, onMount } from "solid-js"

import "./Settings.css"
import type { loadingData } from "../components/Loading.tsx"
import Loading from "../components/Loading.tsx"
import SettingsInterface from "../components/SettingsInterface.tsx"
import { getPortals } from "../lib/api/portals.ts"
import type { settingsData, settingsInterfaceFieldsData } from "../lib/api/settings.ts"
import { getSettings, saveSettings } from "../lib/api/settings.ts"
import type { response } from "../lib/var.ts"

export const [settings, setSettings] = createSignal<settingsData>({ iface: {} })
export const [portals, setPortals] = createSignal<string[]>([])
export const [settingsInterfaceHotspot, setSettingsInterfaceHotspot] = createSignal<string>("")
export const [settingsInterfaceClient, setSettingsInterfaceClient] = createSignal<string>("")

const Settings: Component = () => {
	onMount(async () => {
		setPortals(await getPortals())
		setSettings(await getSettings())
	})

	const [loading, setLoading] = createSignal<loadingData>({ title: "loading", pending: false, msg: "" })

	return (
		<div id="settings">
			<For each={Object.entries(settings().iface) as [string, settingsInterfaceFieldsData][]}>
				{(iface, _) => <SettingsInterface name={iface[0]} iface={iface[1]} />}
			</For>
			<div id="settings-btn">
				<button
					id="settings-revert"
					class="card"
					onClick={async () => {
						setLoading({ title: "Reverting", pending: true, msg: "" })
						setSettings(await getSettings())
						setLoading({ title: "Reverting", pending: false, msg: "" })
					}}
				>
					revert
				</button>
				<button
					id="settings-save"
					class="card green"
					onClick={async () => {
						setLoading({ title: "saving", pending: true, msg: "" })
						const rsp: response = await saveSettings(settings())

						if (rsp.message != "") {
							setLoading({ title: "Saving", pending: true, msg: rsp.message })
							return
						}

						setSettings(rsp.data as settingsData)
						setLoading({ title: "Saving", pending: false, msg: "" })
					}}
				>
					save
				</button>
			</div>
			<Loading data={loading()} />
		</div>
	)
}

export default Settings
