import { atobUnicode } from "./other.ts"
import type { response } from "./var.ts"
import { api } from "./var.ts"

export interface logData {
	time: Date
	level: string
	msg: string
	err: string
	data: string
	ssid: string
}

export const getLog = async (): Promise<logData[]> => {
	const rsp: Response = await fetch(api + "api/log", { credentials: "include" })

	const rspJson: response = await rsp.json()

	if (rsp.status === 200 && rspJson.data) {
		const data: logData[] = []
		const lines = atobUnicode(rspJson.data as unknown as string).split("\n")

		const getLogItem = (line: string, name: string): string => {
			return new RegExp(`(?:^| )${name}="(.*?)"`).exec(line)?.[1] as string
		}

		for (const line of lines) {
			if (line.length < 8) {
				continue
			}

			const time = new Date(Date.parse(getLogItem(line, "time")))

			data.push({
				time: time,
				level: getLogItem(line, "level"),
				msg: getLogItem(line, "msg"),
				err: getLogItem(line, "err"),
				data: getLogItem(line, "data"),
				ssid: getLogItem(line, "ssid"),
			})
		}

		return data
	}

	return []
}
