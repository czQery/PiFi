import type {Component} from "solid-js"

import "./Auth.css"
import {authSave} from "../lib/auth.ts"
import {setLogged} from "../App.tsx"

const Auth: Component = () => {

	let inputRef

	return (
		<form id="c-auth" onsubmit="return false">
			<h1 id="pifi">PiFi</h1>
			<input class="card" ref={inputRef} onKeyPress={async (e) => {
				if (e.key == "Enter") {
					inputRef.disabled = true
					setLogged(await authSave(inputRef.value))
					inputRef.value = ""
					inputRef.disabled = false
				}
			}} type="password" placeholder="password" autocomplete="current-password" required="required" value=""/>
			<button class="card pink" onClick={() => authSave(inputRef.value)}>login</button>
		</form>
	)
}

export default Auth