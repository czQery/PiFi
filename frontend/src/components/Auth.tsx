import {Component} from "solid-js";

import "./Auth.css"
import {authSave} from "../hp/auth";

const Auth: Component = () => {

    let inputRef;

    return (
        <div id="c-auth">
            <input ref={inputRef} onKeyPress={(e) => {(e.key == "Enter" ? authSave(inputRef.value) : undefined)}} type="password" placeholder="password" required="required" value=""/>
            <button onClick={() => authSave(inputRef.value)}>login</button>
        </div>
    )
}

export default Auth