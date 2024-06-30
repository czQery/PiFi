import type {Component} from "solid-js";

import "./Dash.css"
import {createSignal, onMount} from "solid-js";
import {getStats, statsData} from "../lib/api";
import {LucideCpu, LucideMemoryStick} from "lucide-solid";

export const [stats, setStats] = createSignal<statsData>();

const Dash: Component = () => {

    onMount(async () => {
        setStats(await getStats())
    })

    return (
        <div id="dash" class="card">
            <div id="dash-stats">
                <h2>> Stats</h2>
                <div>
                    <LucideCpu/>
                    <span>{stats()?.cpu * 100 + "%"}</span>
                </div>
                <div>
                    <LucideMemoryStick/>
                    <span>{stats()?.mem_used + "/" + stats()?.mem_total}</span>
                </div>
            </div>
            <div id="dash-log">
                <h2>> Log</h2>
            </div>
        </div>
    )
}

export default Dash