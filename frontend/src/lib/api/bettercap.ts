import { api } from "../var.ts"

export interface bettercapWifiData {
	aps: bettercapWifiAPData[]
}

export interface bettercapWifiAPData {
	mac: string
	hostname: string
	vendor: string
	channel: number
	rssi: number
	encryption: string
	cipher: string
	authentication: string
	clients: bettercapWifiClientData[]
}

export interface bettercapWifiClientData {
	mac: string
	vendor: string
	sent: number
	received: number
}

export const getBettercapWifi = async (): Promise<bettercapWifiAPData[]> => {
	const rsp: Response = await fetch(api + "api/bettercap/session/wifi", { credentials: "include" })
	const rspJson: bettercapWifiData = await rsp.json()

	if (rsp.status === 200 && rspJson.aps) {
		rspJson.aps.sort((a, b) => b.rssi - a.rssi)

		return rspJson.aps as bettercapWifiAPData[]
	}

	return []
}
