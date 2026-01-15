import { Dialog } from "@ark-ui/solid"
import { type Component, createEffect, createSignal } from "solid-js"
import { Portal } from "solid-js/web"

export interface optionButtonData {
	name: string
	action: Function
}

export interface optionData {
	open: boolean
	title: string
	message: string
	buttonFirst: optionButtonData
	buttonSecond: optionButtonData
}

interface optionProps {
	data: optionData
}

const Option: Component<optionProps> = props => {
	const [open, setOpen] = createSignal(false)

	const close = () => {
		props.data.open = false
		setOpen(false)
	}

	createEffect(async () => {
		if (props.data.open) {
			setOpen(true)
		} else {
			setOpen(false)
		}
	})

	return (
		// @ts-ignore
		<Dialog.Root
			className={"card"}
			oepn
			open={open()}
			onFocusOutside={close}
			onEscapeKeyDown={close}
			onInteractOutside={close}
			onPointerDownOutside={close}
		>
			<Portal>
				<Dialog.Backdrop />
				<Dialog.Positioner>
					<Dialog.Content>
						<label style={{ "display": "block", "margin-bottom": "5px", "color": "var(--white)" }}>{props.data.title}</label>
						<Dialog.Description style={{ "width": "300px", "margin-bottom": "10px", "color": "var(--white-hover)" }}>
							{props.data.message}
						</Dialog.Description>
						<div style={{ "display": "flex", "gap": "10px", "width": "100%", "justify-content": "right" }}>
							<button class="card" style={{ width: "100px", height: "30px" }} onClick={() => props.data.buttonFirst.action()}>
								{props.data.buttonFirst.name}
							</button>
							<button class="card" style={{ width: "100px", height: "30px" }} onClick={() => props.data.buttonSecond.action()}>
								{props.data.buttonSecond.name}
							</button>
						</div>
					</Dialog.Content>
				</Dialog.Positioner>
			</Portal>
		</Dialog.Root>
	)
}

export default Option
