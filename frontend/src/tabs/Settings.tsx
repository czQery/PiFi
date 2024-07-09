import type {Component} from "solid-js";
import {createSignal, For, Index, onMount} from "solid-js";

import "./Settings.css"
import {getSettings, saveSettings, settingsData, settingsInterfaceFieldsData} from "../lib/settings";
import {Select} from "@ark-ui/solid"
import {Portal} from "solid-js/web";
import Loading, {loadingData} from "../components/Loading";

export const [settings, setSettings] = createSignal<settingsData>({iface: {}});

const Settings: Component = () => {

    onMount(async () => {
        setSettings(await getSettings())
    })

    const modes = ["none", "hotspot", "monitor"]
    const [loading, setLoading] = createSignal<loadingData>({title: "loading", pending: false})

    return (
        <div id="settings">
            <For each={Object.entries(settings().iface)<[string, settingsInterfaceFieldsData][]>}>{(iface, i) =>
                <div class="settings-iface card">
                    <h2>{iface[0]}</h2>
                    <Select.Root items={modes} required={true} value={[iface[1].mode ? iface[1].mode : "none"]}>
                        <Select.Label>Mode</Select.Label>
                        <Select.Control>
                            <Select.Trigger>
                                <Select.ValueText placeholder="select"/>
                                <Select.Indicator>▼</Select.Indicator>
                            </Select.Trigger>
                        </Select.Control>
                        <Portal>
                            <Select.Positioner>
                                <Select.Content>
                                    <Select.ItemGroup id="test">
                                        <Index each={modes<string[]>}>{(item, i) => (
                                            <Select.Item item={item()}>
                                                <Select.ItemText>{item()}</Select.ItemText>
                                            </Select.Item>
                                        )}
                                        </Index>
                                    </Select.ItemGroup>
                                </Select.Content>
                            </Select.Positioner>
                        </Portal>
                    </Select.Root>
                </div>
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
                    await saveSettings()
                    setSettings(await getSettings())
                    setLoading({title: "saving", pending: false})
                }}>save
                </button>
            </div>
            <Loading data={loading()}/>
        </div>
    )
}

export default Settings