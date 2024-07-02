import {render} from "solid-js/web"
import {Route, Router} from "@solidjs/router"
import Dash from "./tabs/Dash";
import App from "./App";

render(() => (
    <Router root={App}>
        <Route path="/dash" component={Dash}/>
        <Route path="/settings" component={() => {
            return <div>settings</div>
        }}/>
        <Route path="/scan" component={() => {
            return <div>scan</div>
        }}/>
        <Route path="/portal" component={() => {
            return <div>portal</div>
        }}/>
    </Router>
), document.getElementById("root")!)
