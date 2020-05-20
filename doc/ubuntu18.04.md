#### ubuntu18.04后修改打开文件数量

1. 方法一:

```shell
# 重启后失效
ulimit -n 4096
```

1. 方法二:


修改文件/etc/systemd/system.conf中以下内容:

```shell
# 重启后生效
#DefaultLimitNOFILE=
DefaultLimitNOFILE=65535
```
