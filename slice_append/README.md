# slice test

slice的底层拷贝非常复杂，看下  
[go中如何将切片中的数据全部直接打印出来_切片传递与指针传递到底有啥区别...
](https://blog.csdn.net/weixin_36114091/article/details/112097928)

`注意：只要在append发生扩容，返回的slice就是指向的新的数组，与原slice就没有关系了`