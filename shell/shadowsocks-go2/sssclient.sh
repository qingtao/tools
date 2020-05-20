#!/bin/bash

# 服务器地址
ADDR=""
# 服务端口
PORT=""
# 本地IP
LOCALIP=""
# 本地服务socks端口
LOCALPORT="1080"
# 加密方法
CRYPTOMOTHED="aes-256-gcm"
# 密码
PASSWORD=""

# go-shadowsocks2的实际名称
CMD="shadowsocks2-linux"
# go-shadowssocks2所在目录,不包含命令
SHADOWSOCKS2DIR="$HOME/tools/goss2"

#####################################
# 以下内容不需要修改             #
#####################################
# 由SHADOWSOCKS2DIR和CMD连接得到绝对路径
SHADOWSOCKS2="$SHADOWSOCKS2DIR/$CMD"

ssspid=""
# 获取进程的PID
get_pid()
{
    # 查找运行中的go-shadowsocks2进程PID
    local cmd="${CMD} -c ss://${CRYPTOMOTHED}:${PASSWORD}@${ADDR}:${PORT} -socks ${LOCALIP}:${LOCALPORT}"
    ssspid=`ps -eo pid,args |grep "${cmd}" |grep -v 'grep' |awk '{print $1}'`
}

case $1 in
    start)
        get_pid
        if [ "a$ssspid" != "a" ]; then
            echo "${CMD} is already running: $ssspid"
            exit 0
        fi

        # 测试连接时使用"sssclient start -v"
        if [ "a$2" != "a" -a "${2}" = "-v" ]; then
            $SHADOWSOCKS2 \
                -c "ss://${CRYPTOMOTHED}:${PASSWORD}@${ADDR}:${PORT}" \
                -socks ${LOCALIP}:${LOCALPORT} -verbose
        else
            $SHADOWSOCKS2 \
                -c "ss://${CRYPTOMOTHED}:${PASSWORD}@${ADDR}:${PORT}" \
                -socks ${LOCALIP}:${LOCALPORT} >/tmp/go-shadowsocks.log 2>&1 &
            echo "${CMD} start success"
        fi

        ;;
    stop)
        get_pid
        if [ "a$ssspid" = "a" ]; then
            echo "${CMD} is not running"
            exit 1
        fi
        
        kill $ssspid
        echo "${CMD} stop success"
        ;;
    status)
        get_pid
        if [ "a$ssspid" = "a" ]; then
            echo "${CMD} is not running"
        else
            echo "${CMD} is running: $ssspid"
        fi
        ;;
    restart)
        $0 stop
        sleep 3
        $0 start
        ;;
    *)
        echo "Usage: $0 (start|stop|restart|status)"
        exit 1
esac

exit 0

