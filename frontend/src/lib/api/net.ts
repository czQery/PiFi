import { api, type response } from "../var.ts"

export const getNet = (name: string) => {
	window.open(api + "api/net/" + name, "_blank")
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
