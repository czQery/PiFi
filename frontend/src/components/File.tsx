import "./File.css"
import { type Component, type JSXElement, Show } from "solid-js"

interface fileProps {
	name: string
	detail?: string
	icon?: JSXElement
	link: string
}

const File: Component<fileProps> = props => {
	return (
		<div class="file">
			<Show when={props.icon} children={props.icon} />
			<span>{props.name}</span>
			<Show when={props.detail}>
				<span style={{ color: "var(--white-hover)" }}>{props.detail}</span>
			</Show>
			<button class="file-button card" onclick={() => window.open(props.link, "_blank")}>download</button>
		</div>
	)
}

export default File
