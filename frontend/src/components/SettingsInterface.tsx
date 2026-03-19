import type {ListCollection} from "@ark-ui/solid"
import {Checkbox, createListCollection, Select} from "@ark-ui/solid"
import {type Accessor, type Component, Show} from "solid-js"
import {Index, Portal} from "solid-js/web"

import type {settingsData} from "../lib/api/settings.ts"

import "./SettingsInterface.css"

import {LucideUnplug} from "lucide-solid"
import {produce, type SetStoreFunction} from "solid-js/store"
import {portals, settingsModes} from "../tabs/Settings.tsx"
import InputBool from "./input/InputBool.tsx"
import InputNumber from "./input/InputNumber.tsx"
import InputString from "./input/InputString.tsx"

interface settingsInterfaceProps {
	name: string
	settings: settingsData
	setSettings: SetStoreFunction<settingsData>
	settingsModes: Accessor<string[]>
}

const channelMin = 1
const channelMax = 13

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
				deselectable={false}
				multiple={false}
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
					<InputString
						name="SSID"
						value={props.settings.hotspot.ssid}
						placeholder="PiFi"
						callback={e => {
							props.setSettings(produce(state => {
								state.hotspot.ssid = e.currentTarget.value
							}))
						}}
					/>
					<InputString
						name="Password"
						value={props.settings.hotspot.password}
						placeholder="none"
						callback={e => {
							props.setSettings(produce(state => {
								state.hotspot.password = e.currentTarget.value
							}))
						}}
					/>
					<InputNumber
						name="Channel"
						value={props.settings.hotspot.channel}
						min={channelMin}
						max={channelMax}
						callback={e => {
							props.setSettings(produce(state => {
								state.hotspot.channel = e.valueAsNumber
							}))
						}}
					/>
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
						deselectable={true}
						multiple={false}
						value={[props.settings.hotspot.portal_source]}
						onValueChange={e => {
							props.setSettings(produce(state => {
								state.hotspot.portal_source = e.value[0] ?? ""
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
					<InputString
						name="SSID"
						value={props.settings.client.ssid}
						placeholder="PiFi"
						callback={e => {
							props.setSettings(produce(state => {
								state.client.ssid = e.currentTarget.value
							}))
						}}
					/>
					<InputString
						name="Password"
						value={props.settings.client.password}
						placeholder="none"
						callback={e => {
							props.setSettings(produce(state => {
								state.client.password = e.currentTarget.value
							}))
						}}
					/>
				</div>
			</Show>

			<Show when={props.settings.iface[props.name].mode === "monitor"}>
				<div class="settings-iface-monitor">
					<InputBool
						name="Wardrive"
						value={props.settings.monitor.wardrive}
						callback={e => {
							props.setSettings(produce(state => {
								state.monitor.wardrive = e.checked as boolean
							}))
						}}
					/>
				</div>
				<div class="settings-iface-monitor">
					<InputBool
						name="Channel-hop"
						value={props.settings.monitor.channel_hop}
						callback={e => {
							props.setSettings(produce(state => {
								state.monitor.channel_hop = e.checked as boolean
							}))
						}}
					/>
					<Show when={!props.settings.monitor.channel_hop}>
						<InputNumber
							name="Channel"
							value={props.settings.monitor.channel}
							min={channelMin}
							max={channelMax}
							callback={e => {
								props.setSettings(produce(state => {
									state.monitor.channel = e.valueAsNumber
								}))
							}}
						/>
					</Show>
				</div>
			</Show>
		</div>
	)
}

export default SettingsInterface
