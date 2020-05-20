#!/bin/bash
# Author: wqt.acc@gmail.com
# -- 读取香港天文台(https://data.weather.gov.hk/gts/time/conversion1_text_c.htm)的年历数据,并保存在本地 --

YEAR_START=1901
YEAR_END=2100
DATADIR="data"
BASEURL="https://data.weather.gov.hk/gts/time/calendar/text/"

if [ ! -d  $DATADIR ]; then
    mkdir ${DATADIR}
fi

curl ${BASEURL}T[${YEAR_START}-${YEAR_END}]c.txt -o "${DATADIR}/T#1c.txt"
echo "download files done!"

files=`ls ${DATADIR}/T*.txt`

for file in $files; do
    iconv -f BIG5 -t UTF8 $file |tee $file
done
echo "conver BIG5 to UTF8 ok!"

exit 0
