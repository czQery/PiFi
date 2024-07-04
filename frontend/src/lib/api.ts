import {api, response} from "./var";

export interface statsData {
    cpu: number
    mem_total: number
    mem_used: number
}

export interface logData {
    time: Date
    level: string
    msg: string
    err: string
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

export const getLog = async (): Promise<logData[]> => {
    const rsp: Response = await fetch(api + "/api/log", {
        credentials: "include"
    })

    const rspJson: response = await rsp.json()

    if (rsp.status === 200 && rspJson.data) {

        let data: logData[] = []
        let lines = atob(rspJson.data as string).split("\n")

        const getLogItem = (line: string, name: string): string => {
            return (new RegExp(`${name}="(.*?)"`).exec(line))?.[1]
        }

        for (const line of lines) {
            if (line.length < 8) {
                continue
            }

            let time = new Date(Date.parse(getLogItem(line, "time")))

            data.push({time: time, level: getLogItem(line, "level"), msg: getLogItem(line, "msg"), err: getLogItem(line, "err")})
        }

        return data
    }

    return []
}