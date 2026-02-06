import type { response } from "../var.ts"
import { api } from "../var.ts"

export interface dbData {
	aps: number
	wigle: dbNetData
	beacondb: dbNetData
	dwpa: dbNetData
}

export interface dbNetData {
	new: number
	net: number
}

export const getDB = async (): Promise<dbData> => {
	const rsp: Response = await fetch(api + "api/db", { credentials: "include" })

	const rspJson: response = await rsp.json()

	if (rsp.status === 200 && rspJson.data) {
		return rspJson.data as dbData
	}

	return {} as dbData
}
