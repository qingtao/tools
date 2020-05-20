#!/bin/bash

# Author: wqt.acc@gmail.com
# Desc: 导出mongodb数据库的所有集合现有索引
# Remark: 移除敏感信息存储为通用工具
# mongodb版本大于3.6以及4.0测试通过

HOST="127.0.0.1:27017"
DATABASE="database"

# 认证用户和认证数据库
AUTH_USER="username"
AUTH_PASS="password"
AUTH_DATABASE="admin"

# mongodb 连接
MONGO="mongo"
MGO="$MONGO ${HOST}/${DATABASE} --username=${AUTH_USER} --password=${AUTH_PASS} -authenticationDatabase=${AUTH_DATABASE}"

FILENAME="${DATABASE}_index.txt"

collections=`$MGO <<"EOF"
show collections;
exit;

EOF`

# 查询集合错误时直接退出
if [ $? -ne 0 ]; then
    exit 1
fi

collections=`echo $collections|sed -E 's/.*server version: [0-9]+.[0-9]+.[0-9]+ (.*)/\1/'`

# 打印查询的集合名称
echo "--- ${DATABASE}库的所有集合 ---"
echo $collections
echo "------------------------------"

echo -n "" |tee $FILENAME

# 循环查询集合索引
for collection in $collections; do
    index=`$MGO <<EOF
db.getCollection("${collection}").getIndexes();
exit;

EOF`

    # 打印并写入文件, 文件名称以数据库名称开头
    echo "--- ${collection} ---" |tee -a $FILENAME
    echo $index|sed -E 's/.*server version: [0-9]+.[0-9]+.[0-9]+ (.*)/\1/' |tee -a $FILENAME
    echo "---------------------" |tee -a $FILENAME

done