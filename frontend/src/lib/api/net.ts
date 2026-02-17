import { api, type response } from "../var.ts"

export interface netDWPAData {
	file: string
	ssid: string
	bssid: string
	net: boolean
}

export const getNetDWPA = async (): Promise<netDWPAData[]> => {
	const rsp: Response = await fetch(api + "api/net/dwpa", { credentials: "include" })
	const rspJson: response = await rsp.json()

	if (rsp.status === 200 && rspJson.data) {
		return rspJson.data as netDWPAData[]
	}

	return [] as netDWPAData[]
}

export const markNet = async (name: string, value: boolean): Promise<response> => {
	const rsp: Response = await fetch(api + "api/net/" + name + "?value=" + value.toString(), { method: "PATCH", credentials: "include" })

	if (rsp.status === 200) {
		return { message: "" } as response
	}

	return await rsp.json()
}

export const upNet = async (name: string): Promise<response> => {
	const rsp: Response = await fetch(api + "api/net/" + name, { method: "POST", credentials: "include" })

	if (rsp.status === 200) {
		return { message: "" } as response
	}

	return await rsp.json()
}
