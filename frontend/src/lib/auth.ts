import {api} from "./var"

import sha256 from "fast-sha256";

export const auth = async (): Promise<boolean> => {

    const rsp: Response = await fetch(api + "api/auth", {
        credentials: "include"
    })

    return rsp.status === 200
}

export const authSave = async (password: string): Promise<boolean> => {
    const hash: Uint8Array = sha256(new TextEncoder().encode(password))
    const hashArray: number[] = Array.from(hash)
    const hashString: string = hashArray.map((n: number) => n.toString(16).padStart(2, "0")).join("")
    document.cookie = "token="+hashString+";path=/"
    return await auth()
}