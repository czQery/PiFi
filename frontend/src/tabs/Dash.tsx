import type {Component} from "solid-js";
import {createSignal, onMount} from "solid-js";

import "./Dash.css"
import {getStats, statsData} from "../lib/api";
import {LucideCpu, LucideMemoryStick, LucideSettings, LucideWifi} from "lucide-solid";
import {useNavigate} from "@solidjs/router";

export const [stats, setStats] = createSignal<statsData>();

const Dash: Component = () => {

    const navigate = useNavigate()

    onMount(async () => {
        setStats(await getStats())
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
                    <span>{stats()?.cpu + "%"}</span>
                </div>
                <div class="card">
                    <LucideMemoryStick/>
                    <span>{stats()?.mem_used + "/" + stats()?.mem_total + "GB"}</span>
                </div>
            </div>
            <div id="dash-log" class="card">
                <span>INFO[02/07/2024 - 11:03:06] config - successfully loaded</span>
                <span>WARN[02/07/2024 - 11:03:06] dist - load failed error="open ./dist: no such file or directory"</span>
                <span>INFO[02/07/2024 - 11:03:06] main - nmcli loaded</span>
                <span>INFO[02/07/2024 - 11:03:06] fiber - started</span>
            </div>
        </div>
    )
}

export default Dash