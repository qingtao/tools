#!/bin/bash
#
# 注意:
#   1. 本脚本用于设置sss的桌面图标, 方便打开;
#   2. 下载golang的package在go get前使用export https_proxy=socks5://127.0.0.1:1080;
#   3. vscode需要安装polipo, 然后设置vscode的代理为http://127.0.0.1:8123;
#   4. 然后没有了, 可以顺利的下载插件;

echo ""

echo "-- 请运行go get下载go-shadowsocks2或者下载二进制文件 --"
echo ""
echo "go get -v -u github.com/shadowsocks/go-shadowsocks2"
echo ""

# 这个参数按照sssclient实际路径设置
CURRENT_DIR=`pwd`

cat > shadowsocks.desktop <<EOF
[Desktop Entry]
Version=2.0
Terminal=false
Type=Application
Name=go-shadowsocks2
Exec=sh ${CURRENT_DIR}/sssclient.sh start
Icon=${CURRENT_DIR}/shadowsocks.png
Name[zh_CN]=影梭Go版v2

EOF

if [ ! -f "${CURRENT_DIR}/shadowsocks.desktop" ]; then
    echo "generate shadowsocks.desktop failed!"
    exit 1
fi

cat << "EOF"
设置sssclient.sh中下列参数:
ADDR=""
PORT=""
CRYPTOMOTHED=""
PASSWORD=""
CMD=""
SHADOWSOCKS2DIR=""

-- 说明 --
1. 如果一切正常，运行shadowsocks.desktop;
2. 查询运行状态：./sssclient.sh status;

-- 可选 --
1. cp -f shadowsocks.desktop $HOME/.local/share/applications/

EOF
