/**
 * analyzing data types
 *
 * @param variable
 * @return {String} - data type.
 *
 * @example
 * type({}); // "object"
 * type([]); // "array"
 * type(5); // "number"
 * type(null); // "null"
 * type(); // "undefined"
 * type(/abcd/); // "regex"
 * type(new Date()); // "date"
 * type("test"); // String
 */
const dataType = (obj) => {
    var str = Object.prototype.toString.call(variable);
    return str.match(/\[object (.*?)\]/)[1].toLowerCase();
}

/**
 * to judge whether variable is Object
 *
 * @param obj
 * @return {Boolean} - whether variable is Object .
 */
const isObject = (obj) => {
    return dataType(obj) === 'object';
}

/**
 * to judge whether variable is String
 *
 * @param obj
 * @return {Boolean} - whether variable is String .
 */
const isString = (obj) => {
    return dataType(obj) === 'String';
}

module.exports = {
    isObject,
    isString,
}
