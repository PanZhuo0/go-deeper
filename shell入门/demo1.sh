#!/bin/bash

# 定义文件名
filename="example.txt"

# 检查文件是否存在
if [ -f "$filename" ]; then
    echo "文件 $filename 存在。"
    # 可以在这里添加一些操作，比如读取文件内容
    cat "$filename"
else
    echo "文件 $filename 不存在。"
    # 可以在这里添加一些操作，比如创建文件
    touch "$filename"
    echo "文件已创建。"
fi