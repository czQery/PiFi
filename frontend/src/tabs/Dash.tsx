import type {Component} from "solid-js"
import {createSignal, Index, onMount, Show} from "solid-js"

import "./Dash.css"
import {LucideCpu, LucideMemoryStick, LucideSettings, LucideWifi, LucideWifiOff} from "lucide-solid"
import {useNavigate} from "@solidjs/router"
import {getStats, statsData} from "../lib/stats"
import {getLog, logData} from "../lib/log"
import {addZero} from "../lib/other";

export const [stats, setStats] = createSignal<statsData>({cpu: 0, mem_total: 0, mem_used: 0, hotspot: {ssid: ""}})
export const [log, setLog] = createSignal<logData[]>([])

const Dash: Component = () => {

    const navigate = useNavigate()

    onMount(async () => {
        setStats(await getStats())
        setLog(await getLog())
    })

    return (
        <div id="dash">
            <div id="dash-overview" class="card" data-disabled={(stats().hotspot.ssid !== "" ? "false" : "true")}>
                <Show when={stats().hotspot.ssid !== ""<boolean>} fallback={<LucideWifiOff id="dash-overview-icon"/>}>
                    <LucideWifi id="dash-overview-icon"/>
                </Show>
                <h2>{(stats().hotspot.ssid !== "" ? stats().hotspot.ssid : "xxxx")}</h2>
                <h4>Status: {(stats().hotspot.ssid !== "" ? (stats().hotspot.portal ? "portal" : "normal") : "off")}</h4>
                <span>Clients: ?, Passwords: ?</span>
                <button id="dash-overview-settings" onClick={() => navigate("/settings")}>
                    <LucideSettings/>
                    <span>settings</span>
                </button>
            </div>
            <div id="dash-stats">
                <div class="card">
                    <LucideCpu/>
                    <span>{stats().cpu + "%"}</span>
                </div>
                <div class="card">
                    <LucideMemoryStick/>
                    <span>{stats().mem_used + "/" + stats().mem_total + "GB"}</span>
                </div>
            </div>
            <div id="dash-log" class="card">
                <Index each={log()<logData[]>}>{(log, i) =>
                    <li>
                        <span style={{
                            color: ((): string => {
                                switch (log().level.slice(0, 4)) {
                                    case "erro":
                                        return "var(--pink)"
                                    case "warn":
                                        return "var(--orange)"
                                    default:
                                        return "var(--blue)"
                                }
                            })()
                        }}>{log().level.slice(0, 4).toUpperCase()}</span>
                        <span style={{opacity: 0.6}}>{"[" + addZero(log().time.getHours()) + ":" + addZero(log().time.getMinutes()) + ":" + addZero(log().time.getSeconds()) + "]"}</span>
                        <span>{log().msg}</span>
                        <Show when={log().err<string>}>
                            <span style={{color: "var(--pink)", "margin-left": "auto"}}>{"err=" + log().err}</span>
                        </Show>
                        <Show when={log().data<string>}>
                            <span style={{color: "var(--green)", "margin-left": "auto"}}>{"data=" + log().data}</span>
                        </Show>
                        <Show when={log().ssid<string>}>
                            <span style={{color: "var(--green)", "margin-left": "auto"}}>{"ssid=" + log().ssid}</span>
                        </Show>
                    </li>
                }</Index>
            </div>
        </div>
    )
}

export default Dash