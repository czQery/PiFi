import type { ListCollection } from "@ark-ui/solid"
import { Checkbox, createListCollection, Field, NumberInput, Select } from "@ark-ui/solid"
import type { Component } from "solid-js"
import { createEffect, createSignal, Show } from "solid-js"
import { Index, Portal } from "solid-js/web"

import type { settingsInterfaceFieldsData } from "../lib/settings.ts"

import "./SettingsInterface.css"

import { LucideUnplug } from "lucide-solid"
import { portals, setSettingsInterfaceHotspot, settingsInterfaceHotspot } from "../tabs/Settings.tsx"

interface settingsInterfaceProps {
	name: string
	iface: settingsInterfaceFieldsData
}

const SettingsInterface: Component<settingsInterfaceProps> = props => {
	const portalsCollection: ListCollection<string> = createListCollection({ items: portals() })
	const [portal, setPortal] = createSignal<boolean>(props.iface.portal)
	const [portalSource, setPortalSource] = createSignal<string>(props.iface.portal_source ? props.iface.portal_source : portalsCollection.items[0])

	const modesCollection: ListCollection<string> = createListCollection({ items: ["none", "hotspot", "monitor"] })
	const [mode, setMode] = createSignal<string>(props.iface.mode ? props.iface.mode : modesCollection.items[0])

	const [channel, setChannel] = createSignal<number>(props.iface.channel === 0 || props.iface.channel > 14 ? 1 : props.iface.channel)

	createEffect(() => {
		props.iface.mode = mode()
		props.iface.portal = portal()
		props.iface.portal_source = portalSource()
		props.iface.channel = channel()
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
				<Show when={!props.iface.ready}>
					<LucideUnplug class="card" />
				</Show>
			</div>
			<Select.Root required={true} immediate={true} value={[mode()]} onValueChange={e => setMode(e.value[0])} collection={modesCollection}>
				<Select.Label>Mode</Select.Label>
				<Select.Control>
					<Select.Trigger>
						<Select.ValueText placeholder="select" />
						<Select.Indicator>▼</Select.Indicator>
					</Select.Trigger>
				</Select.Control>
				<Portal>
					<Select.Positioner>
						<Select.Content>
							<Select.ItemGroup id="test">
								<Index each={modesCollection.items}>
									{item => (
										<Show
											when={settingsInterfaceHotspot() !== props.name && settingsInterfaceHotspot() !== "" && item() === "hotspot"}
											fallback={
												<Select.Item item={item()}>
													<Select.ItemText>{item()}</Select.ItemText>
												</Select.Item>
											}
										>
											{/*@ts-ignore*/}
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
			<Show when={mode() === "hotspot"}>
				<div class="settings-iface-hotspot">
					<Field.Root>
						<Field.Label>SSID</Field.Label>
						<Field.Input placeholder={"PiFi"} value={props.iface.ssid} onInput={e => (props.iface.ssid = e.currentTarget.value)} />
						<Field.ErrorText>Error Info</Field.ErrorText>
					</Field.Root>
					<Field.Root>
						<Field.Label>Password</Field.Label>
						<Field.Input placeholder={"none"} value={props.iface.password} onInput={e => (props.iface.password = e.currentTarget.value)} />
						<Field.ErrorText>Error Info</Field.ErrorText>
					</Field.Root>
					<NumberInput.Root value={channel().toString()} min={1} max={14} onValueChange={e => (setChannel(e.valueAsNumber))}>
						<NumberInput.Label>Channel</NumberInput.Label>
						<NumberInput.Input />
						<NumberInput.Control>
							<NumberInput.DecrementTrigger>-</NumberInput.DecrementTrigger>
							<NumberInput.IncrementTrigger>+</NumberInput.IncrementTrigger>
						</NumberInput.Control>
					</NumberInput.Root>
				</div>
				<div class="settings-iface-hotspot">
					<Checkbox.Root checked={portal()} onCheckedChange={e => (setPortal(e.checked as boolean))}>
						<Checkbox.Label>Portal</Checkbox.Label>
						<Checkbox.Control>
							<div>
								<span></span>
								<div></div>
							</div>
						</Checkbox.Control>
						<Checkbox.HiddenInput />
					</Checkbox.Root>
					<Select.Root
						required={true}
						immediate={true}
						value={[portalSource()]}
						onValueChange={e => setPortalSource(e.value[0])}
						collection={portalsCollection}
					>
						<Select.Label>Portal source</Select.Label>
						<Select.Control>
							<Select.Trigger>
								<Select.ValueText placeholder="select" />
								<Select.Indicator>▼</Select.Indicator>
							</Select.Trigger>
						</Select.Control>
						<Portal>
							<Select.Positioner>
								<Select.Content>
									<Select.ItemGroup id="test">
										<Index each={portalsCollection.items}>
											{item => (
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
