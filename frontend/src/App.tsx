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
                    <Tabs.Root value="info">
                        <Tabs.List>
                            <Tabs.Trigger value="info">Info</Tabs.Trigger>
                            <Tabs.Trigger value="portal">Portal</Tabs.Trigger>
                            <Tabs.Trigger value="scan">Scan</Tabs.Trigger>
                            <Tabs.Trigger value="other">Other</Tabs.Trigger>
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
