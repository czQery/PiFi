import "./DB.css"
import { LucideFile, LucideFileCheck, LucideGlobe, LucideRadio, LucideRadioTower } from "lucide-solid"
import { type Component, createEffect, createSignal, For, Match, Switch } from "solid-js"
import DialogLoading, { type dialogLoadingData } from "../components/dialog/DialogLoading.tsx"
import DialogWrapper, { type dialogWrapperButtonData, type dialogWrapperData } from "../components/dialog/DialogWrapper.tsx"
import File from "../components/File.tsx"
import Net from "../components/Net.tsx"
import { type dbData, getDB } from "../lib/api/db.ts"
import { getNetDWPA, type netDWPAData } from "../lib/api/net.ts"
import { api } from "../lib/var.ts"

export const [db, setDB] = createSignal<dbData>({ aps: 0, wigle: { new: 0, net: 0 }, beacondb: { new: 0, net: 0 }, dwpa: { new: 0, net: 0 } } as dbData)
export const [dataDWPA, setDataDWPA] = createSignal<netDWPAData[]>([] as netDWPAData[])

const DB: Component = () => {
	const [loading, setLoading] = createSignal<dialogLoadingData>({ title: "Loading", pending: false, msg: "" })
	const [option, setOption] = createSignal<dialogWrapperData>({
		open: false,
		title: "Manual",
		message: "You can submit the data by your self, and then mark them as uploaded.",
		helper: null,
		buttonFirst: { name: "mark" } as dialogWrapperButtonData,
		buttonSecond: { name: "download" } as dialogWrapperButtonData,
	})

	createEffect(async () => {
		if (!loading().pending) {
			setDB(await getDB())
			setDataDWPA(await getNetDWPA())
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
			<DialogWrapper data={option()}>
				<Switch>
					<Match when={option().helper === "dwpa"}>
						<For each={dataDWPA() as netDWPAData[]}>
							{(ap, _) => (
								<File
									icon={ap.net ? <LucideFileCheck /> : <LucideFile />}
									name={ap.ssid}
									detail={"[" + ap.bssid + "]"}
									link={api + "api/cap/" + ap.file}
								/>
							)}
						</For>
					</Match>
				</Switch>
			</DialogWrapper>
			<Net name="Wigle" link="wigle.net" db={db().wigle} icon={<LucideGlobe />} option={option} setOption={setOption} setLoading={setLoading} />
			<Net name="BeaconDB" link="beacondb.net" db={db().beacondb} icon={<LucideRadio />} option={option} setOption={setOption} setLoading={setLoading} />
			<Net
				name="DWPA"
				link="wpa-sec.stanev.org"
				db={db().dwpa}
				icon={<LucideRadioTower />}
				option={option}
				setOption={setOption}
				setLoading={setLoading}
			/>
			<DialogLoading data={loading} setData={setLoading} />
		</div>
	)
}

export default DB
