# kc
这是一个操作k8s api的client工具，可以类似kubectl去操作kubernetets资源

## Prepare
如下是windows中编译的过程：

1.打开Terminal 执行命令：
set GOARCH=amd64
set GOOS=linux
go build xx.go

2.会生成一个没有后缀的xx二进制文件

3.将该文件放入linux系统/usr/local/bin目录下

4.赋予权限 chmod 777 xx

5.执行 ./xx

## Example
执行 ```kc get po -n example```
效果如下：
```
+------+---------+-----------+------------------+
| NAME | STATUS  | NAMESPACE |       NODE       |
+------+---------+-----------+------------------+
| test | Running |  example  | infra.main.node2 |
+------+---------+-----------+------------------+
```

## Usage
```kc -h```

