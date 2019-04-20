这是一个实现golang对json字符串进排序，并根据Key进行排序，实现js的签名，golang验签的功能

### 规则
- GET 方式
hmacMD5(sort(querystring)+timestamp)
- POST 方式,根据content-type来做区分
1. Application/json
hmacMD5(sort(querystring)+ 有序json字符串 + timestamp)
2. 其他
hmacMD5(sort(querystring,form表单整体排充))
