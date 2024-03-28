import {render} from "solid-js/web"
import {Route, Router} from "@solidjs/router"
import App from "./App"

render(() => (
    <Router root={App}>
        <Route path="/dash" component={() => {
            return <div>dash</div>
        }}/>
        <Route path="/scan" component={() => {
            return <div>scan</div>
        }}/>
        <Route path="/portal" component={() => {
            return <div>portal</div>
        }}/>
    </Router>
), document.getElementById("root")!)
