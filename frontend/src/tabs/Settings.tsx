import type {Component} from "solid-js"
import {createSignal, For, onMount} from "solid-js"

import "./Settings.css"
import {getSettings, saveSettings, settingsData, settingsInterfaceFieldsData} from "../lib/settings"
import Loading, {loadingData} from "../components/Loading"
import SettingsInterface from "../components/SettingsInterface"
import {response} from "../lib/var";
import {getPortals} from "../lib/portals";

export const [settings, setSettings] = createSignal<settingsData>({iface: {}})
export const [portals, setPortals] = createSignal<string[]>([])
export const [settingsInterfaceHotspot, setSettingsInterfaceHotspot] = createSignal<string>("")

const Settings: Component = () => {

    onMount(async () => {
        setPortals(await getPortals())
        setSettings(await getSettings())
    })

    const [loading, setLoading] = createSignal<loadingData>({title: "loading", pending: false, msg: ""})

    return (
        <div id="settings">
            <For each={Object.entries(settings().iface)<[string, settingsInterfaceFieldsData][]>}>{(iface, i) =>
                <SettingsInterface name={iface[0]} iface={iface[1]}/>
            }
            </For>
            <div id="settings-btn">
                <button id="settings-revert" class="card" onClick={async () => {
                    setLoading({title: "Reverting", pending: true, msg: ""})
                    setSettings(await getSettings())
                    setLoading({title: "Reverting", pending: false, msg: ""})

                }}>revert
                </button>
                <button id="settings-save" class="card green" onClick={async () => {
                    setLoading({title: "saving", pending: true, msg: ""})
                    let req: response = await saveSettings(settings())

                    if (req.message != "") {
                        setLoading({title: "Saving", pending: true, msg: req.message})
                        return
                    }

                    setSettings(req.data as settingsData)
                    setLoading({title: "Saving", pending: false, msg: ""})
                }}>save
                </button>
            </div>
            <Loading data={loading()}/>
        </div>
    )
}

export default Settings