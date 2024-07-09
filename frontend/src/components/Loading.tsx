import {Component, createEffect, createSignal, Setter} from "solid-js";
import {Portal} from "solid-js/web";
import {Dialog, Progress} from "@ark-ui/solid";

export interface loadingData {
    title: string
    pending: boolean
}

interface loadingProps {
    data: loadingData
}

const Loading: Component<loadingProps> = (props) => {

    const [pending, setPending] = createSignal(false)
    const [name, setName] = createSignal("loading")
    const [progress, setProgress] = createSignal(0)

    const fakePending = async (pending: () => boolean, progress: Setter<number>) => {
        for (let i = 0; i < 80; i++) {
            if (!pending()) return

            progress(i)
            await sleep(15)
        }
    }

    const sleep = (ms: number) => {
        return new Promise(r => setTimeout(r, ms))
    }

    createEffect(async () => {
        if (props.data.pending) {
            setName(props.data.title)
            setPending(true)
            fakePending(pending, setProgress).then()
        } else {
            await sleep(200)
            setPending(false)
        }
    })

    return (
        <Dialog.Root className={"card"} open={pending()} onOpenChange={(e) => setPending(e.open)}>
            <Portal>
                <Dialog.Backdrop/>
                <Dialog.Positioner>
                    <Dialog.Content>
                        <Progress.Root value={progress()}>
                            <Progress.Label>{name()}</Progress.Label>
                            <Progress.Track>
                                <Progress.Range/>
                            </Progress.Track>
                        </Progress.Root>
                    </Dialog.Content>
                </Dialog.Positioner>
            </Portal>
        </Dialog.Root>
    )
}

export default Loading