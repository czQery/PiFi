import "./DB.css"
import {LucideGlobe, LucideRadio, LucideRadioTower} from "lucide-solid"
import {type Component, createSignal, onMount} from "solid-js"
import Net from "../components/Net.tsx"
import {type dbData, getDB} from "../lib/db.ts"

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
			<Net name="Wigle" db={db().wigle} icon={<LucideGlobe />} />
			<Net name="BeaconDB" db={db().beacondb} icon={<LucideRadio />} />
			<Net name="DWPA" db={db().dwpa} icon={<LucideRadioTower />} />
		</div>
	)
}

export default DB
