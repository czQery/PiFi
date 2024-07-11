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