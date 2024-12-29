export const addZero = (num: number) => {

    let str = num.toString()

    switch (str.length) {
        case 0:
            return "00"
        case 1:
            return ("0" + str)
        default:
            return str
    }
}

// source: https://stackoverflow.com/questions/30106476/using-javascripts-atob-to-decode-base64-doesnt-properly-decode-utf-8-strings
export const atobUnicode = (str: string) => {
    return decodeURIComponent(atob(str).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)
    }).join(''))
}