import { Route, Router } from "@solidjs/router"
import { render } from "solid-js/web"
import App from "./App.tsx"
import Dash from "./tabs/Dash.tsx"
import Scan from "./tabs/Scan.tsx"
import Settings from "./tabs/Settings.tsx"

export const base: string = "/pifi"

render(() => (
	<Router root={App} base={base}>
		<Route path={"dash"} component={Dash} />
		<Route path={"settings"} component={Settings} />
		<Route path={"scan"} component={Scan} />
		<Route
			path={"portal"}
			component={() => {
				return <div>portal</div>
			}}
		/>
	</Router>
), document.getElementById("root")!)
