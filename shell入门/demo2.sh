echo "请输入你的姓名"
# read name
name=$1
channel=$2
echo "你好,$name,欢迎来到$channel!"

# $# 传递给脚本或参数的位置参数的个数 
# $? 上一个命令的退出状态码
# $* 传递给脚本或函数的位置参数
# $@ 传递给脚本或函数的位置参数
# $$ 当前shell的进程ID
# $! 最后一个后台命令的进程ID
# $0 当前脚本名称
# $1-n 脚本或函数的位置参数

# 环境变量,只会保持在会话存在时
export name=zhangsan
export channel=xxx 

# 环境变量，持久有效   
# 用户家目录下的 .bashrc 每次打开终端前执行(建议放这) .profile用户登录时执行只执行一次
# /etc/bash下的环境变量 对所有所有用户有效,
# Attention: -------> 修改.bash.rc记得source一下

