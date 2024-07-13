import {api, response} from "./var"

export interface logData {
    time: Date
    level: string
    msg: string
    err: string
}

export const getLog = async (): Promise<logData[]> => {
    const rsp: Response = await fetch(api + "api/log", {
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