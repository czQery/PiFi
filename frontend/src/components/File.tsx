import "./File.css"
import type { Component } from "solid-js"

interface fileProps {
	name: string
	link: string
}

const File: Component<fileProps> = props => {
	return (
		<div class="file">
			<span>{props.name}</span>
			<button class="file-button card green" onclick={() => window.open(props.link, "_blank")}>download</button>
		</div>
	)
}

export default File
