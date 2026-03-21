import "./AP.css"
import { LucideWifi } from "lucide-solid"
import { type Accessor, type Component, For, type Setter, Show } from "solid-js"
import type { bettercapWifiAPData, bettercapWifiClientData } from "../lib/api/bettercap.ts"
import { addZeroDecimal, parseAgo } from "../lib/other.ts"
import type { loadingData } from "./Loading.tsx"
import type { optionData } from "./Option.tsx"

interface APProps {
	ap: bettercapWifiAPData
	option: Accessor<optionData>
	setOption: Setter<optionData>
	setLoading: Setter<loadingData>
}

const AP: Component<APProps> = props => {
	return (
		<div class="ap card">
			<div class="ap-flex" style="justify-content: space-between;">
				<div class="ap-flex">
					<LucideWifi />
					<span class="title">{props.ap.hostname}</span>
				</div>
				<div class="ap-flex">
					<span>[RSSI:</span>
					<span class="value mono">{props.ap.rssi}</span>
					<span>]</span>
					<span>[CH:</span>
					<span class="value mono">{props.ap.channel}</span>
					<span>]</span>
				</div>
			</div>
			<div class="ap-flex" style="justify-content: space-between;">
				<div class="ap-flex">
					<span>BSSID:</span>
					<span class="value mono">{props.ap.mac}</span>
				</div>
				<div class="ap-flex">
					<span>[</span>
					<span style={{ color: props.ap.encryption === "OPEN" ? "var(--green)" : "var(--pink)" }}>{props.ap.encryption}</span>
					<span>]</span>
				</div>
			</div>
			<div class="ap-flex">
				<span>VENDOR:</span>
				<span class="value">{props.ap.vendor !== "" ? props.ap.vendor : "unknown"}</span>
			</div>
			<div class="ap-clients">
				<hr class="h" style="margin: 2.5px;" />
				<Show when={props.ap.clients.length !== 0} fallback={<span>No clients!</span>}>
					<For each={props.ap.clients as bettercapWifiClientData[]}>
						{(client, _) => (
							<div class="ap-flex" style="justify-content: space-between;">
								<span class="value mono">{client.mac}</span>
								<span class="value mono">{addZeroDecimal(client.received / 1000000) + "/" + addZeroDecimal(client.sent / 10000) + "MB"}</span>
							</div>
						)}
					</For>
				</Show>
			</div>
			<div class="ap-flex ap-footer">
				<span class="value">{parseAgo(props.ap.last_seen)}</span>
				<button class="card" onClick={() => props.setOption({ ...props.option(), open: true, message: props.ap.hostname + " [" + props.ap.mac + "]" })}>
					attack
				</button>
			</div>
		</div>
	)
}

export default AP
