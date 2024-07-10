import type {Component} from "solid-js"
import {createEffect, createSignal, Index, Show} from "solid-js"
import {settingsInterfaceFieldsData} from "../lib/settings"
import {Field, NumberInput, Select} from "@ark-ui/solid"
import {Portal} from "solid-js/web"

import "./SettingsInterface.css"

interface settingsInterfaceProps {
    name: string
    iface: settingsInterfaceFieldsData
}

const SettingsInterface: Component<settingsInterfaceProps> = (props) => {

    const modes = ["none", "hotspot", "monitor"]

    const [mode, setMode] = createSignal<string>(props.iface.mode)

    createEffect(() => {
        props.iface.mode = mode()
    })

    return (
        <div class="settings-iface card">
            <h2>{props.name}</h2>
            <Select.Root items={modes} required={true} value={[props.iface.mode ? props.iface.mode : "none"]} onValueChange={(e) => setMode(e.value[0])}>
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
            <Show when={mode() === "hotspot"<boolean>}>
                <div class="settings-iface-hotspot">
                    <Field.Root>
                        <Field.Label>SSID</Field.Label>
                        <Field.Input placeholder={"PiFi"} onInput={(e) => props.iface.ssid = e.currentTarget.value}/>
                        <Field.ErrorText>Error Info</Field.ErrorText>
                    </Field.Root>
                    <Field.Root>
                        <Field.Label>Password</Field.Label>
                        <Field.Input placeholder={"none"} onInput={(e) => props.iface.password = e.currentTarget.value}/>
                        <Field.ErrorText>Error Info</Field.ErrorText>
                    </Field.Root>
                    <NumberInput.Root value={"1"} min={1} max={14} onValueChange={(e) => props.iface.channel = e.valueAsNumber}>
                        <NumberInput.Label>Channel</NumberInput.Label>
                        <NumberInput.Input/>
                        <NumberInput.Control>
                            <NumberInput.DecrementTrigger>-</NumberInput.DecrementTrigger>
                            <NumberInput.IncrementTrigger>+</NumberInput.IncrementTrigger>
                        </NumberInput.Control>
                    </NumberInput.Root>
                </div>
            </Show>
        </div>
    )
}

export default SettingsInterface