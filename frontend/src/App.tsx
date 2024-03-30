import type {Component} from "solid-js"

import "./App.css"

import {Menu, Tabs} from "@ark-ui/solid"
import {useLocation, useNavigate} from "@solidjs/router"
import {LucideFileCog, LucideLayoutDashboard, LucideLogOut, LucideMenu, LucideWifi} from "lucide-solid";

const App: Component = (props) => {
    const {pathname} = useLocation()
    const navigate = useNavigate()

    navigate("/dash")

    return (
        <>
            <header>
                <h1>PiFi</h1>
                <Menu.Root id="header-menu" positioning={{"gutter": 10, "placement": "bottom-end"}}>
                    <Menu.Trigger><LucideMenu/></Menu.Trigger>
                    <Menu.Positioner>
                        <Menu.Content>
                            <Tabs.Root orientation="vertical" value={pathname.substring(1)}
                                       onValueChange={({value}) => {
                                           navigate("/" + value, {replace: true})
                                       }}>
                                <Tabs.List>
                                    <div id="header-menu-list">
                                        <Tabs.Trigger value="dash"><LucideLayoutDashboard/></Tabs.Trigger>
                                        <Tabs.Trigger value="scan"><LucideWifi/></Tabs.Trigger>
                                        <Tabs.Trigger value="portal"><LucideFileCog/></Tabs.Trigger>
                                        <Tabs.Indicator/>
                                    </div>
                                    <button id="header-menu-logout"><LucideLogOut/></button>
                                </Tabs.List>
                            </Tabs.Root>
                        </Menu.Content>
                    </Menu.Positioner>
                </Menu.Root>
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
