import "./DB.css"
import { LucideGlobe, LucideRadio, LucideRadioTower } from "lucide-solid"
import { type Component, createSignal, onMount } from "solid-js"
import { type dbData, getDB } from "../lib/db.ts"

export const [db, setDB] = createSignal<dbData>({ aps: 0, wigle: { new: 0, net: 0 }, beacondb: { new: 0, net: 0 }, dwpa: { new: 0, net: 0 } } as dbData)

const DB: Component = () => {
	onMount(async () => {
		setDB(await getDB())
	})

	return (
		<div id="db">
			<div id="db-stats" class="card">
				<span class="title">Database</span>
				<div class="db-flex">
					<span>Access points:</span>
					<span class="value">{db().aps.toString()}</span>
				</div>
			</div>
			<div id="db-net-wigle" class="card">
				<div class="db-flex">
					<LucideGlobe />
					<span class="title">Wigle</span>
				</div>
				<div class="db-flex">
					<span>New:</span>
					<span class="value">{db().wigle.new.toString()}</span>
				</div>
				<div class="db-flex">
					<span>Net:</span>
					<span class="value">{db().wigle.net.toString()}</span>
				</div>
				<button class="card green">upload</button>
			</div>
			<div id="db-net-beacondb" class="card">
				<div class="db-flex">
					<LucideRadio />
					<span class="title">BeaconDB</span>
				</div>
				<div class="db-flex">
					<span>New:</span>
					<span class="value">{db().beacondb.new.toString()}</span>
				</div>
				<div class="db-flex">
					<span>Net:</span>
					<span class="value">{db().beacondb.net.toString()}</span>
				</div>
				<button class="card green">upload</button>
			</div>
			<div id="db-net-dwpa" class="card">
				<div class="db-flex">
					<LucideRadioTower />
					<span class="title">DWPA</span>
				</div>
				<div class="db-flex">
					<span>New:</span>
					<span class="value">{db().dwpa.new.toString()}</span>
				</div>
				<div class="db-flex">
					<span>Net:</span>
					<span class="value">{db().dwpa.net.toString()}</span>
				</div>
				<button class="card green">upload</button>
			</div>
		</div>
	)
}

export default DB
