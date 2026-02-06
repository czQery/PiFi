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
