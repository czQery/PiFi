import { Field } from "@ark-ui/solid"
import type { Component } from "solid-js"

interface inputBoolProps {
	name: string
	value: string
	placeholder: string
	callback: (e: any) => void
}

const InputString: Component<inputBoolProps> = props => {
	return (
		<Field.Root>
			<Field.Label>{props.name}</Field.Label>
			<Field.Input placeholder={props.placeholder} value={props.value} onInput={e => props.callback(e)} />
		</Field.Root>
	)
}

export default InputString
