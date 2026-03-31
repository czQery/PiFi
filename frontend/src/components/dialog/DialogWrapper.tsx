import { Dialog } from "@ark-ui/solid"
import { type Component, createEffect, createSignal, type JSXElement, Show } from "solid-js"
import { Portal } from "solid-js/web"

export interface dialogWrapperButtonData {
	name: string
	action: Function
}

export interface dialogWrapperData {
	open: boolean
	title: string
	message: string
	helper: any // used to pass something to parent
	buttonFirst: dialogWrapperButtonData | null
	buttonSecond: dialogWrapperButtonData | null
}

interface dialogWrapperProps {
	data: dialogWrapperData
	children?: JSXElement
}

const DialogWrapper: Component<dialogWrapperProps> = props => {
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
		<Dialog.Root className={"card"} open={open()} onFocusOutside={close} onEscapeKeyDown={close} onInteractOutside={close} onPointerDownOutside={close}>
			<Portal>
				<Dialog.Backdrop />
				<Dialog.Positioner>
					<Dialog.Content>
						{/*this is not needed  but the text blinks for time when closing without it*/}
						<Show when={open()}>
							<label style={{ "display": "block", "margin-bottom": "5px", "color": "var(--white)" }}>{props.data.title}</label>
							<Dialog.Description style={{ "width": "300px", "margin-bottom": "10px", "color": "var(--white-hover)" }}>
								{props.data.message}
							</Dialog.Description>
							<Show when={props.children}>
								<div
									style={{
										"display": "flex",
										"flex-direction": "column",
										"gap": "5px",
										"margin-bottom": "10px",
										"max-height": "500px",
										"overflow-y": "auto",
									}}
								>
									{props.children}
								</div>
							</Show>
							<div style={{ "display": "flex", "gap": "10px", "width": "100%", "justify-content": "right" }}>
								<Show when={props.data.buttonFirst}>
									<button class="card" style={{ width: "100px", height: "30px" }} onClick={() => props.data.buttonFirst!.action()}>
										{props.data.buttonFirst!.name}
									</button>
								</Show>
								<Show when={props.data.buttonSecond}>
									<button class="card" style={{ width: "100px", height: "30px" }} onClick={() => props.data.buttonSecond!.action()}>
										{props.data.buttonSecond!.name}
									</button>
								</Show>
							</div>
						</Show>
					</Dialog.Content>
				</Dialog.Positioner>
			</Portal>
		</Dialog.Root>
	)
}

export default DialogWrapper
