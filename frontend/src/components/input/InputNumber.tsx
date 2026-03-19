import { NumberInput } from "@ark-ui/solid"
import { type Component } from "solid-js"

interface inputBoolProps {
	name: string
	value: number
	min: number
	max: number
	callback: (e: any) => void
}

const InputNumber: Component<inputBoolProps> = props => {
	return (
		<NumberInput.Root value={(props.value || props.min).toString()} min={props.min} max={props.max} onValueChange={e => props.callback(e)}>
			<NumberInput.Label>Channel</NumberInput.Label>
			<NumberInput.Input class="mono" />
			<NumberInput.Control>
				<NumberInput.DecrementTrigger>-</NumberInput.DecrementTrigger>
				<NumberInput.IncrementTrigger>+</NumberInput.IncrementTrigger>
			</NumberInput.Control>
		</NumberInput.Root>
	)
}

export default InputNumber
