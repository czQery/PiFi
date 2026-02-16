import {type Component, createSignal, For, type JSXElement, onMount, Show} from "solid-js"
import type {dbNetData} from "../lib/api/db.ts"
import {getNetDWPA, markNet, upNet} from "../lib/api/net.ts"
import {api, type response} from "../lib/var.ts"
import File from "./File.tsx"
import type {optionButtonData} from "./Option.tsx"

interface netProps {
	name: string
	db: dbNetData
	icon: JSXElement
	option: Function
	setOption: Function
	setLoading: Function
}

const Net: Component<netProps> = props => {
	const [dataDWPA, setDataDWPA] = createSignal<string[]>([])

	onMount(async () => {
		setDataDWPA(await getNetDWPA())
	})

	return (
		<div id="db-net-wigle" class="card">
			<div class="db-flex">
				{props.icon}
				<span class="title">{props.name}</span>
			</div>
			<div class="db-flex">
				<span>New:</span>
				<span class="value">{props.db.new.toString()}</span>
			</div>
			<div class="db-flex">
				<span>Net:</span>
				<span class="value">{props.db.net.toString()}</span>
			</div>
			{/*TODO: remove this Show wrapper after implementing all nets*/}
			<Show when={props.name.toLowerCase() !== "beacondb"}>
				<div class="db-buttons">
					<button
						class="card"
						onClick={() =>
							props.setOption({
								...props.option(),
								open: true,
								children: ((): Element | null => {
									switch (props.name.toLowerCase()) {
										case "dwpa":
											return <For each={dataDWPA()}>{(file, _) => <File name={file} link={api + "api/cap/" + file} />}</For> as Element
										default:
											return null
									}
								})(),
								buttonFirst: {
									name: "mark all",
									action: async () => {
										props.setOption({ ...props.option(), open: false })
										props.setLoading({ title: "Marking", pending: true, msg: "" })
										const rsp: response = await markNet(props.name.toLowerCase(), true)

										if (rsp.message != "") {
											props.setLoading({ title: "Marking", pending: true, msg: rsp.message })
											return
										}

										props.setLoading({ title: "Marking", pending: false, msg: "" })
									},
								},
								buttonSecond: ((): optionButtonData | null => {
									switch (props.name.toLowerCase()) {
										case "wigle":
											return { name: "download", action: () => window.open(api + "api/net/wigle", "_blank") }
										default:
											return null
									}
								})(),
							})}
					>
						manual
					</button>
					<button
						class="card green"
						onClick={async () => {
							props.setLoading({ title: "Uploading", pending: true, msg: "" })
							const rsp: response = await upNet(props.name.toLowerCase())

							if (rsp.message != "") {
								props.setLoading({ title: "Uploading", pending: true, msg: rsp.message })
								return
							}

							props.setLoading({ title: "Uploading", pending: false, msg: "" })
						}}
					>
						upload
					</button>
				</div>
			</Show>
		</div>
	)
}

export default Net
