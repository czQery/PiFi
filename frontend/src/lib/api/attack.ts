import { api, type response } from "../var.ts"

export const getAttackDeauth = async (): Promise<string[]> => {
	const rsp: Response = await fetch(api + "api/attack/deauth", { credentials: "include" })
	const rspJson: response = await rsp.json()

	if (rsp.status === 200) {
		return rspJson.data as string[]
	}

	return []
}

export const setAttackDeauth = async (enable: boolean, target: string): Promise<response> => {
	const rsp: Response = await fetch(api + "api/attack/deauth?target=" + target, { method: enable ? "POST" : "DELETE", credentials: "include" })
	if (rsp.status === 200) {
		return { message: "" } as response
	}

	return await rsp.json()
}
