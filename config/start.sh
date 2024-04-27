#!/bin/bash

# 文件所在目录
HOME="$WORK/com.lh.basic"

# 拉取最新代码
git pull

# 打包go脚本
rm -rf ${HOME}/output/*
go build -o ${HOME}/output/lhBasic.sh ${HOME}/main.go
${HOME}/output/lhBasic.sh &