import type {Component} from "solid-js"
import {createSignal, For, onMount} from "solid-js"

import "./Settings.css"
import {getSettings, settingsData, settingsInterfaceFieldsData} from "../lib/settings"
import Loading, {loadingData} from "../components/Loading"
import SettingsInterface from "../components/SettingsInterface"

export const [settings, setSettings] = createSignal<settingsData>({iface: {}})

const Settings: Component = () => {

    onMount(async () => {
        setSettings(await getSettings())
    })
    const [loading, setLoading] = createSignal<loadingData>({title: "loading", pending: false})

    return (
        <div id="settings">
            <For each={Object.entries(settings().iface)<[string, settingsInterfaceFieldsData][]>}>{(iface, i) =>
                <SettingsInterface name={iface[0]} iface={iface[1]}/>
            }
            </For>
            <div id="settings-btn">
                <button id="settings-revert" class="card" onClick={async () => {
                    setLoading({title: "reverting", pending: true})
                    setSettings(await getSettings())
                    setLoading({title: "reverting", pending: false})

                }}>revert
                </button>
                <button id="settings-save" class="card green" onClick={async () => {
                    setLoading({title: "saving", pending: true})
                    //setSettings(await saveSettings(settings()))
                    console.log(settings())
                    setLoading({title: "saving", pending: false})
                }}>save
                </button>
            </div>
            <Loading data={loading()}/>
        </div>
    )
}

export default Settings