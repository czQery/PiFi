import type {Component} from "solid-js";
import {createSignal, For, Index, onMount} from "solid-js";

import "./Settings.css"
import {getSettings, settingsData, settingsInterfaceFieldsData} from "../lib/settings";
import {Select} from "@ark-ui/solid"
import {Portal} from "solid-js/web";

export const [settings, setSettings] = createSignal<settingsData>({iface: {}});

const Settings: Component = () => {

    onMount(async () => {
        setSettings(await getSettings())
    })

    const items = ["none", "hotspot", "monitor"]

    return (
        <div id="settings">
            <For each={Object.entries(settings().iface)<[string, settingsInterfaceFieldsData][]>}>{(iface, i) =>
                <div class="settings-iface card">
                    <h2>{iface[0]}</h2>
                    <Select.Root items={items} required={true} value={[iface[1].mode ? iface[1].mode : "none"]}>
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
                                        <Index each={items<string[]>}>{(item,i) => (
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
        </div>
    )
}

export default Settings