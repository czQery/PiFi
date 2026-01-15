import { type Component, type JSXElement, Show } from "solid-js"
import type { dbNetData } from "../lib/db.ts"

interface netProps {
	name: string
	db: dbNetData
	icon: JSXElement
	option: Function
	setOption: Function
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
			{/*TODO: remove this after implementing other nets*/}
			<Show when={props.name.toLowerCase() === "wigle"}>
				<div class="db-buttons">
					<button
						class="card"
						onClick={() =>
							props.setOption({
								...props.option(),
								open: true,
								buttonSecond: { name: "download", action: () => window.open("/api/net/" + props.name.toLowerCase(), "_blank") },
							})}
					>
						manual
					</button>
					<button class="card green">upload</button>
				</div>
			</Show>
		</div>
	)
}

export default Net
