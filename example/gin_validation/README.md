# gin_validation

## example

参考：
[学会gin参数校验之validator库，看这一篇就足够了](https://blog.csdn.net/qq_39397165/article/details/108173108)
[gin - validator 参数校验](https://studygolang.com/articles/28368)
```shell script
# localhost:8080/time2
$ curl "localhost:8080/time2?create_time=2020-10-11&update_time=2020-10-11"        
{"code":1000,"msg":"param is error"}%   

# localhost:8080/time
curl "localhost:8080/time?create_time=2020-10-11&update_time=2020-10-11" 
{"error":"Key: 'Info.CreateTime' Error:Field validation for 'CreateTime' failed on the 'timing' tag\nKey: 'Info.UpdateTime' Error:Field validation for 'UpdateTime' failed on the 'timing' tag"}
```

## customize
久和参考的是：[gin自定义验证](http://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/%E5%8F%82%E6%95%B0%E9%AA%8C%E8%AF%81/%E8%87%AA%E5%AE%9A%E4%B9%89%E9%AA%8C%E8%AF%81.html)
```shell script
# 这个其实和example是一样的，只不过是修改@温久和的报错，
# 久和引用的包是gopkg.in/go-playground/validator.v8
# 会有报错,
$ curl -X GET "http://localhost:8080/5lmh?check_in=2019-11-07&check_out=2019-11-20"
{"error":"Key: 'Booking.CheckIn' Error:Field validation for 'CheckIn' failed on the 'bookableDate' tag"}
```





