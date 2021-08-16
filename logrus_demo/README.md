# logrus demo

## 主要解决不输出时间的问题

```go
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
```
在logrus 1.7和1.8来看
看文章写只加`TimestampFormat`是不好使的  
需要加上`FullTimestamp`才好使  
另外记住`TimestampFormat`的格式是`2006-01-02 15:04:05`