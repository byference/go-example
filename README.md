# 从零开始的go语言学习



## 笔记
- 内存分配
    - new: 用来分配值类型内存，如int，struct,返回的是指针
    - make: 用来分配引用类型内存，如chan，map等等

- 数组
    - 长度也是数组类型的一部分，所以[5]int 和[10]int是不同的类型
    - 数组是值类型

- 切片（slice）
    - 切片是数组的一个引用
    - slice定义的时候无需设置长度
    - 切片的长度是可以改变的
    - 切片的便利方式与数组一样
    - len()可以求切片的长度
    - cap()可以求出切片的最大容量
    - copy(target, source)可以拷贝切片
    - string底层也是一个byte数组，因此也可以进行切片操作






- 其他
    - slice，channel，map都是用make来初始化的

## 参考资料

- [Go指南](http://go-tour-zh.appspot.com/welcome/1)