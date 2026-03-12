export const addZero = (num: number) => {
	const str = num.toString()

	switch (str.length) {
		case 0:
			return "00"
		case 1:
			return "0" + str
		default:
			return str
	}
}

export const addZeroDecimal = (num: number) => {
	let str = num.toFixed(1)

	switch (str.length) {
		case 3:
			return "00" + str
		case 4:
			return "0" + str
		default:
			return str
	}
}

// source: https://stackoverflow.com/questions/30106476/using-javascripts-atob-to-decode-base64-doesnt-properly-decode-utf-8-strings
export const atobUnicode = (str: string) => {
	return decodeURIComponent(atob(str).split("").map(c => "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2)).join(""))
}

// source: https://stackoverflow.com/questions/1026069/how-do-i-make-the-first-letter-of-a-string-uppercase-in-javascript
export const toTitle = (str: string) => {
	return String(str).charAt(0).toUpperCase() + String(str).slice(1)
}
