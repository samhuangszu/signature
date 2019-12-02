这是一个实现golang对json字符串进排序，并根据Key进行排序，实现js的签名，golang验签的功能

### 关于key 
双方约定

### 规则
- GET 方式
hmacMD5(sort(queryString)+timestamp)
- POST 方式,根据content-type来做区分
1. Application/json或Application/javascript
hmacMD5(sort(querystring)+ 有序json字符串 + timestamp)
2. 其他
hmacMD5(sort(querystring,form表单整体排序))

### 加签方式 
request header中添加字段：
SIGNATURE  签名结果 
TIMESTAMP  当前时间截 

### 验签方式 
从request header中获取到如下值：
SIGNATURE
TIMESTAMP
及根据规则由服务端生成一个sign 进行比对
1.SIGNATURE!=sign 验签失败 
2.TIMESTAMP 不在服务允许范围，验签失败

### PS
端与服务对时，做比对，校正端与服务器时间
