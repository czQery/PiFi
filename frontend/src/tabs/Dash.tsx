import type { Component } from "solid-js"
import { createSignal, Index, onMount, Show } from "solid-js"

import "./Dash.css"
import { useNavigate } from "@solidjs/router"
import { LucideCpu, LucideLocateFixed, LucideLocateOff, LucideMemoryStick, LucideSettings, LucideWifi, LucideWifiOff } from "lucide-solid"
import type { logData } from "../lib/log.ts"
import { getLog } from "../lib/log.ts"
import { addZero } from "../lib/other.ts"
import type { statsData } from "../lib/stats.ts"
import { getStats } from "../lib/stats.ts"

export const [stats, setStats] = createSignal<statsData>(
	{ cpu: 0, mem_total: 0, mem_used: 0, hotspot: { ssid: "" }, gps: { lat: 0, lon: 0, alt: 0, mode: 0, time: 0 } } as statsData,
)
export const [log, setLog] = createSignal<logData[]>([])

const Dash: Component = () => {
	const navigate = useNavigate()

	onMount(async () => {
		setStats(await getStats())
		setLog(await getLog())
	})

	return (
		<div id="dash">
			<div id="dash-overview" class="card" data-disabled={stats().hotspot.ssid !== "" ? "false" : "true"}>
				<Show when={stats().hotspot.ssid !== ""} fallback={<LucideWifiOff id="dash-overview-icon" />}>
					<LucideWifi id="dash-overview-icon" />
				</Show>
				<h2>{stats().hotspot.ssid !== "" ? stats().hotspot.ssid : "xxxx"}</h2>
				<h4>Status: {stats().hotspot.ssid !== "" ? stats().hotspot.portal ? "portal" : "normal" : "off"}</h4>
				<span>Clients: ?, Passwords: ?</span>
				<button id="dash-overview-settings" onClick={() => navigate("/settings")}>
					<LucideSettings />
					<span>settings</span>
				</button>
			</div>
			<div id="dash-gps" class="card">
				<div>
					<Show when={stats().gps.mode >= 2} fallback={<LucideLocateOff />}>
						<LucideLocateFixed />
					</Show>
					<span>{"Coordinates: " + stats().gps.lat.toString() + ", " + stats().gps.lon.toString()}</span>
				</div>
				<span>{stats().gps.mode >= 2 ? stats().gps.alt.toString() + "m" : "no location"}</span>
			</div>
			<div id="dash-stats">
				<div class="card">
					<LucideCpu />
					<span>{stats().cpu + "%"}</span>
				</div>
				<div class="card">
					<LucideMemoryStick />
					<span>{(stats().mem_used / 1000000000).toFixed(2) + "/" + (stats().mem_total / 1000000000).toFixed(2) + "GB"}</span>
				</div>
			</div>
			<div id="dash-log" class="card">
				<Index each={log() as logData[]}>
					{(log, _) => (
						<li>
							<span
								style={{
									color: ((): string => {
										switch (log().level.slice(0, 4)) {
											case "erro":
												return "var(--pink)"
											case "warn":
												return "var(--orange)"
											case "debu":
												return "var(--white)"
											default:
												return "var(--blue)"
										}
									})(),
								}}
							>
								{log().level.slice(0, 4).toUpperCase()}
							</span>
							<span style={{ opacity: 0.6 }}>
								{"[" + addZero(log().time.getHours()) + ":" + addZero(log().time.getMinutes()) + ":" + addZero(log().time.getSeconds()) + "]"}
							</span>
							<span>{log().msg}</span>
							<Show when={log().err as string}>
								<span style={{ "color": "var(--pink)", "margin-left": "auto" }}>{"err=" + log().err}</span>
							</Show>
							<Show when={log().data as string}>
								<span style={{ "color": "var(--green)", "margin-left": "auto" }}>{"data=" + log().data}</span>
							</Show>
							<Show when={log().ssid as string}>
								<span style={{ "color": "var(--green)", "margin-left": "auto" }}>{"ssid=" + log().ssid}</span>
							</Show>
							<Show when={log().iface as string}>
								<span style={{ "color": "var(--blue)", "margin-left": "auto" }}>{"iface=" + log().iface}</span>
							</Show>
						</li>
					)}
				</Index>
			</div>
		</div>
	)
}

export default Dash
