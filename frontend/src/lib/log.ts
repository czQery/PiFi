import { atobUnicode } from "./other.ts"
import type { response } from "./var.ts"
import { api } from "./var.ts"

export interface logData {
	time: Date
	level: string
	msg: string
	items: logDataItem[]
}

export interface logDataItem {
	name: string
	value: string
	color: string
}

export const getLog = async (): Promise<logData[]> => {
	const rsp: Response = await fetch(api + "api/log", { credentials: "include" })

	const rspJson: response = await rsp.json()

	if (rsp.status === 200 && rspJson.data) {
		const data: logData[] = []
		const lines = atobUnicode(rspJson.data as unknown as string).split("\n")
		const regex = /([^\s=]+)="([^"]*)"/g

		for (const line of lines) {
			if (line.length < 8) {
				continue
			}

			let entry: logData = { items: [] as logDataItem[] } as logData

			for (const match of line.matchAll(regex)) {
				switch (match[1]) {
					case "time":
						entry.time = new Date(match[2])
						break
					case "level":
						entry.level = match[2]
						break
					case "msg":
						entry.msg = match[2]
						break
					case "err":
						entry.items.push({ name: "err", value: match[2], color: "var(--red)" })
						break
					case "data":
					case "ssid":
					case "rssi":
						entry.items.push({ name: match[1], value: match[2], color: "var(--green)" })
						break
					case "observer":
					case "transids":
					case "iface":
						entry.items.push({ name: match[1], value: match[2], color: "var(--blue)" })
						break
				}
			}

			data.push(entry)
		}

		return data
	}

	return []
}
