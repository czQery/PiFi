import type { response } from "./var.ts"
import { api } from "./var.ts"

export interface statsData {
	cpu: number
	mem_total: number
	mem_used: number
	hotspot: statsHotspotData
	gps: statsGPSData
}

export interface statsHotspotData {
	ssid: string
	portal: boolean
}

export interface statsGPSData {
	lat: number
	lon: number
	alt: number

	mode: number
	time: number
}

export const getStats = async (): Promise<statsData> => {
	const rsp: Response = await fetch(api + "api/stats", { credentials: "include" })

	const rspJson: response = await rsp.json()

	if (rsp.status === 200 && rspJson.data) {
		return rspJson.data as statsData
	}

	return { cpu: 0, mem_total: 0, mem_used: 0, hotspot: { ssid: "", portal: false }, gps: { lat: 0, lon: 0, alt: 0, mode: 0, time: 0 } }
}
