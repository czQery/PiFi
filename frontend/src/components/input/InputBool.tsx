import { Checkbox } from "@ark-ui/solid"
import type { Component } from "solid-js"

interface inputBoolProps {
	name: string
	value: boolean
	callback: (e: any) => void
}

const InputBool: Component<inputBoolProps> = props => {
	return (
		<Checkbox.Root checked={props.value} onCheckedChange={e => props.callback(e)}>
			<Checkbox.Label>{props.name}</Checkbox.Label>
			<Checkbox.Control>
				<div>
					<span></span>
					<div></div>
				</div>
			</Checkbox.Control>
			<Checkbox.HiddenInput />
		</Checkbox.Root>
	)
}

export default InputBool
