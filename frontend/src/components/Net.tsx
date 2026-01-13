import type {Component, JSXElement} from "solid-js"
import type {dbNetData} from "../lib/db.ts"

interface NetProps {
	name: string
	db: dbNetData
	icon: JSXElement
}

const Net: Component<NetProps> = props => {
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
			<div class="db-buttons">
				<button class="card" onClick={() => window.open("/api/net/" + props.name.toLowerCase(), "_blank")}>raw</button>
				<button class="card green">upload</button>
			</div>
		</div>
	)
}

export default Net
