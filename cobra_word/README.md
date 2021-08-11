# cobra的使用

* 总体
```shell script
> go run main.go --help
Usage:
   [command]

Available Commands:
  help        Help about any command
  word        单词格式转换

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.

```


* 子命令
```shell script
$ go run main.go word --help
该子命令支持各种单词格式转换，模式如下：
1：全部转大写
2：全部转小写
3：下划线转大写驼峰
4：下划线转小写驼峰
5：驼峰转下划线

Usage:
   word [flags]

Flags:
  -h, --help         help for word
  -m, --mode int8    请输入单词转换的模式
  -s, --str string   请输入单词内容

```

cmd/root.go里，rootCmd实际上是没有定义"Use"，其实可以把最终的执行名称写进入

比如，最后执行程序的名称是"word"，那么就在rootCmd定义为Use:"word"