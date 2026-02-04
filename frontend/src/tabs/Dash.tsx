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
				<ul>
					<Index each={log() as logData[]}>
						{(entry, _) => (
							<li>
								<span
									style={{
										color: ((): string => {
											switch (entry().level.slice(0, 4)) {
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
									{entry().level.slice(0, 4).toUpperCase()}
								</span>
								<span style={{ opacity: 0.6 }}>
									{"[" + addZero(entry().time.getHours()) + ":" + addZero(entry().time.getMinutes()) + ":"
										+ addZero(entry().time.getSeconds())
										+ "]"}
								</span>
								<span style={{ "margin-right": "10px" }}>{entry().msg}</span>
								<Index each={entry().items}>
									{(item, i) => (
										<span
											style={{
												"color": item().color,
												"margin-left": (i === 0 ? "auto" : ""),
												"margin-right": (i !== entry().items.length - 1 ? "10px" : ""),
											}}
										>
											{item().name + "=" + item().value}
										</span>
									)}
								</Index>
							</li>
						)}
					</Index>
				</ul>
			</div>
		</div>
	)
}

export default Dash
