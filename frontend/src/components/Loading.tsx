import {Component, createEffect, createSignal, Setter, Show} from "solid-js"
import {Portal} from "solid-js/web"
import {Dialog, Progress} from "@ark-ui/solid"

export interface loadingData {
    title: string
    pending: boolean
    msg: string
}

interface loadingProps {
    data: loadingData
}

const Loading: Component<loadingProps> = (props) => {

    const [pending, setPending] = createSignal(false)
    const [name, setName] = createSignal("loading")
    const [message, setMessage] = createSignal("")
    const [progress, setProgress] = createSignal(0)

    const fakePending = async (pending: () => boolean, progress: Setter<number>) => {
        for (let i = 0; i < 80; i++) {
            if (!pending()) return

            if (message() != "") {
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
        if (props.data.pending) {
            setName(props.data.title)
            setMessage(props.data.msg)
            setPending(true)
            fakePending(pending, setProgress).then()
        } else {
            await sleep(200)
            setPending(false)
        }
    })

    return (
        <Dialog.Root className={"card"} open={pending()} closeOnEscape={false} closeOnInteractOutside={false}>
            <Portal>
                <Dialog.Backdrop/>
                <Dialog.Positioner>
                    <Dialog.Content>
                        <Show when={message() == ""<boolean>}>
                            <Progress.Root value={progress()}>
                                <Progress.Label>{name()}</Progress.Label>
                                <Progress.Track>
                                    <Progress.Range/>
                                </Progress.Track>
                            </Progress.Root>
                        </Show>
                        <Show when={message() != ""<boolean>}>
                            <label style={{display: "block", "margin-bottom": "5px", color: "var(--white)"}}>Error</label>
                            <Dialog.Description style={{width: "300px", color: "var(--pink)"}}>{message()}</Dialog.Description>
                            <Dialog.CloseTrigger onClick={() => {
                                props.data.pending = false
                                setPending(false)
                            }}>close</Dialog.CloseTrigger>
                        </Show>
                    </Dialog.Content>
                </Dialog.Positioner>
            </Portal>
        </Dialog.Root>
    )
}

export default Loading