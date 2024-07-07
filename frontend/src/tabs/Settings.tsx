import type {Component} from "solid-js";
import {createSignal, For, Index, onMount} from "solid-js";

import "./Settings.css"
import {getSettings, saveSettings, settingsData, settingsInterfaceFieldsData} from "../lib/settings";
import {Dialog, Select, Progress} from "@ark-ui/solid"
import {Portal} from "solid-js/web";

export const [settings, setSettings] = createSignal<settingsData>({iface: {}});

const Settings: Component = () => {

    onMount(async () => {
        setSettings(await getSettings())
    })

    const items = ["none", "hotspot", "monitor"]
    const [isOpen, setIsOpen] = createSignal(false)

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
                                        <Index each={items<string[]>}>{(item, i) => (
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
                    setSettings(await getSettings())
                }}>revert
                </button>
                <button id="settings-save" class="card green" onClick={async () => {
                    setIsOpen(true)
                    await saveSettings()
                }}>save
                </button>
                <Dialog.Root className={"card"} open={isOpen()} onOpenChange={(e) => setIsOpen(e.open)}>
                    <Portal>
                        <Dialog.Backdrop/>
                        <Dialog.Positioner>
                            <Dialog.Content>
                                <Progress.Root>
                                    <Progress.Label>saving</Progress.Label>
                                    <Progress.Track>
                                        <Progress.Range/>
                                    </Progress.Track>
                                </Progress.Root>
                            </Dialog.Content>
                        </Dialog.Positioner>
                    </Portal>
                </Dialog.Root>
            </div>
        </div>
    )
}

export default Settings