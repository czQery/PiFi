import { type Accessor, type Component, type JSXElement, type Setter, Show } from "solid-js"
import type { dbNetData } from "../lib/api/db.ts"
import { markNet, upNet } from "../lib/api/net.ts"
import { api, type response } from "../lib/var.ts"
import type { dialogLoadingData } from "./dialog/DialogLoading.tsx"
import type { dialogWrapperButtonData, dialogWrapperData } from "./dialog/DialogWrapper.tsx"

interface netProps {
	name: string
	link: string
	db: dbNetData
	icon: JSXElement
	option: Accessor<dialogWrapperData>
	setOption: Setter<dialogWrapperData>
	setLoading: Setter<dialogLoadingData>
}

const Net: Component<netProps> = props => {
	return (
		<div id={"db-net-" + props.name.toLowerCase()} class="card">
			<div class="db-flex" style="grid-column: 1 / 3;">
				{props.icon}
				<span class="title">{props.name}</span>
			</div>
			<a class="link mono" href={"https://" + props.link}>{props.link}</a>
			<div class="db-flex">
				<span>New:</span>
				<span class="value mono">{props.db.new.toString()}</span>
			</div>
			<div class="db-flex">
				<span>Net:</span>
				<span class="value mono">{props.db.net.toString()}</span>
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
								helper: props.name.toLowerCase(),
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
								buttonSecond: ((): dialogWrapperButtonData | null => {
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
