import {render} from "solid-js/web"
import {Route, Router} from "@solidjs/router"
import App from "./App"
import Dash from "./tabs/Dash";

render(() => (
    <Router root={App}>
        <Route path="/dash" component={Dash}/>
        <Route path="/scan" component={() => {
            return <div>scan</div>
        }}/>
        <Route path="/portal" component={() => {
            return <div>portal</div>
        }}/>
    </Router>
), document.getElementById("root")!)
