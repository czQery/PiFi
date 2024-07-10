import {api, response} from "./var"

export interface statsData {
    cpu: number
    mem_total: number
    mem_used: number
}

export const getStats = async (): Promise<statsData> => {
    const rsp: Response = await fetch(api + "/api/stats", {
        credentials: "include"
    })

    const rspJson: response = await rsp.json()

    if (rsp.status === 200 && rspJson.data) {
        return rspJson.data as statsData
    }

    return {cpu: 0, mem_total: 0, mem_used: 0}
}