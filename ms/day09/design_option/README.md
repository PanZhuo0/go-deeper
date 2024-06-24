# Go语言中的option 设计模式

解决的问题
    1.go 语言中如何设置函数默认参数

```python
def foo(msg,name='zhangsan'):
    # python中的默认参数
    pass

foo("hello")
foo("hello",name="lisi")

```

Option 模式
1.利用Go的不定长参数


