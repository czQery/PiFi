import type {response} from "./var.ts"
import {api} from "./var.ts"

export const getPortals = async (): Promise<string[]> => {
	const rsp: Response = await fetch(api + "api/portals", {
		credentials: "include"
	})

	const rspJson: response = await rsp.json()

	if (rsp.status === 200 && rspJson.data) {
		return rspJson.data as string[]
	}

	return []
}