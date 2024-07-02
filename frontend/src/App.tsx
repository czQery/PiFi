import {Component, createEffect, createSignal, onMount, Show} from "solid-js";

import "./App.css"

import {Menu, Tabs} from "@ark-ui/solid"
import {useLocation, useNavigate} from "@solidjs/router"
import {LucideFileCog, LucideLayoutDashboard, LucideLogOut, LucideMenu, LucideRadar, LucideSettings, LucideWifi} from "lucide-solid";
import Auth from "./components/Auth";
import {auth, authSave} from "./lib/auth";

export const [logged, setLogged] = createSignal<undefined | boolean>();

const App: Component = (props: { children }) => {
    const {pathname} = useLocation()
    const navigate = useNavigate()

    onMount(async () => {
        if (pathname === "/" || pathname === "") {
            navigate("/dash")
        }

        setLogged(await auth())
    })

    createEffect(() => {
        if (logged() !== undefined) {
            document.getElementById("root")!.style.display = ""
            document.getElementById("loading")!.style.display = "none"
        }
    })

    return (
        <Show when={logged()<undefined | boolean>} fallback={<Auth/>} keyed>
            <header>
                <h1 id="pifi">PiFi</h1>
                <Menu.Root id="header-menu" positioning={{"gutter": 10, "placement": "bottom-end"}}>
                    <Menu.Trigger><LucideMenu/></Menu.Trigger>
                    <Menu.Positioner>
                        <Menu.Content>
                            <Tabs.Root orientation="vertical" value={(pathname.substring(1) !== "" ? pathname.substring(1) : "dash")} onValueChange={({value}) => navigate("/" + value)}>
                                <Tabs.List>
                                    <div id="header-menu-list">
                                        <Tabs.Trigger value="dash"><LucideLayoutDashboard/></Tabs.Trigger>
                                        <Tabs.Trigger value="settings"><LucideSettings/></Tabs.Trigger>
                                        <Tabs.Trigger value="scan"><LucideRadar/></Tabs.Trigger>
                                        <Tabs.Trigger value="portal"><LucideFileCog/></Tabs.Trigger>
                                        <Tabs.Indicator/>
                                    </div>
                                    <button id="header-menu-logout" onClick={async () => {
                                        setLogged(false)
                                        authSave("").then()
                                    }}><LucideLogOut/></button>
                                </Tabs.List>
                            </Tabs.Root>
                        </Menu.Content>
                    </Menu.Positioner>
                </Menu.Root>
            </header>
            <main class="container">
                {props.children}
            </main>
        </Show>
    )
}

export default App
