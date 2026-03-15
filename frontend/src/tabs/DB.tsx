import "./DB.css"
import { LucideGlobe, LucideRadio, LucideRadioTower } from "lucide-solid"
import { type Component, createEffect, createSignal } from "solid-js"
import Loading, { type loadingData } from "../components/Loading.tsx"
import Net from "../components/Net.tsx"
import Option, { type optionButtonData, type optionData } from "../components/Option.tsx"
import { type dbData, getDB } from "../lib/api/db.ts"

export const [db, setDB] = createSignal<dbData>({ aps: 0, wigle: { new: 0, net: 0 }, beacondb: { new: 0, net: 0 }, dwpa: { new: 0, net: 0 } } as dbData)

const DB: Component = () => {
	const [loading, setLoading] = createSignal<loadingData>({ title: "Loading", pending: false, msg: "" })
	const [option, setOption] = createSignal<optionData>({
		open: false,
		title: "Manual",
		message: "You can submit the data by your self, and then mark them as uploaded.",
		children: null,
		buttonFirst: { name: "mark" } as optionButtonData,
		buttonSecond: { name: "download" } as optionButtonData,
	})

	createEffect(async () => {
		if (!loading().pending) {
			setDB(await getDB())
		}
	})

	return (
		<div id="db">
			<div id="db-stats" class="card">
				<span class="title">Database</span>
				<div class="db-flex">
					<span>Access points:</span>
					<span class="value mono">{db().aps.toString()}</span>
				</div>
			</div>
			<Option data={option()} />
			<Net name="Wigle" db={db().wigle} icon={<LucideGlobe />} option={option} setOption={setOption} setLoading={setLoading} />
			<Net name="BeaconDB" db={db().beacondb} icon={<LucideRadio />} option={option} setOption={setOption} setLoading={setLoading} />
			<Net name="DWPA" db={db().dwpa} icon={<LucideRadioTower />} option={option} setOption={setOption} setLoading={setLoading} />
			<Loading data={loading} setData={setLoading} />
		</div>
	)
}

export default DB
