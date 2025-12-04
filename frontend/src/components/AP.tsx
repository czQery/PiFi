import "./AP.css"
import { LucideWifi } from "lucide-solid"
import type { Component } from "solid-js"
import type { bettercapWifiData } from "../lib/bettercap.ts"

interface APProps {
	ap: bettercapWifiData
}

const AP: Component<APProps> = props => {
	return (
		<div class="ap card">
			<div class="ap-flex" style="justify-content: space-between;">
				<div class="ap-flex">
					<LucideWifi />
					<span class="ap-title">{props.ap.hostname}</span>
				</div>
				<div class="ap-flex">
					<span>[RSSI:</span>
					<span class="ap-value">{props.ap.rssi}</span>
					<span>]</span>
					<span>[CH:</span>
					<span class="ap-value">{props.ap.channel}</span>
					<span>]</span>
				</div>
			</div>
			<div class="ap-flex" style="justify-content: space-between;">
				<div class="ap-flex">
					<span>BSSID:</span>
					<span class="ap-value">{props.ap.mac}</span>
				</div>
				<div class="ap-flex">
					<span>[</span>
					<span style={{ color: (props.ap.encryption === "OPEN" ? "var(--green)" : "var(--pink)") }}>{props.ap.encryption}</span>
					<span>]</span>
				</div>
			</div>
			<div class="ap-flex">
				<span>VENDOR:</span>
				<span class="ap-value">{props.ap.vendor !== "" ? props.ap.vendor : "unknown"}</span>
			</div>
		</div>
	)
}

export default AP
