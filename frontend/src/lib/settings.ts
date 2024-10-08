import {api, response} from "./var"

export interface settingsData {
    iface: settingsInterfaceData
}

export interface settingsInterfaceData {
    [key: string]: settingsInterfaceFieldsData
}

export interface settingsInterfaceFieldsData {
    mode: string
    ready: boolean
    ssid: string
    password: string
    channel: number
    portal: boolean
    portal_source: string
}

export const getSettings = async (): Promise<settingsData> => {
    const rsp: Response = await fetch(api + "api/settings", {
        credentials: "include"
    })

    const rspJson: response = await rsp.json()

    if (rsp.status === 200 && rspJson.data) {
        return rspJson.data as settingsData
    }

    return {iface: {}}
}

export const saveSettings = async (data: settingsData):Promise<response> => {
    const rsp: Response = await fetch(api + "api/settings", {
        method: "POST",
        credentials: "include",
        body: JSON.stringify(data),
    })

    const rspJson: response = await rsp.json()

    if (rsp.status === 200 && rspJson.data) {
        return {message: "", data: rspJson.data as settingsData}
    }

    return {message: rspJson.message, data: {iface: {}}}
}