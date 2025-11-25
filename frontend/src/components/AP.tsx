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
			<div class="ap-title">
				<LucideWifi />
				<span>{props.ap.hostname}</span>
			</div>
			<span>BSSID: {props.ap.mac}</span>
			<span>CHANNEL: {props.ap.channel}</span>
		</div>
	)
}

export default AP
