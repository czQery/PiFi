import type {Component} from "solid-js";
import {createSignal, Index, onMount, Show} from "solid-js";

import "./Dash.css"
import {getLog, getStats, logData, statsData} from "../lib/api";
import {LucideCpu, LucideMemoryStick, LucideSettings, LucideWifi} from "lucide-solid";
import {useNavigate} from "@solidjs/router";

export const [stats, setStats] = createSignal<statsData>({cpu: 0, mem_total: 0, mem_used: 0});
export const [log, setLog] = createSignal<logData[]>([]);

const Dash: Component = () => {

    const navigate = useNavigate()

    onMount(async () => {
        setStats(await getStats())
        setLog(await getLog())
    })

    return (
        <div id="dash">
            <div id="dash-overview" class="card">
                <LucideWifi id="dash-overview-icon"/>
                <h2>PizzaQY</h2>
                <h4>Mode: hotspot</h4>
                <span>Clients: 2, Passwords: 15</span>
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
                        <span style={{opacity: 0.6}}>{"[" + log().time.getHours() + ":" + log().time.getMinutes() + ":" + log().time.getSeconds() + "]"}</span>
                        <span>{log().msg}</span>
                        <Show when={log().err<string>}>
                            <span style={{color: "var(--pink)", "margin-left": "auto"}}>{"err=" + log().err}</span>
                        </Show>
                    </li>
                }</Index>
            </div>
        </div>
    )
}

export default Dash