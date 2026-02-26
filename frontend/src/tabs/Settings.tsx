import { type Component, createEffect } from "solid-js"
import { createSignal, For, onMount } from "solid-js"

import "./Settings.css"
import { createStore } from "solid-js/store"
import type { loadingData } from "../components/Loading.tsx"
import Loading from "../components/Loading.tsx"
import SettingsInterface from "../components/SettingsInterface.tsx"
import { getPortals } from "../lib/api/portals.ts"
import type { settingsData, settingsInterfaceFieldsData } from "../lib/api/settings.ts"
import { getSettings, saveSettings } from "../lib/api/settings.ts"
import type { response } from "../lib/var.ts"

export const [settings, setSettings] = createStore<settingsData>({ iface: {} } as settingsData)
export const [settingsModes, setSettingsModes] = createSignal<string[]>([])
export const [portals, setPortals] = createSignal<string[]>([])

const Settings: Component = () => {
	onMount(async () => {
		setPortals(await getPortals())
		setSettings(await getSettings())
	})

	const [loading, setLoading] = createSignal<loadingData>({ title: "loading", pending: false, msg: "" })

	createEffect(() => {
		let modes = []
		for (const iface in settings.iface) {
			if (settings.iface[iface].mode === "none") continue
			modes.push(settings.iface[iface].mode)
		}
		setSettingsModes(modes)
	})

	return (
		<div id="settings">
			<For each={Object.entries(settings.iface) as [string, settingsInterfaceFieldsData][]}>
				{(iface, _) => <SettingsInterface name={iface[0]} settings={settings} setSettings={setSettings} settingsModes={settingsModes} />}
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
						const rsp: response = await saveSettings(settings)

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
