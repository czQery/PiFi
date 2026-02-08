import { type Component, type JSXElement, Show } from "solid-js"
import type { dbNetData } from "../lib/api/db.ts"
import { getNet, markNet, upNet } from "../lib/api/net.ts"
import type { response } from "../lib/var.ts"

interface netProps {
	name: string
	db: dbNetData
	icon: JSXElement
	option: Function
	setOption: Function
	setLoading: Function
}

const Net: Component<netProps> = props => {
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
			{/*TODO: remove this Show wrapper after implementing other nets*/}
			<Show when={props.name.toLowerCase() === "wigle"}>
				<div class="db-buttons">
					<button
						class="card"
						onClick={() =>
							props.setOption({
								...props.option(),
								open: true,
								buttonFirst: {
									name: "mark",
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
								buttonSecond: { name: "download", action: () => getNet(props.name.toLowerCase()) },
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
