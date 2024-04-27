#!/bin/bash

FILE="lhBasic.sh"

tpid=$(ps -ef | grep $FILE | grep -v grep | awk '{print $2}')

if [ -z "$tpid" ]
   then
   echo "查询进程不存在"
else
   kill -9 $tpid
   echo "$tpid进程终止"
fi