import type {Component} from "solid-js"

import "./App.css"

import {Tabs} from "@ark-ui/solid"
import {useLocation, useNavigate} from "@solidjs/router"

const App: Component = (props) => {
    const { pathname } = useLocation()
    const navigate = useNavigate()

    navigate("/dash")

    return (
        <>
            <header>
                <div class="container">
                    <h1>PiFi</h1>
                    <Tabs.Root value={pathname.substring(1)} onValueChange={({value}) => {
                        navigate("/" + value, {replace: true})
                    }}>
                        <Tabs.List>
                            <Tabs.Trigger value="dash">dash</Tabs.Trigger>
                            <Tabs.Trigger value="scan">scan</Tabs.Trigger>
                            <Tabs.Trigger value="portal">portal</Tabs.Trigger>
                            <Tabs.Indicator/>
                        </Tabs.List>
                    </Tabs.Root>
                </div>
            </header>
            <main class="container">
                {
                    //@ts-ignore
                    props.children
                }
            </main>
        </>
    )
}

export default App
