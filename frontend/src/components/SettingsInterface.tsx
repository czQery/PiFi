import type { ListCollection } from "@ark-ui/solid"
import { Checkbox, createListCollection, Field, NumberInput, Select } from "@ark-ui/solid"
import { type Accessor, type Component, Show } from "solid-js"
import { Index, Portal } from "solid-js/web"

import type { settingsData } from "../lib/api/settings.ts"

import "./SettingsInterface.css"

import { LucideUnplug } from "lucide-solid"
import { produce, type SetStoreFunction } from "solid-js/store"
import { portals, settingsModes } from "../tabs/Settings.tsx"

interface settingsInterfaceProps {
	name: string
	settings: settingsData
	setSettings: SetStoreFunction<settingsData>
	settingsModes: Accessor<string[]>
}

const SettingsInterface: Component<settingsInterfaceProps> = props => {
	const portalsCollection: ListCollection<string> = createListCollection({ items: portals() })
	const modesCollection: ListCollection<string> = createListCollection({ items: ["none", "hotspot", "client", "monitor"] })

	return (
		<div class="settings-iface card">
			<div class="settings-iface-title">
				<h2>{props.name}</h2>
				<Show when={!props.settings.iface[props.name].ready}>
					<LucideUnplug class="card" />
				</Show>
			</div>
			<Select.Root
				required={true}
				immediate={true}
				value={[props.settings.iface[props.name].mode]}
				onValueChange={e => {
					props.setSettings(produce(state => {
						state.iface[props.name].mode = e.value[0]
					}))
				}}
				collection={modesCollection}
			>
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
							<Select.ItemGroup>
								<Index each={modesCollection.items}>
									{item => (
										<Show
											when={!settingsModes().includes(item()) || props.settings.iface[props.name].mode === item()}
											fallback={
												/*@ts-ignore*/


													<Select.Item item={item()} data-disabled aria-disabled disabled>
														<Select.ItemText data-disabled aria-disabled>{item()}</Select.ItemText>
													</Select.Item>

											}
										>
											<Select.Item item={item()}>
												<Select.ItemText>{item()}</Select.ItemText>
											</Select.Item>
										</Show>
									)}
								</Index>
							</Select.ItemGroup>
						</Select.Content>
					</Select.Positioner>
				</Portal>
			</Select.Root>
			<Show when={props.settings.iface[props.name].mode === "hotspot"}>
				<div class="settings-iface-hotspot">
					<Field.Root>
						<Field.Label>SSID</Field.Label>
						<Field.Input
							placeholder={"PiFi"}
							value={props.settings.hotspot.ssid}
							onInput={e => {
								props.setSettings(produce(state => {
									state.hotspot.ssid = e.currentTarget.value
								}))
							}}
						/>
						<Field.ErrorText>Error Info</Field.ErrorText>
					</Field.Root>
					<Field.Root>
						<Field.Label>Password</Field.Label>
						<Field.Input
							placeholder={"none"}
							value={props.settings.hotspot.password}
							onInput={e => {
								props.setSettings(produce(state => {
									state.hotspot.password = e.currentTarget.value
								}))
							}}
						/>
						<Field.ErrorText>Error Info</Field.ErrorText>
					</Field.Root>
					<NumberInput.Root
						value={props.settings.hotspot.channel.toString()}
						min={1}
						max={14}
						onValueChange={e => {
							props.setSettings(produce(state => {
								state.hotspot.channel = e.valueAsNumber
							}))
						}}
					>
						<NumberInput.Label>Channel</NumberInput.Label>
						<NumberInput.Input />
						<NumberInput.Control>
							<NumberInput.DecrementTrigger>-</NumberInput.DecrementTrigger>
							<NumberInput.IncrementTrigger>+</NumberInput.IncrementTrigger>
						</NumberInput.Control>
					</NumberInput.Root>
				</div>
				<div class="settings-iface-hotspot">
					<Checkbox.Root
						checked={props.settings.hotspot.portal}
						onCheckedChange={e => {
							props.setSettings(produce(state => {
								state.hotspot.portal = e.checked as boolean
							}))
						}}
					>
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
						value={[props.settings.hotspot.portal_source]}
						onValueChange={e => {
							props.setSettings(produce(state => {
								state.hotspot.portal_source = e.value[0]
							}))
						}}
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
									<Select.ItemGroup>
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
			<Show when={props.settings.iface[props.name].mode === "client"}>
				<div class="settings-iface-client">
					<Field.Root>
						<Field.Label>SSID</Field.Label>
						<Field.Input
							placeholder={"PiFi"}
							value={props.settings.client.ssid}
							onInput={e => {
								props.setSettings(produce(state => {
									state.client.ssid = e.currentTarget.value
								}))
							}}
						/>
						<Field.ErrorText>Error Info</Field.ErrorText>
					</Field.Root>
					<Field.Root>
						<Field.Label>Password</Field.Label>
						<Field.Input
							placeholder={"none"}
							value={props.settings.client.password}
							onInput={e => {
								props.setSettings(produce(state => {
									state.client.password = e.currentTarget.value
								}))
							}}
						/>
						<Field.ErrorText>Error Info</Field.ErrorText>
					</Field.Root>
				</div>
			</Show>

			<Show when={props.settings.iface[props.name].mode === "monitor"}>
				<div class="settings-iface-monitor">
					<Checkbox.Root
						checked={props.settings.monitor.deauth}
						onCheckedChange={e => {
							props.setSettings(produce(state => {
								state.monitor.deauth = e.checked as boolean
							}))
						}}
					>
						<Checkbox.Label>Deauth</Checkbox.Label>
						<Checkbox.Control>
							<div>
								<span></span>
								<div></div>
							</div>
						</Checkbox.Control>
						<Checkbox.HiddenInput />
					</Checkbox.Root>
					<Checkbox.Root
						checked={props.settings.monitor.assoc}
						onCheckedChange={e => {
							props.setSettings(produce(state => {
								state.monitor.assoc = e.checked as boolean
							}))
						}}
					>
						<Checkbox.Label>Assoc</Checkbox.Label>
						<Checkbox.Control>
							<div>
								<span></span>
								<div></div>
							</div>
						</Checkbox.Control>
						<Checkbox.HiddenInput />
					</Checkbox.Root>
				</div>
			</Show>
		</div>
	)
}

export default SettingsInterface
