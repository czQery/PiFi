import { Dialog, Progress } from "@ark-ui/solid"
import type { Accessor, Component, Setter } from "solid-js"
import { createEffect, createSignal, Show } from "solid-js"
import { Portal } from "solid-js/web"

export interface dialogLoadingData {
	title: string
	pending: boolean
	msg: string
}

interface dialogLoadingProps {
	data: Accessor<dialogLoadingData>
	setData: Setter<dialogLoadingData>
}

const DialogLoading: Component<dialogLoadingProps> = props => {
	const [progress, setProgress] = createSignal(0)

	const fakeProgress = async () => {
		for (let i = 0; i < 80; i++) {
			if (!props.data().pending) return

			if (props.data().msg != "") {
				setProgress(100)
				return
			}

			setProgress(i)
			await sleep(15)
		}
	}

	const sleep = (ms: number) => {
		return new Promise(r => setTimeout(r, ms))
	}

	createEffect(async () => {
		if (props.data().pending) {
			fakeProgress().then()
		}
	})

	return (
		// @ts-ignore
		<Dialog.Root className={"card"} open={props.data().pending} closeOnEscape={false} closeOnInteractOutside={false}>
			<Portal>
				<Dialog.Backdrop />
				<Dialog.Positioner>
					<Dialog.Content>
						<Show when={props.data().msg == ""}>
							<Progress.Root value={progress()}>
								<Progress.Label>{props.data().title}</Progress.Label>
								<Progress.Track>
									<Progress.Range />
								</Progress.Track>
							</Progress.Root>
						</Show>
						<Show when={props.data().msg != ""}>
							<label style={{ "display": "block", "margin-bottom": "5px", "color": "var(--white)" }}>Error</label>
							<Dialog.Description style={{ width: "300px", color: "var(--pink)" }}>{props.data().msg}</Dialog.Description>
							<Dialog.CloseTrigger
								onClick={() => {
									props.setData({ ...props.data(), pending: false })
								}}
							>
								close
							</Dialog.CloseTrigger>
						</Show>
					</Dialog.Content>
				</Dialog.Positioner>
			</Portal>
		</Dialog.Root>
	)
}

export default DialogLoading
