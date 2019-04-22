var md5 = require('./md5.js')
export function Sign(obj, method, timeStamp) {
    if (!obj) {
        obj = ''
    }
    function sortKey(o) {
        let arr = Object.keys(o)
        let sortArr = arr.sort()
        let newO = {}
        sortArr.map(item => {
            newO[item] = o[item]
        })

        for (let k in newO) {
            if (newO[k] instanceof Array) {
                newO[k] = newO[k].sort()
                // console.log(newO[k])
                newO[k].map((tmp, index) => {
                    if (tmp instanceof Object) {
                        // console.log(tmp)
                        newO[k][index] = sortKey(tmp)
                    }
                })
            }

            if (newO[k] instanceof Object) {
                newO[k] = sortKey(newO[k])
            }
        }
        return newO
    }

    let key = ''
    let json = sortKey(obj)
    // post原样处理
    // get方法取对象第一层排序
    if (method.toUpperCase() == 'POST') {
        if (JSON.stringify(json) === '{}') {
            key = `${timeStamp}`
        } else {
            key = `${JSON.stringify(json)}${timeStamp}`
        }
    } else {
        let str = ''
        for (let k in json) {
            str += `${k}=${json[k]}&`
        }
        let arr = str.split('&')
        arr.length -= 1
        key = `${arr.join('&')}${timeStamp}`
    }
    let sign = md5(key, '0f90529eeccc1539b5cf6f0101a97ff2')
    return sign
}