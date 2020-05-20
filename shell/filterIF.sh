#!/bin/bash
# Author: wqt.acc@gmail.com
#
# 本代码用途：
# 通过获取snmp设备的IF-MIB::ifNumber.0得到网络端口数量（包含虚拟端口等）
# 依次查找网络端口的ifIndex,ifDescr,ifAlias，并打印到终端

#网络设备IP地址
IP=$1
#snmp v2的认证字符串
COMMUNITY=$2

LLLL=99999999
SPEC=99999999

#端口序号
snmpIndex='IF-MIB::ifIndex'
#端口描述
snmpDescr='IF-MIB::ifDescr'
#端口别名
snmpAlias='IF-MIB::ifAlias'
#端口数量是if-MIB::ifNumber.0
snmpIfNumber='IF-MIB::ifNumber'

#
IPFORMAT='([0-9]{1,3}\.){3}[0-9]{1,3}'

#帮助信息
Usage() {
    local NAME=$0
    cat <<EOF
Usage: findIF IP COMMUNITY [port_number]

    IP:         192.168.1.1
    COMMUNITY:  public
Optional:
    port_number: special port, 1-?

Example: $NAME 192.168.1.1 public
.
EOF
}

IP=`echo -n $1 | egrep -o "$IPFORMAT"`
#检查参数数量
if [ $# -ne 3 ]; then
    if [ $# -ne 2 -o "a${IP}" = "a" ]; then
        Usage
        exit 1
    fi
else
    specTmp=`echo -n $3 |egrep -o '[0-9]+'`
    #test
    #echo '------------------'
    #echo $specTmp
    #echo '------------------'
    if [ "a${specTmp}" != "a" ]; then
        SPEC=$specTmp
    fi
fi

#echo '------------------'
#echo $SPEC
#echo '------------------'

#使用snmpwalk命令获取设备的信息
getInfo() {
    #$1是否指定
    if [ "a$1" = "a" ]; then
        echo -n 'getInfo() $1 $2 $3, $1 is ip, $2 is one of $snmpIndex, $snmpDescr,$snmpAlias, $3 is number'
        return 2
    fi
    #判断$2
    case $2 in
        $snmpIndex|$snmpDescr|$snmpAlias|$snmpIfNumber)
            echo -n `snmpwalk -v 2c -c $COMMUNITY $IP ${2}.$3`
            return 0
            ;;
        *)
            echo -n 'getInfo() $1 $2 $3, $1 is ip, $2 is one of $snmpIndex, $snmpDescr,$snmpAlias, $3 is number'
            return 1
    esac
}

#获取网络端口数量
number=`getInfo $IP ${snmpIfNumber} 0`

#获取number是否是空字符串
if [ "a${number}" = "a" ]; then
    echo 'getInfo failed, return nothing'
    exit 1
fi

#获取网络端口数
number=`echo -n $number |sed -r -n 's/[0-9A-Za-z:.-]+ = INTEGER: ([0-9]+)/\1/p'`
if [ ${number} -eq 0 ]; then
    echo 'the number of interfaces is zero'
    exit 1
fi

onlyOne=0
i=1

if [ $SPEC -le 0 ]; then
    Usage
    exit 1
elif [ $SPEC -ne $LLLL ]; then
    i=$SPEC
    number=$SPEC
    onlyOne=1
fi

#循环打印端口的index|descr|alias
while [ $i -le $number ]; do
    res=`getInfo $IP ${snmpIndex} $i`
    noSuch=`echo -n $res |grep -i -o 'No Such Instance'`

    echo ${res}

    if [ "a${noSuch}" != "a" ]; then
        ((number++))
    else
        echo `getInfo $IP ${snmpDescr} $i`
        echo `getInfo $IP ${snmpAlias} $i`
    fi

    if [ $onlyOne -eq 1 ]; then
        break
    fi
    ((i++))
done

#end
