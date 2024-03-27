import type {Component} from "solid-js"

import "./App.css"

import logo from "./assets/logo.svg";
import {Tabs} from "@ark-ui/solid";

const App: Component = () => {
    return (
        <>
            <header>
                <div class="container">
                    <img id="header-logo" src={logo} alt="logo"/>
                    <Tabs.Root value="dash">
                        <Tabs.List>
                            <Tabs.Trigger value="dash">dash</Tabs.Trigger>
                            <Tabs.Trigger value="scan">scan</Tabs.Trigger>
                            <Tabs.Trigger value="portal">portal</Tabs.Trigger>
                            <Tabs.Indicator/>
                        </Tabs.List>
                    </Tabs.Root>
                </div>
            </header>
            <main></main>
        </>
    )
}

export default App
