import {Component} from "solid-js";

import "./Auth.css"
import {authSave} from "../lib/auth";
import {setLogged} from "../App";

const Auth: Component = () => {

    let inputRef;

    return (
        <div id="c-auth">
            <h1 id="pifi">PiFi</h1>
            <input class="card" ref={inputRef} onKeyPress={async (e) => {
                if (e.key == "Enter") {
                    inputRef.disabled = true
                    setLogged(await authSave(inputRef.value))
                    inputRef.value = ""
                    inputRef.disabled = false
                }
            }} type="password" placeholder="password" required="required" value=""/>
            <button class="card pink" onClick={() => authSave(inputRef.value)}>login</button>
        </div>
    )
}

export default Auth