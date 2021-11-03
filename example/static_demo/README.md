# static demo

static file server

## Introduction

[static github link](https://github.com/gin-contrib/static)

```bash
# 访问ping，得到test
$ curl http://localhost:8080/ping        
test

# 访问coredemo.log，就是访问/tmp/coredemo.log，获取该文件内容
$ curl http://localhost:8080/coredemo.log                                                                                                                                                                                                                                  (base) [10:52:01] my_vue  curl http://localhost:8080/coredemo.log
[Debug] 2021-11-03T00:37:08+08:00       "hello" map[]
[Fatal] 2021-11-03T00:37:08+08:00       "fatal" map[]
[Debug] 2021-11-03T10:12:45+08:00       "hello" map[]
[Fatal] 2021-11-03T10:12:45+08:00       "fatal" map[]

# 查看/tmpcoredemo.log的内容，与http访问是一致的
$ cat /tmp/coredemo.log 
[Debug] 2021-11-03T00:37:08+08:00       "hello" map[]
[Fatal] 2021-11-03T00:37:08+08:00       "fatal" map[]
[Debug] 2021-11-03T10:12:45+08:00       "hello" map[]
[Fatal] 2021-11-03T10:12:45+08:00       "fatal" map[]

```