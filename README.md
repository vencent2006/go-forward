# go-examples

go practice examples

## Introduction

    casual coding，write some demos for review

## List

| Name                | Type                   | Desc                                                         |
| ------------------- | ---------------------- | ------------------------------------------------------------ |
| chat_demo_im        | IM                     | IM with websocket and tcp                                                             |
| chat_demo_websocket | IM                     | IM with websocket                                            |
| cobra_word          | package，cmd           | cobra for command line tool                                  |
| error_unwrap        | package，error         | above Go 1.13，errors standard usage(wrap, unwrap,as, is)    |
| escape_analysis     | memory，gc             | memory escape analysis                                       |
| false_sharing       | memory，padding        | use memory padding to impove alloc and size                  |
| flag_parse          | package，config        | flag                                                         |
| gin_validation      | package，valid         | valid                                                        |
| gops_demo           | analyze tool           | Google official tool `gops`                                      |
| logrus_demo         | package，log           | logrus config time display                                   |
| print_caller        | practice               | print function caller                                        |
| serialization       | network, serialization | compare big endian ,little endian; protocolbuf, json, customize |
| string_concat       | package, string        | compare string cost: "A"+"B" vs fmt.Sprintf                  |
| zero_garbage        | memory                 | alloc and sync.Pool                                          |