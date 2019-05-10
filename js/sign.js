var md5 = require('./md5.js')
import { isObject } from './object'
/*
* sign 对请求参数进行签名
* param url:请求的路径，获取queryString
* param data:发送的数据，必须为对象或null
* param contentType: 为json发送时，data发送前必须转为string
* param timeStamp:时间截
* param key:加密密钥，这里做了一次md5
* return {string}
*/
export function sign(url, data, contentType, timeStamp, key = "") {
    let obj = {}
    if (isObject(data)) {
        Object.assign(obj, data)
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
                newO[k].map((tmp, index) => {
                    if (tmp instanceof Object) {
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
    function queryObj(url) {
        if (!url) {
            return {}
        }
        if (url.indexOf("?") == -1) {
            return {}
        }
        url = url.split("?")[1];
        const urlObj = {}
        let urlArr = url.split('&')
        urlArr.forEach((item) => {
            let urlItem = item.split('=')
            Object.assign(urlObj, {
                [urlItem[0]]: decodeURIComponent(urlItem[1])
            })
        })
        return urlObj
    }
    let content = ''
    // contentType: application/json,application/javascript
    let ct = contentType.toLocaleLowerCase()
    if (ct.indexOf('application/json') != -1 ||
        ct.indexOf('application/javascript') != -1) {
        let json = sortKey(obj)
        if (JSON.stringify(json) === '{}') {
            content = `${timeStamp}`
        } else {
            content = `${JSON.stringify(json)}${timeStamp}`
        }
        let qstr = JSON.stringify(sortKey(queryObj(url)))
        if (qstr !== '{}') {
            content = `${qstr}${content}`
        }
    } else {
        Object.assign(obj, queryObj(url))
        let json = sortKey(obj)
        for (let k in json) {
            content += `${encodeURIComponent(k)}=${encodeURIComponent(json[k])}&`
        }
        content = `${content.substr(0, content.length - 1)}${timeStamp}`
    }
    key = `${key}${timeStamp}`
    key = md5(key).toLocaleLowerCase()
    let sign = md5(content, key)
    return sign.toLocaleLowerCase()
}