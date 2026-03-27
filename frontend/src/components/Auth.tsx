import { type Component, onMount } from "solid-js"

import "./Auth.css"
import { setLogged } from "../App.tsx"
import { authSave } from "../lib/api/auth.ts"

const Auth: Component = () => {
	let inputRef!: HTMLInputElement

	const login = async () => {
		inputRef.disabled = true
		setLogged(await authSave(inputRef.value))
		inputRef.value = ""
		inputRef.disabled = false
	}

	onMount(() => {
		if (inputRef) inputRef.focus()
	})

	return (
		// @ts-ignore
		<form id="c-auth" onsubmit="return false">
			<h1 id="pifi">PiFi</h1>
			<input
				class="card"
				ref={inputRef}
				onKeyPress={async e => {
					if (e.key == "Enter") await login()
				}}
				type="password"
				placeholder="password"
				autocomplete="current-password"
				// @ts-ignore
				required="required"
				value=""
			/>
			<button class="card pink" onClick={async () => await login()}>login</button>
		</form>
	)
}

export default Auth
