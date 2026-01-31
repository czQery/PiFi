import "./Error.css"
import type { Component } from "solid-js"

interface errorProps {
	title: string
	msg: string
}

const Error: Component<errorProps> = props => {
	return (
		<div class="error card">
			<h2>{props.title}</h2>
			<span>{props.msg}</span>
		</div>
	)
}

export default Error
