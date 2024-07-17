import type {Component} from "solid-js"
import {createEffect, createSignal, Index, onMount, Show} from "solid-js"
import {settingsInterfaceFieldsData} from "../lib/settings"
import {Checkbox, Field, NumberInput, Select} from "@ark-ui/solid"
import {Portal} from "solid-js/web"

import "./SettingsInterface.css"
import {LucideUnplug} from "lucide-solid";
import {setSettingsInterfaceHotspot, settingsInterfaceHotspot} from "../tabs/Settings";

interface settingsInterfaceProps {
    name: string
    iface: settingsInterfaceFieldsData
}

const SettingsInterface: Component<settingsInterfaceProps> = (props) => {

    const modes: string[] = ["none", "hotspot", "monitor"]
    const [mode, setMode] = createSignal<string>(props.iface.mode)

    const portals: string[] = ["test", "nn"]

    onMount(() => {
        if (props.iface.channel === 0 || props.iface.channel > 14) {
            props.iface.channel = 1
        }

        if (!props.iface.portal_source) {
            props.iface.portal_source = portals[0]
        }
    })

    createEffect(() => {
        props.iface.mode = mode()
        if (mode() === "hotspot") {
            setSettingsInterfaceHotspot(props.name)
        } else if (settingsInterfaceHotspot() === props.name) {
            setSettingsInterfaceHotspot("")
        }
    })

    return (
        <div class="settings-iface card">
            <div class="settings-iface-title">
                <h2>{props.name}</h2>
                <Show when={!props.iface.ready<boolean>}>
                    <LucideUnplug class="card"/>
                </Show>
            </div>
            <Select.Root items={modes} required={true} immediate={true} value={[props.iface.mode ? props.iface.mode : "none"]} onValueChange={(e) => setMode(e.value[0])}>
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
                                    <Show when={settingsInterfaceHotspot() !== props.name && settingsInterfaceHotspot() !== "" && item() === "hotspot"<boolean>} fallback={
                                        <Select.Item item={item()}>
                                            <Select.ItemText>{item()}</Select.ItemText>
                                        </Select.Item>
                                    }>
                                        <Select.Item item={item()} data-disabled aria-disabled disabled>
                                            <Select.ItemText data-disabled aria-disabled>{item()}</Select.ItemText>
                                        </Select.Item>
                                    </Show>
                                )}
                                </Index>
                            </Select.ItemGroup>
                        </Select.Content>
                    </Select.Positioner>
                </Portal>
            </Select.Root>
            <Show when={mode() === "hotspot"<boolean>}>
                <div class="settings-iface-hotspot">
                    <Field.Root>
                        <Field.Label>SSID</Field.Label>
                        <Field.Input placeholder={"PiFi"} value={props.iface.ssid} onInput={(e) => props.iface.ssid = e.currentTarget.value}/>
                        <Field.ErrorText>Error Info</Field.ErrorText>
                    </Field.Root>
                    <Field.Root>
                        <Field.Label>Password</Field.Label>
                        <Field.Input placeholder={"none"} value={props.iface.password} onInput={(e) => props.iface.password = e.currentTarget.value}/>
                        <Field.ErrorText>Error Info</Field.ErrorText>
                    </Field.Root>
                    <NumberInput.Root value={props.iface.channel.toString()} min={1} max={14} onValueChange={(e) => props.iface.channel = e.valueAsNumber}>
                        <NumberInput.Label>Channel</NumberInput.Label>
                        <NumberInput.Input/>
                        <NumberInput.Control>
                            <NumberInput.DecrementTrigger>-</NumberInput.DecrementTrigger>
                            <NumberInput.IncrementTrigger>+</NumberInput.IncrementTrigger>
                        </NumberInput.Control>
                    </NumberInput.Root>
                </div>
                <div class="settings-iface-hotspot">
                    <Checkbox.Root checked={props.iface.portal} onCheckedChange={(e) => props.iface.portal = e.checked as boolean}>
                        <Checkbox.Label>Portal</Checkbox.Label>
                        <Checkbox.Control>
                            <div>
                                <span></span>
                                <div></div>
                            </div>
                        </Checkbox.Control>
                        <Checkbox.HiddenInput/>
                    </Checkbox.Root>
                    <Select.Root items={portals} required={true} immediate={true} value={[props.iface.portal_source ? props.iface.portal_source : portals[0]]} onValueChange={(e) => props.iface.portal_source = e.value[0]}>
                        <Select.Label>Portal source</Select.Label>
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
                                        <Index each={portals<string[]>}>{(item, i) => (
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
            </Show>
        </div>
    )
}

export default SettingsInterface