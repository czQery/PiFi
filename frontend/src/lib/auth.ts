import {api} from "./var";

export const auth = async (): Promise<boolean> => {
    const rsp: Response = await fetch(api + "/api/auth", {
        credentials: "include"
    })

    return rsp.status === 200
}

export const authSave = async (password: string): Promise<boolean> => {
    const hash: string = await sha256(password)
    document.cookie = "token="+hash
    return await auth()
}

export const sha256 = async (input: string): Promise<string> => {
    const msgBuf: Uint8Array = new TextEncoder().encode(input)
    const hashBuf: ArrayBuffer = await crypto.subtle.digest("SHA-256", msgBuf)
    const hashArr: number[] = Array.from(new Uint8Array(hashBuf))
    return hashArr.map((n: number) => n.toString(16).padStart(2, "0")).join("")
}