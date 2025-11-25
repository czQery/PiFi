import { api } from "./var.ts"

export interface bettercapWifiDataResponse {
	aps: bettercapWifiData[]
}

export interface bettercapWifiData {
	mac: string
	hostname: string
	channel: number
	encryption: string
	cipher: string
	authentication: string
}

export const getBettercapWifi = async (): Promise<bettercapWifiData[]> => {
	const rsp: Response = await fetch(api + "api/bettercap/session/wifi", { credentials: "include" })

	const rspJson: bettercapWifiDataResponse = await rsp.json()

	if (rsp.status === 200 && rspJson.aps) {
		return rspJson.aps as bettercapWifiData[]
	}

	return []
}
