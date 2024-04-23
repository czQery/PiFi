import {Component, createEffect, createSignal, onMount, Show} from "solid-js";

import "./App.css"

import {Menu, Tabs} from "@ark-ui/solid"
import {useLocation, useNavigate} from "@solidjs/router"
import {LucideFileCog, LucideLayoutDashboard, LucideLogOut, LucideMenu, LucideWifi} from "lucide-solid";
import Auth from "./components/Auth";
import {auth, authSave} from "./hp/auth";

export const [logged, setLogged] = createSignal<undefined | boolean>();

const App: Component = (props: { children }) => {
    const {pathname} = useLocation()
    const navigate = useNavigate()

    navigate("/dash")

    onMount(async () => {
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
                            <Tabs.Root orientation="vertical" value={pathname.substring(1)} onValueChange={({value}) => navigate("/" + value, {replace: true})}>
                                <Tabs.List>
                                    <div id="header-menu-list">
                                        <Tabs.Trigger value="dash"><LucideLayoutDashboard/></Tabs.Trigger>
                                        <Tabs.Trigger value="scan"><LucideWifi/></Tabs.Trigger>
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
