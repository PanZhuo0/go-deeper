# linux 运维常用命令
linux 运维常用的命令
```bash
top
    应该需要system 少一点, user态的多一点(这样性能会好一点)
    IO wait:IO等待，IO过多是不正常的,说明代码
    soft interupt:软中断,用于网络IO,太多也说明系统存在一定问题,尽可能让soft interrupt均匀 使
        用RPC/网卡多队列,来让每个核心都用起来、DPDK等等、总之就是让CPU尽可能的去均衡工作
    st stolen:偷取,


    按e:each
    按c:command
```
```bash
    nmon
        用户态、内核态、IO

    按k:内核状态,context:上下文切换、Interrupts中断
        如果上下文切换、中断过多、可以反推出存在的问题
    按n:网络状态,packin:入包数量,packout:出包数量
    m:内存状态
    ... 很多可用的选择,而且UI也不错

```
```bash
    nload
        查看网卡的实时流量
```
```bash
    tcpflow
        查看网络情况
    tcpflow -e http
```
```bash
    demsg -T 查看
    vim /var/log/message
```
```bash
    ifconfig 查看网卡是否丢包
    netstat 成本有点高
        用netstat -s 这样成本低一点,可以查看整体网络的情况
    ss -s: 查看网络的基本情况
    vmstat: 虚拟存储的情况
    iostat: IO的情况
        r/s:读取/s
        w/s:写入/s

    iobench： 

    iotop: 
        用于查看IO被占用率
    lsof -p $pid
        查看打开文件的进程是谁?

    strace
    perf top
        可以用来查看内存/CPU的情况
```
```bash
    ethtools等等
```

![常用linux监控指令](https://www.brendangregg.com/Perf/linux_observability_tools.png "监控指令")
> 建议:有时间多去学习一下linux内核，可以提高对程序运行的理解,便于优化你的程序


